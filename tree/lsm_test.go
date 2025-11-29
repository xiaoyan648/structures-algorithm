package tree

import (
	"fmt"
	"os"
	"testing"
	"time"
)

// TestLSMBasicOperations 测试LSM树的基本操作
func TestLSMBasicOperations(t *testing.T) {
	// 创建临时目录用于测试
	dataDir := "/tmp/lsm_test"
	defer os.RemoveAll(dataDir)

	// 创建LSM引擎
	engine := NewLSMEngine(dataDir, 1024, 3, 2) // 1KB内存表，3层，每层最多2个SSTable
	defer engine.Close()

	// 测试基本写入和读取
	t.Run("基本写入读取", func(t *testing.T) {
		// 写入一些数据
		testData := map[string]string{
			"key1": "value1",
			"key2": "value2",
			"key3": "value3",
		}

		for key, value := range testData {
			err := engine.Put([]byte(key), []byte(value))
			if err != nil {
				t.Fatalf("写入失败: %v", err)
			}
		}

		// 读取数据
		for key, expectedValue := range testData {
			value, exists := engine.Get([]byte(key))
			if !exists {
				t.Fatalf("键 %s 不存在", key)
			}
			if string(value) != expectedValue {
				t.Fatalf("键 %s 的值不匹配，期望: %s, 实际: %s", key, expectedValue, string(value))
			}
		}
	})

	// 测试更新操作
	t.Run("更新操作", func(t *testing.T) {
		// 更新现有键
		err := engine.Put([]byte("key1"), []byte("updated_value1"))
		if err != nil {
			t.Fatalf("更新失败: %v", err)
		}

		value, exists := engine.Get([]byte("key1"))
		if !exists {
			t.Fatalf("更新后的键不存在")
		}
		if string(value) != "updated_value1" {
			t.Fatalf("更新失败，期望: updated_value1, 实际: %s", string(value))
		}
	})

	// 测试删除操作
	t.Run("删除操作", func(t *testing.T) {
		// 删除一个键
		err := engine.Delete([]byte("key2"))
		if err != nil {
			t.Fatalf("删除失败: %v", err)
		}

		// 验证键已被删除
		_, exists := engine.Get([]byte("key2"))
		if exists {
			t.Fatalf("键应该已被删除")
		}
	})
}

// TestLSMMemTableFlush 测试内存表刷新机制
func TestLSMMemTableFlush(t *testing.T) {
	dataDir := "/tmp/lsm_flush_test"
	defer os.RemoveAll(dataDir)

	// 创建小内存表的引擎来测试刷新
	engine := NewLSMEngine(dataDir, 100, 3, 2) // 100字节的小内存表
	defer engine.Close()

	// 写入大量数据触发内存表刷新
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("key_%d", i)
		value := fmt.Sprintf("value_%d_%s", i, "这是一个比较长的值用来测试内存表刷新机制")
		err := engine.Put([]byte(key), []byte(value))
		if err != nil {
			t.Fatalf("写入失败: %v", err)
		}
	}

	// 验证数据仍然可以读取
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("key_%d", i)
		expectedValue := fmt.Sprintf("value_%d_%s", i, "这是一个比较长的值用来测试内存表刷新机制")

		value, exists := engine.Get([]byte(key))
		if !exists {
			t.Fatalf("键 %s 不存在", key)
		}
		if string(value) != expectedValue {
			t.Fatalf("键 %s 的值不匹配", key)
		}
	}

	// 检查统计信息
	stats := engine.GetStats()
	t.Logf("引擎统计信息: %+v", stats)
}

// TestLSMCompaction 测试压缩机制
func TestLSMCompaction(t *testing.T) {
	dataDir := "/tmp/lsm_compaction_test"
	defer os.RemoveAll(dataDir)

	// 创建每层最多1个SSTable的引擎来快速触发压缩
	engine := NewLSMEngine(dataDir, 50, 3, 1)
	defer engine.Close()

	// 写入多批数据触发多次压缩
	for batch := 0; batch < 5; batch++ {
		for i := 0; i < 3; i++ {
			key := fmt.Sprintf("batch_%d_key_%d", batch, i)
			value := fmt.Sprintf("batch_%d_value_%d", batch, i)
			err := engine.Put([]byte(key), []byte(value))
			if err != nil {
				t.Fatalf("写入失败: %v", err)
			}
		}

		// 等待一下让压缩完成
		time.Sleep(10 * time.Millisecond)
	}

	// 验证所有数据仍然可以读取
	for batch := 0; batch < 5; batch++ {
		for i := 0; i < 3; i++ {
			key := fmt.Sprintf("batch_%d_key_%d", batch, i)
			expectedValue := fmt.Sprintf("batch_%d_value_%d", batch, i)

			value, exists := engine.Get([]byte(key))
			if !exists {
				t.Fatalf("键 %s 不存在", key)
			}
			if string(value) != expectedValue {
				t.Fatalf("键 %s 的值不匹配", key)
			}
		}
	}

	// 检查统计信息
	stats := engine.GetStats()
	t.Logf("压缩后统计信息: %+v", stats)
}

// TestLSMConcurrent 测试并发操作
func TestLSMConcurrent(t *testing.T) {
	dataDir := "/tmp/lsm_concurrent_test"
	defer os.RemoveAll(dataDir)

	engine := NewLSMEngine(dataDir, 1024, 3, 2)
	defer engine.Close()

	// 并发写入
	done := make(chan bool, 10)

	for i := 0; i < 10; i++ {
		go func(workerID int) {
			defer func() { done <- true }()

			for j := 0; j < 10; j++ {
				key := fmt.Sprintf("worker_%d_key_%d", workerID, j)
				value := fmt.Sprintf("worker_%d_value_%d", workerID, j)

				err := engine.Put([]byte(key), []byte(value))
				if err != nil {
					t.Errorf("并发写入失败: %v", err)
					return
				}
			}
		}(i)
	}

	// 等待所有goroutine完成
	for i := 0; i < 10; i++ {
		<-done
	}

	// 验证数据
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			key := fmt.Sprintf("worker_%d_key_%d", i, j)
			expectedValue := fmt.Sprintf("worker_%d_value_%d", i, j)

			value, exists := engine.Get([]byte(key))
			if !exists {
				t.Fatalf("键 %s 不存在", key)
			}
			if string(value) != expectedValue {
				t.Fatalf("键 %s 的值不匹配", key)
			}
		}
	}
}

// BenchmarkLSMWrite 写入性能测试
func BenchmarkLSMWrite(b *testing.B) {
	dataDir := "/tmp/lsm_bench_write"
	defer os.RemoveAll(dataDir)

	engine := NewLSMEngine(dataDir, 1024*1024, 3, 2) // 1MB内存表
	defer engine.Close()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("key_%d", i)
		value := fmt.Sprintf("value_%d", i)
		engine.Put([]byte(key), []byte(value))
	}
}

// BenchmarkLSMRead 读取性能测试
func BenchmarkLSMRead(b *testing.B) {
	dataDir := "/tmp/lsm_bench_read"
	defer os.RemoveAll(dataDir)

	engine := NewLSMEngine(dataDir, 1024*1024, 3, 2)
	defer engine.Close()

	// 预先写入一些数据
	for i := 0; i < 1000; i++ {
		key := fmt.Sprintf("key_%d", i)
		value := fmt.Sprintf("value_%d", i)
		engine.Put([]byte(key), []byte(value))
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("key_%d", i%1000)
		engine.Get([]byte(key))
	}
}

// ExampleLSMEngine 演示LSM引擎的使用
func ExampleLSMEngine() {
	// 创建LSM引擎
	engine := NewLSMEngine("/tmp/example_lsm", 1024, 3, 2)
	defer engine.Close()

	// 写入数据
	engine.Put([]byte("name"), []byte("张三"))
	engine.Put([]byte("age"), []byte("25"))
	engine.Put([]byte("city"), []byte("北京"))

	// 读取数据
	if name, exists := engine.Get([]byte("name")); exists {
		fmt.Printf("姓名: %s\n", string(name))
	}

	if age, exists := engine.Get([]byte("age")); exists {
		fmt.Printf("年龄: %s\n", string(age))
	}

	// 更新数据
	engine.Put([]byte("age"), []byte("26"))

	// 删除数据
	engine.Delete([]byte("city"))

	// 检查删除结果
	if _, exists := engine.Get([]byte("city")); !exists {
		fmt.Println("城市信息已删除")
	}

	// 输出:
	// 姓名: 张三
	// 年龄: 25
	// 城市信息已删除
}
