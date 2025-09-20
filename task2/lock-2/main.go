package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	// 使用 int64 类型的原子计数器
	var counter int64

	var wg sync.WaitGroup
	numGoroutines := 10
	incrementsPerGoroutine := 1000

	// 启动10个协程
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			// 每个协程对计数器进行1000次原子递增操作
			for j := 0; j < incrementsPerGoroutine; j++ {
				atomic.AddInt64(&counter, 1)
			}

			fmt.Printf("协程 %d 完成工作\n", id)
		}(i)
	}

	// 等待所有协程完成
	wg.Wait()

	// 输出最终结果
	expected := int64(numGoroutines * incrementsPerGoroutine)
	actual := atomic.LoadInt64(&counter)
	fmt.Printf("\n最终计数器值: %d\n", actual)
	fmt.Printf("期望值: %d\n", expected)

	if actual == expected {
		fmt.Println("✅ 测试通过: 实际值与期望值一致")
	} else {
		fmt.Println("❌ 测试失败: 实际值与期望值不一致")
	}

	// 演示其他原子操作
	fmt.Println("\n=== 其他原子操作演示 ===")

	// 原子比较并交换 (Compare and Swap)
	var value int64 = 100
	swapped := atomic.CompareAndSwapInt64(&value, 100, 200)
	fmt.Printf("比较并交换 (100->200): 成功=%t, 新值=%d\n", swapped, value)

	// 再次尝试相同的比较并交换
	swapped = atomic.CompareAndSwapInt64(&value, 100, 300)
	fmt.Printf("比较并交换 (100->300): 成功=%t, 新值=%d\n", swapped, value)

	// 原子存储和加载
	atomic.StoreInt64(&value, 500)
	loaded := atomic.LoadInt64(&value)
	fmt.Printf("存储和加载: 值=%d\n", loaded)
}
