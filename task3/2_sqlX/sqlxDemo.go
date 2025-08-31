package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Employee struct {
	Id         int    `db:"id"`
	Name       string `db:"name"`
	Department string `db:"department"`
	Salary     int    `db:"salary"`
}

func main() {
	//题目1：使用SQL扩展库进行查询
	//假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
	//要求 ：
	//编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
	//编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
	db, err := sqlx.Connect("mysql", "root:123456@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err)
	}

	defer db.Close()
	var techEmp []Employee
	err = db.Select(&techEmp, "select id, name, department, salary from employees where department=?", "技术部")
	if err != nil {
		panic(err)
	}
	fmt.Printf("search tech employee: %+v\n\n", techEmp)

	//search highest salary employee
	var highestEmp []Employee
	err = db.Select(&highestEmp, "select * from employees order by salary desc limit 1")
	if err != nil {
		panic(err)
	}

	fmt.Printf("highest employess %+v", highestEmp)
}
