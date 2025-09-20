package main

import (
	"fmt"
	"sync"
)

func main() {
	// 创建一个缓冲大小为10的整数通道
	ch := make(chan int, 10)

	var wg sync.WaitGroup

	// 生产者协程：发送100个整数到通道
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(ch) // 确保通道被关闭

		for i := 1; i <= 100; i++ {
			ch <- i
			fmt.Printf("生产者发送: %d\n", i)
		}
		fmt.Println("生产者完成发送")
	}()

	// 消费者协程：从通道接收并打印整数
	wg.Add(1)
	go func() {
		defer wg.Done()

		for num := range ch {
			fmt.Printf("消费者接收: %d\n", num)
		}
		fmt.Println("消费者完成接收")
	}()

	// 等待所有协程完成
	wg.Wait()
	fmt.Println("程序执行完毕")
}
