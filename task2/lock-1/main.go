package main

import (
	"fmt"
	"sync"
)

// Counter 结构体，包含一个计数器和互斥锁
type Counter struct {
	value int
	mutex sync.Mutex
}

// Increment 安全地增加计数器
func (c *Counter) Increment() {
	c.mutex.Lock()         // 获取锁
	defer c.mutex.Unlock() // 确保锁会被释放
	c.value++
}

// GetValue 安全地获取计数器值
func (c *Counter) GetValue() int {
	c.mutex.Lock()         // 获取锁
	defer c.mutex.Unlock() // 确保锁会被释放
	return c.value
}

func main() {
	// 创建计数器实例
	counter := Counter{}

	var wg sync.WaitGroup
	numGoroutines := 10
	incrementsPerGoroutine := 1000

	// 启动10个协程
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			// 每个协程对计数器进行1000次递增操作
			for j := 0; j < incrementsPerGoroutine; j++ {
				counter.Increment()
			}

			fmt.Printf("协程 %d 完成工作\n", id)
		}(i)
	}

	// 等待所有协程完成
	wg.Wait()

	// 输出最终结果
	expected := numGoroutines * incrementsPerGoroutine
	actual := counter.GetValue()
	fmt.Printf("\n最终计数器值: %d\n", actual)
	fmt.Printf("期望值: %d\n", expected)

	if actual == expected {
		fmt.Println("✅ 测试通过: 实际值与期望值一致")
	} else {
		fmt.Println("❌ 测试失败: 实际值与期望值不一致")
	}
}
