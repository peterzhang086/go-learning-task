package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	//启动奇数协程
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i += 2 {
			fmt.Println("奇数: %d\n", i)
		}
	}()

	//启动偶数协程
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 2; i <= 10; i += 2 {
			fmt.Println("偶数: %d\n", i)
		}
	}()

	//等待所有协程完成
	wg.Wait()
	fmt.Println("所有数字打印完成")
}
