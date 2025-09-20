package main

import "fmt"

// Person 结构体定义
type Person struct {
	Name string
	Age  int
}

// Employee 结构体定义，组合了 Person
type Employee struct {
	Person     // 匿名嵌入，实现组合
	EmployeeID string
	Department string
}

// PrintInfo 方法，输出员工信息
func (e Employee) PrintInfo() {
	fmt.Printf("员工信息:\n")
	fmt.Printf("  姓名: %s\n", e.Name) // 直接访问嵌入结构的字段
	fmt.Printf("  年龄: %d\n", e.Age)  // 直接访问嵌入结构的字段
	fmt.Printf("  工号: %s\n", e.EmployeeID)
	fmt.Printf("  部门: %s\n", e.Department)
}

// 可选的：为 Person 结构体添加方法
func (p Person) Introduce() {
	fmt.Printf("大家好，我是%s，今年%d岁。\n", p.Name, p.Age)
}

func main() {
	// 创建 Employee 实例
	emp := Employee{
		Person: Person{
			Name: "张三",
			Age:  30,
		},
		EmployeeID: "E1001",
		Department: "技术部",
	}

	// 调用 Employee 的方法
	emp.PrintInfo()
	fmt.Println()

	// 演示组合的特性：Employee 可以直接访问 Person 的方法
	fmt.Println("员工自我介绍:")
	emp.Introduce()
	fmt.Println()

	// 也可以直接访问嵌入结构的字段
	fmt.Printf("直接访问字段: %s的工号是%s\n", emp.Name, emp.EmployeeID)
}
