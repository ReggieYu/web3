package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

type Book struct {
	ID     int     `db:"id""`
	Title  string  `db:"title""`
	Author string  `db:"author"`
	Price  float64 `db:"price"`
}

func main() {
	//题目2：实现类型安全映射
	//假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
	//要求 ：
	//定义一个 Book 结构体，包含与 books 表对应的字段。
	//编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。
	db, err := sqlx.Connect("mysql", "root:123456@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err)
	}

	defer db.Close()
	var books []Book
	query := "select * from books where price > ?"
	err = db.Select(&books, query, 50)
	if err != nil {
		log.Fatalf("search error %v", err)
	}

	//print result
	fmt.Printf("search books %+v", books)
}
