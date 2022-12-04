package main

import (
	"Go_Code/src/mysql/global"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type sex string

const (
	Male   sex = "male"
	Female sex = "female"
)

type StudentInfo struct {
	Number         int    `db:"number"`
	Name           string `db:"name"`
	Sex            sex    `db:"sex"`
	IdNumber       string `db:"id_number"`
	Department     string `db:"department"`
	Major          string `db:"major"`
	EnrollmentTime string `db:"enrollment_time"`
}

type StudentScore struct {
	Number  int    `db:"number"`
	Subject string `db:"subject"`
	Score   int    `db:"score"`
}

func init() {
	s := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s", "root", "123456", "127.0.0.1:3306", "go_code", "utf8")
	var err error

	global.SqlxDB, err = sqlx.Connect("mysql", s)
	if err != nil {
		panic(err)
	}
}

func main() {
	// create table
	createTable()

	// insert
	insert()

	// select
	selectTable()

	// function
	function()

	// group by
	groupBy()
}
