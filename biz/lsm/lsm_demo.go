package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github/driftingboy/structures-algorithm/base/tree"
)

func main() {
	fmt.Println("=== LSM树演示程序 ===")

	// 创建临时目录
	dataDir := "/tmp/lsm_demo"
	defer os.RemoveAll(dataDir)

	// 创建LSM引擎
	fmt.Println("1. 创建LSM引擎...")
	engine := tree.NewLSMEngine(dataDir, 1024, 3, 2) // 1KB内存表，3层，每层最多2个SSTable
	defer engine.Close()

	// 演示基本操作
	fmt.Println("\n2. 基本写入操作...")
	start := time.Now()

	// 写入一些数据
	testData := map[string]string{
		"用户1": "张三",
		"用户2": "李四",
		"用户3": "王五",
		"用户4": "赵六",
		"用户5": "钱七",
	}

	for key, value := range testData {
		err := engine.Put([]byte(key), []byte(value))
		if err != nil {
			log.Fatalf("写入失败: %v", err)
		}
		fmt.Printf("  写入: %s -> %s\n", key, value)
	}

	writeTime := time.Since(start)
	fmt.Printf("写入完成，耗时: %v\n", writeTime)

	// 演示读取操作
	fmt.Println("\n3. 读取操作...")
	start = time.Now()

	for key := range testData {
		value, exists := engine.Get([]byte(key))
		if exists {
			fmt.Printf("  读取: %s -> %s\n", key, string(value))
		} else {
			fmt.Printf("  读取失败: %s\n", key)
		}
	}

	readTime := time.Since(start)
	fmt.Printf("读取完成，耗时: %v\n", readTime)

	// 演示更新操作
	fmt.Println("\n4. 更新操作...")
	engine.Put([]byte("用户1"), []byte("张三（已更新）"))

	if value, exists := engine.Get([]byte("用户1")); exists {
		fmt.Printf("  更新后: 用户1 -> %s\n", string(value))
	}

	// 演示删除操作
	fmt.Println("\n5. 删除操作...")
	engine.Delete([]byte("用户3"))

	if _, exists := engine.Get([]byte("用户3")); !exists {
		fmt.Println("  用户3 已成功删除")
	}

	// 演示内存表刷新
	fmt.Println("\n6. 触发内存表刷新...")
	fmt.Println("  写入大量数据以触发内存表刷新到磁盘...")

	for i := 0; i < 20; i++ {
		key := fmt.Sprintf("批量数据_%d", i)
		value := fmt.Sprintf("这是第%d条批量数据，内容比较长，用来测试内存表刷新机制", i)
		engine.Put([]byte(key), []byte(value))
	}

	fmt.Println("  内存表已刷新到磁盘")

	// 验证数据仍然可以读取
	fmt.Println("\n7. 验证刷新后的数据...")
	for i := 0; i < 5; i++ {
		key := fmt.Sprintf("批量数据_%d", i)
		if value, exists := engine.Get([]byte(key)); exists {
			fmt.Printf("  验证: %s -> %s\n", key, string(value))
		}
	}

	// 演示压缩机制
	fmt.Println("\n8. 触发压缩机制...")
	fmt.Println("  写入更多数据以触发压缩...")

	for batch := 0; batch < 3; batch++ {
		for i := 0; i < 5; i++ {
			key := fmt.Sprintf("压缩测试_批次%d_数据%d", batch, i)
			value := fmt.Sprintf("这是压缩测试数据，批次%d，数据%d", batch, i)
			engine.Put([]byte(key), []byte(value))
		}
		time.Sleep(10 * time.Millisecond) // 让压缩有时间完成
	}

	fmt.Println("  压缩操作已完成")

	// 显示统计信息
	fmt.Println("\n9. 引擎统计信息...")
	stats := engine.GetStats()
	fmt.Printf("  内存表大小: %d 字节\n", stats["memtable_size"])
	fmt.Printf("  内存表最大大小: %d 字节\n", stats["memtable_max_size"])

	if levelCounts, ok := stats["sstable_counts_by_level"].([]int); ok {
		fmt.Println("  各层SSTable数量:")
		for level, count := range levelCounts {
			if count > 0 {
				fmt.Printf("    第%d层: %d个SSTable\n", level, count)
			}
		}
	}

	// 性能测试
	fmt.Println("\n10. 性能测试...")

	// 写入性能测试
	start = time.Now()
	for i := 0; i < 1000; i++ {
		key := fmt.Sprintf("性能测试_%d", i)
		value := fmt.Sprintf("性能测试数据_%d", i)
		engine.Put([]byte(key), []byte(value))
	}
	writeTime = time.Since(start)
	fmt.Printf("  写入1000条数据耗时: %v (平均: %v/条)\n", writeTime, writeTime/1000)

	// 读取性能测试
	start = time.Now()
	for i := 0; i < 1000; i++ {
		key := fmt.Sprintf("性能测试_%d", i)
		engine.Get([]byte(key))
	}
	readTime = time.Since(start)
	fmt.Printf("  读取1000条数据耗时: %v (平均: %v/条)\n", readTime, readTime/1000)

	fmt.Println("\n=== 演示完成 ===")
	fmt.Println("LSM树的主要特点:")
	fmt.Println("1. 写入性能极高 - 所有写入都在内存中完成")
	fmt.Println("2. 批量持久化 - 内存表满时批量写入磁盘")
	fmt.Println("3. 分层存储 - 数据按层级组织，新数据在高层")
	fmt.Println("4. 自动压缩 - 定期合并SSTable，优化存储空间")
	fmt.Println("5. 支持删除 - 通过逻辑删除标记实现")
}
