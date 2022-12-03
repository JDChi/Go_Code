package main

import (
	"Go_Code/src/mysql/global"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const _createFirstTable = `CREATE TABLE IF NOT EXISTS first_table (
    first_column INT,
    second_column VARCHAR(100)
    )
`

func init() {
	s := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s", "root", "123456", "127.0.0.1:3306", "go_code", "utf8")
	var err error

	global.SqlxDB, err = sqlx.Connect("mysql", s)
	if err != nil {
		panic(err)
	}
}

func main() {
	rows, err := global.SqlxDB.MustExec(_createFirstTable).RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Printf("create first table rows = %d", rows)
}
