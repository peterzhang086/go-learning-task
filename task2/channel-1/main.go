package main

import (
	"fmt"
)

func main() {
	// 创建一个整数通道
	ch := make(chan int)

	// 启动生产者协程，生成数字并发送到通道
	go func() {
		for i := 1; i <= 10; i++ {
			ch <- i // 发送数字到通道
		}
		close(ch) // 发送完成后关闭通道
	}()

	// 在主协程中接收并打印数字
	for num := range ch {
		fmt.Printf("接收到的数字: %d\n", num)
	}

	fmt.Println("所有数字已接收完成")
}
