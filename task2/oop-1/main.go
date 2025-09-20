package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func printShapeInfo(s Shape) {
	fmt.Printf("面积: %.2f, 周长: %.2f\n", s.Area(), s.Perimeter())
}

func main() {
	// 创建 Rectangle 实例
	rect := Rectangle{Width: 5, Height: 3}
	fmt.Println("矩形:")
	printShapeInfo(rect)

	// 创建 Circle 实例
	circle := Circle{Radius: 4}
	fmt.Println("圆形:")
	printShapeInfo(circle)

	// 演示多态性：将不同形状存储在同一个切片中
	fmt.Println("\n使用多态:")
	shapes := []Shape{rect, circle}
	for i, shape := range shapes {
		fmt.Printf("形状 %d: ", i+1)
		printShapeInfo(shape)
	}
}
