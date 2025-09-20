package main

import "fmt"

func addTen(num *int) {
	*num += 10
}

func doubleSlice(slicePtr *[]int) {
	slice := *slicePtr
	for i := range slice {
		slice[i] *= 2
	}
}

func main() {
	num := 5
	addTen(&num)
	fmt.Println("修改后的整数值：", num)

	slice := []int{1, 2, 3, 4, 5}
	doubleSlice(&slice)
	fmt.Println("修改后的切片：", slice)
}
