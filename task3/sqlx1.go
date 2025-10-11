package main

import (
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var sqxDB *sqlx.DB
var sqxerr error

func init() {
	// 配置MySQL连接参数，添加认证插件设置
	cfg := mysql.Config{
		User:   "root",
		Passwd: "root",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "gorm", //数据库名gorm
		Params: map[string]string{
			"charset":              "utf8mb4",
			"parseTime":            "True",
			"loc":                  "Local",
			"allowNativePasswords": "true",
		},
	}

	// 使用sqlx连接数据库
	sqxDB, sqxerr = sqlx.Connect("mysql", cfg.FormatDSN())
	if sqxerr != nil {
		panic(fmt.Sprintf("无法连接数据库: %v", sqxerr))
	}

}

type Employee struct {
	Id         int     `db:"id"`
	Name       string  `db:"name"`
	Department string  `db:"department"`
	Salary     float64 `db:"salary"`
}

type Book struct {
	Id     int
	Title  string
	Author string
	Price  int
}

func main() {

	// 取技术部的人员信息
	var employees []Employee
	sql := "select id, name, department, salary from employees where department = ?"
	sqxDB.Select(&employees, sql, "技术部")
	fmt.Println(employees)

	// 取工资最高的人员信息
	var topSalaryEmployee Employee
	sql = "select id, name, department, salary from employees order by salary desc limit 1"
	sqxDB.Get(&topSalaryEmployee, sql)
	fmt.Println(topSalaryEmployee)

	var books []Book
	sql = "select * from books where price > ?"
	sqxDB.Select(&books, sql, 50)
	fmt.Println(books)

}
