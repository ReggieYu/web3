package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	//假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
	//要求 ：
	//编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1)/test?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	err = transferMoney(db, 1, 2, 100.00)
	if err != nil {
		log.Printf("transfer money failed: %v", err)
	} else {
		log.Printf("transfer money success")
	}
}

func transferMoney(db *sql.DB, fromAccountId int, toAccountId int, amount float64) error {
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("begin transaction failed %w", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	var balance float64
	err = tx.QueryRow("select balance from accounts where id = ? for update", fromAccountId).Scan(&balance)
	if err != nil {
		return fmt.Errorf("query error %w", err)
	}

	if balance < amount {
		return errors.New("balance is not enough")
	}

	_, err = tx.Exec("update accounts set balance = balance - ? where id = ?", amount, fromAccountId)
	if err != nil {
		return fmt.Errorf("sustract balance failed %w", err)
	}

	_, err = tx.Exec("update accounts set balance = balance + ? where id = ?", amount, toAccountId)
	if err != nil {
		return fmt.Errorf("plus balance failed %w", err)
	}

	_, err = tx.Exec("insert into transactions (from_account_id, to_account_id, amount) values (?, ?, ?)", fromAccountId, toAccountId, amount)
	if err != nil {
		return fmt.Errorf("insert transaction failed %w", err)
	}

	return tx.Commit()
}
