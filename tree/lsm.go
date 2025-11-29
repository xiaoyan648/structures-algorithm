package tree

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"sync"
	"time"
)

// LSM树实现 - Log-Structured Merge-Tree
// 用于高性能的写入操作，广泛应用于现代数据库系统

// Entry 表示一个键值对条目
type Entry struct {
	Key       []byte
	Value     []byte
	Timestamp int64
	Deleted   bool // 标记是否被删除（用于逻辑删除）
}

// MemTable 内存表，用于缓存写入操作
type MemTable struct {
	data    map[string]*Entry
	size    int64 // 当前大小（字节）
	maxSize int64 // 最大大小阈值
	mutex   sync.RWMutex
}

// NewMemTable 创建新的内存表
func NewMemTable(maxSize int64) *MemTable {
	return &MemTable{
		data:    make(map[string]*Entry),
		maxSize: maxSize,
	}
}

// Put 向内存表添加键值对
func (mt *MemTable) Put(key, value []byte) {
	mt.mutex.Lock()
	defer mt.mutex.Unlock()

	keyStr := string(key)
	entry := &Entry{
		Key:       key,
		Value:     value,
		Timestamp: time.Now().UnixNano(),
		Deleted:   false,
	}

	// 计算大小变化
	oldSize := int64(0)
	if oldEntry, exists := mt.data[keyStr]; exists {
		oldSize = int64(len(oldEntry.Key) + len(oldEntry.Value))
	}
	newSize := int64(len(key) + len(value))
	mt.size = mt.size - oldSize + newSize

	mt.data[keyStr] = entry
}

// Get 从内存表获取值
func (mt *MemTable) Get(key []byte) (*Entry, bool) {
	mt.mutex.RLock()
	defer mt.mutex.RUnlock()

	entry, exists := mt.data[string(key)]
	return entry, exists
}

// Delete 逻辑删除键
func (mt *MemTable) Delete(key []byte) {
	mt.mutex.Lock()
	defer mt.mutex.Unlock()

	entry := &Entry{
		Key:       key,
		Value:     nil,
		Timestamp: time.Now().UnixNano(),
		Deleted:   true,
	}

	// 更新大小
	oldSize := int64(0)
	if oldEntry, exists := mt.data[string(key)]; exists {
		oldSize = int64(len(oldEntry.Key) + len(oldEntry.Value))
	}
	mt.size = mt.size - oldSize + int64(len(key))

	mt.data[string(key)] = entry
}

// ShouldFlush 检查是否需要刷新到磁盘
func (mt *MemTable) ShouldFlush() bool {
	mt.mutex.RLock()
	defer mt.mutex.RUnlock()
	return mt.size >= mt.maxSize
}

// GetAllEntries 获取所有条目（用于刷新到磁盘）
func (mt *MemTable) GetAllEntries() []*Entry {
	mt.mutex.RLock()
	defer mt.mutex.RUnlock()

	entries := make([]*Entry, 0, len(mt.data))
	for _, entry := range mt.data {
		entries = append(entries, entry)
	}

	// 按键排序
	sort.Slice(entries, func(i, j int) bool {
		return string(entries[i].Key) < string(entries[j].Key)
	})

	return entries
}

// Clear 清空内存表
func (mt *MemTable) Clear() {
	mt.mutex.Lock()
	defer mt.mutex.Unlock()
	mt.data = make(map[string]*Entry)
	mt.size = 0
}

// SSTable 排序字符串表，存储在磁盘上的不可变数据结构
type SSTable struct {
	filePath string
	entries  []*Entry
	index    map[string]int64 // 键到文件偏移的映射
	mutex    sync.RWMutex
}

// NewSSTable 创建新的SSTable
func NewSSTable(filePath string) *SSTable {
	return &SSTable{
		filePath: filePath,
		index:    make(map[string]int64),
	}
}

// WriteToDisk 将条目写入磁盘文件
func (sst *SSTable) WriteToDisk(entries []*Entry) error {
	sst.mutex.Lock()
	defer sst.mutex.Unlock()

	// 确保目录存在
	dir := filepath.Dir(sst.filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	file, err := os.Create(sst.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// 写入条目
	for _, entry := range entries {
		offset, _ := file.Seek(0, io.SeekCurrent)
		sst.index[string(entry.Key)] = offset

		// 写入键长度
		if err := binary.Write(file, binary.LittleEndian, int32(len(entry.Key))); err != nil {
			return err
		}
		// 写入键
		if _, err := file.Write(entry.Key); err != nil {
			return err
		}
		// 写入值长度
		if err := binary.Write(file, binary.LittleEndian, int32(len(entry.Value))); err != nil {
			return err
		}
		// 写入值
		if _, err := file.Write(entry.Value); err != nil {
			return err
		}
		// 写入时间戳
		if err := binary.Write(file, binary.LittleEndian, entry.Timestamp); err != nil {
			return err
		}
		// 写入删除标记
		if err := binary.Write(file, binary.LittleEndian, entry.Deleted); err != nil {
			return err
		}
	}

	sst.entries = entries
	return nil
}

// ReadFromDisk 从磁盘读取SSTable
func (sst *SSTable) ReadFromDisk() error {
	sst.mutex.Lock()
	defer sst.mutex.Unlock()

	file, err := os.Open(sst.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	var entries []*Entry
	index := make(map[string]int64)

	for {
		offset, _ := file.Seek(0, io.SeekCurrent)

		// 读取键长度
		var keyLen int32
		if err := binary.Read(file, binary.LittleEndian, &keyLen); err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		// 读取键
		key := make([]byte, keyLen)
		if _, err := file.Read(key); err != nil {
			return err
		}

		// 读取值长度
		var valueLen int32
		if err := binary.Read(file, binary.LittleEndian, &valueLen); err != nil {
			return err
		}

		// 读取值
		value := make([]byte, valueLen)
		if _, err := file.Read(value); err != nil {
			return err
		}

		// 读取时间戳
		var timestamp int64
		if err := binary.Read(file, binary.LittleEndian, &timestamp); err != nil {
			return err
		}

		// 读取删除标记
		var deleted bool
		if err := binary.Read(file, binary.LittleEndian, &deleted); err != nil {
			return err
		}

		entry := &Entry{
			Key:       key,
			Value:     value,
			Timestamp: timestamp,
			Deleted:   deleted,
		}

		entries = append(entries, entry)
		index[string(key)] = offset
	}

	sst.entries = entries
	sst.index = index
	return nil
}

// Get 从SSTable获取值
func (sst *SSTable) Get(key []byte) (*Entry, bool) {
	sst.mutex.RLock()
	defer sst.mutex.RUnlock()

	offset, exists := sst.index[string(key)]
	if !exists {
		return nil, false
	}

	file, err := os.Open(sst.filePath)
	if err != nil {
		return nil, false
	}
	defer file.Close()

	// 定位到指定偏移
	if _, err := file.Seek(offset, io.SeekStart); err != nil {
		return nil, false
	}

	// 读取条目
	var keyLen int32
	if err := binary.Read(file, binary.LittleEndian, &keyLen); err != nil {
		return nil, false
	}

	keyBytes := make([]byte, keyLen)
	if _, err := file.Read(keyBytes); err != nil {
		return nil, false
	}

	var valueLen int32
	if err := binary.Read(file, binary.LittleEndian, &valueLen); err != nil {
		return nil, false
	}

	value := make([]byte, valueLen)
	if _, err := file.Read(value); err != nil {
		return nil, false
	}

	var timestamp int64
	if err := binary.Read(file, binary.LittleEndian, &timestamp); err != nil {
		return nil, false
	}

	var deleted bool
	if err := binary.Read(file, binary.LittleEndian, &deleted); err != nil {
		return nil, false
	}

	return &Entry{
		Key:       keyBytes,
		Value:     value,
		Timestamp: timestamp,
		Deleted:   deleted,
	}, true
}

// LSMEngine LSM存储引擎
type LSMEngine struct {
	dataDir     string
	memTable    *MemTable
	sstables    [][]*SSTable // 分层存储，每层包含多个SSTable
	maxLevels   int
	maxSSTables int // 每层最大SSTable数量
	mutex       sync.RWMutex
}

// NewLSMEngine 创建新的LSM引擎
func NewLSMEngine(dataDir string, memTableSize int64, maxLevels, maxSSTables int) *LSMEngine {
	// 确保数据目录存在
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		panic(fmt.Sprintf("无法创建数据目录: %v", err))
	}

	engine := &LSMEngine{
		dataDir:     dataDir,
		memTable:    NewMemTable(memTableSize),
		sstables:    make([][]*SSTable, maxLevels),
		maxLevels:   maxLevels,
		maxSSTables: maxSSTables,
	}

	// 初始化每层的SSTable列表
	for i := 0; i < maxLevels; i++ {
		engine.sstables[i] = make([]*SSTable, 0)
	}

	return engine
}

// Put 写入键值对
func (lsm *LSMEngine) Put(key, value []byte) error {
	lsm.mutex.Lock()
	defer lsm.mutex.Unlock()

	// 写入内存表
	lsm.memTable.Put(key, value)

	// 检查是否需要刷新到磁盘
	if lsm.memTable.ShouldFlush() {
		return lsm.flushMemTable()
	}

	return nil
}

// Get 读取键值对
func (lsm *LSMEngine) Get(key []byte) ([]byte, bool) {
	lsm.mutex.RLock()
	defer lsm.mutex.RUnlock()

	// 首先在内存表中查找
	if entry, exists := lsm.memTable.Get(key); exists {
		if entry.Deleted {
			return nil, false
		}
		return entry.Value, true
	}

	// 然后在SSTable中查找（从高层到低层）
	for level := 0; level < lsm.maxLevels; level++ {
		for _, sstable := range lsm.sstables[level] {
			if entry, exists := sstable.Get(key); exists {
				if entry.Deleted {
					return nil, false
				}
				return entry.Value, true
			}
		}
	}

	return nil, false
}

// Delete 删除键
func (lsm *LSMEngine) Delete(key []byte) error {
	lsm.mutex.Lock()
	defer lsm.mutex.Unlock()

	// 在内存表中标记删除
	lsm.memTable.Delete(key)

	// 检查是否需要刷新
	if lsm.memTable.ShouldFlush() {
		return lsm.flushMemTable()
	}

	return nil
}

// flushMemTable 将内存表刷新到磁盘
func (lsm *LSMEngine) flushMemTable() error {
	// 获取内存表中的所有条目
	entries := lsm.memTable.GetAllEntries()
	if len(entries) == 0 {
		return nil
	}

	// 创建新的SSTable
	filePath := filepath.Join(lsm.dataDir, fmt.Sprintf("level0_sstable_%d.sst", time.Now().UnixNano()))
	sstable := NewSSTable(filePath)

	// 写入磁盘
	if err := sstable.WriteToDisk(entries); err != nil {
		return err
	}

	// 添加到第0层
	lsm.sstables[0] = append(lsm.sstables[0], sstable)

	// 清空内存表
	lsm.memTable.Clear()

	// 检查是否需要压缩
	if len(lsm.sstables[0]) > lsm.maxSSTables {
		return lsm.compact(0)
	}

	return nil
}

// compact 压缩指定层
func (lsm *LSMEngine) compact(level int) error {
	if level >= lsm.maxLevels-1 {
		return nil // 最后一层不压缩
	}

	// 获取当前层需要压缩的SSTable
	sstablesToCompact := lsm.sstables[level]
	if len(sstablesToCompact) == 0 {
		return nil
	}

	// 合并所有SSTable的条目
	var allEntries []*Entry
	keyMap := make(map[string]*Entry) // 用于去重，保留最新的条目

	for _, sstable := range sstablesToCompact {
		for _, entry := range sstable.entries {
			keyStr := string(entry.Key)
			if existing, exists := keyMap[keyStr]; !exists || entry.Timestamp > existing.Timestamp {
				keyMap[keyStr] = entry
			}
		}
	}

	// 转换为切片并排序
	for _, entry := range keyMap {
		allEntries = append(allEntries, entry)
	}
	sort.Slice(allEntries, func(i, j int) bool {
		return string(allEntries[i].Key) < string(allEntries[j].Key)
	})

	// 创建新的SSTable
	filePath := filepath.Join(lsm.dataDir, fmt.Sprintf("level%d_sstable_%d.sst", level+1, time.Now().UnixNano()))
	newSSTable := NewSSTable(filePath)

	// 写入磁盘
	if err := newSSTable.WriteToDisk(allEntries); err != nil {
		return err
	}

	// 删除旧的SSTable文件
	for _, sstable := range sstablesToCompact {
		os.Remove(sstable.filePath)
	}

	// 更新SSTable列表
	lsm.sstables[level] = make([]*SSTable, 0)
	lsm.sstables[level+1] = append(lsm.sstables[level+1], newSSTable)

	// 递归检查下一层是否需要压缩
	if len(lsm.sstables[level+1]) > lsm.maxSSTables {
		return lsm.compact(level + 1)
	}

	return nil
}

// Close 关闭LSM引擎
func (lsm *LSMEngine) Close() error {
	lsm.mutex.Lock()
	defer lsm.mutex.Unlock()

	// 刷新内存表
	if lsm.memTable.size > 0 {
		return lsm.flushMemTable()
	}

	return nil
}

// GetStats 获取引擎统计信息
func (lsm *LSMEngine) GetStats() map[string]interface{} {
	lsm.mutex.RLock()
	defer lsm.mutex.RUnlock()

	stats := make(map[string]interface{})
	stats["memtable_size"] = lsm.memTable.size
	stats["memtable_max_size"] = lsm.memTable.maxSize

	levelStats := make([]int, lsm.maxLevels)
	for i := 0; i < lsm.maxLevels; i++ {
		levelStats[i] = len(lsm.sstables[i])
	}
	stats["sstable_counts_by_level"] = levelStats

	return stats
}
