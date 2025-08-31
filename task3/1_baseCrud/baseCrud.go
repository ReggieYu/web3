package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Student struct {
	Id    int
	Name  string
	Age   int
	Grade string
}

func main() {
	//题目1：基本CRUD操作
	//假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
	//要求 ：
	//编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
	//编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
	//编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
	//编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	//insert student
	//insertStudent(db, "张三", 20, "三年级")
	//res := searchStudent(db, 15)
	//for _, s := range res {
	//	fmt.Printf("query student: %+v\n", s)
	//}
	//updateStudent(db, "四年级", "张三")
	deleteStudent(db, 15)
}

func deleteStudent(db *sql.DB, age int) {
	_, err := db.Exec("delete from students where age < ?", age)
	if err != nil {
		log.Printf("delete student failed: %v", err)
	} else {
		log.Println("delete student success")
	}
}

func updateStudent(db *sql.DB, grade string, name string) {
	_, err := db.Exec("update students set grade = ? where name = ?", grade, name)
	if err != nil {
		log.Printf("update student failed: %v", err)
	} else {
		log.Println("update student success")
	}
}

func searchStudent(db *sql.DB, age int) []Student {
	rows, err := db.Query("select * from students where age > ?", age)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	var students []Student
	for rows.Next() {
		var s Student
		err := rows.Scan(&s.Id, &s.Name, &s.Age, &s.Grade)
		if err != nil {
			log.Fatal(err)
		}
		students = append(students, s)
	}

	return students
}

func insertStudent(db *sql.DB, name string, age int, grade string) {
	stmt, err := db.Prepare("insert into students(name, age, grade) values (?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	res, err := stmt.Exec(name, age, grade)
	if err != nil {
		log.Fatal(err)
	}

	id, _ := res.LastInsertId()
	fmt.Printf("insert success, id: %d \n\n", id)
}
