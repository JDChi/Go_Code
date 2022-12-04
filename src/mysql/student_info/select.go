package main

import (
	"Go_Code/src/mysql/global"
	"fmt"
	"github.com/jmoiron/sqlx"
)

const selectAll = `
   SELECT * FROM student_info
`

const selectDistinct = `
   SELECT DISTINCT department FROM student_info
`

const selectLimit = `
   SELECT number, name, id_number, major FROM student_info LIMIT 2
`

const selectOrder = `
   SELECT * FROM student_score ORDER BY score DESC
`

func selectTable() {
	var studentInfos []StudentInfo
	err := global.SqlxDB.Select(&studentInfos, selectAll)
	if err != nil {
		panic(err)
	}
	fmt.Printf("len of studentInfos = %d\n", len(studentInfos))
	for _, studentInfo := range studentInfos {
		fmt.Printf("user info = %v is male = %v\n", studentInfo, studentInfo.Sex == Male)
	}

	err = global.SqlxDB.Select(&studentInfos, selectDistinct)
	if err != nil {
		panic(err)
	}
	fmt.Printf("distinct len of studentInfos = %d\n", len(studentInfos))
	for _, studentInfo := range studentInfos {
		fmt.Printf("user info = %v\n", studentInfo)
	}
	err = global.SqlxDB.Select(&studentInfos, selectLimit)
	if err != nil {
		panic(err)
	}
	fmt.Printf("limit len of studentInfos = %d\n", len(studentInfos))
	for _, studentInfo := range studentInfos {
		fmt.Printf("user info = %v\n", studentInfo)
	}

	var studentScores []StudentScore
	err = global.SqlxDB.Select(&studentScores, selectOrder)
	if err != nil {
		panic(err)
	}
	fmt.Printf("order len of studentInfos = %d\n", len(studentScores))
	for _, studentScore := range studentScores {
		fmt.Printf("user info = %v\n", studentScore)
	}

	// using StructScan
	fmt.Printf("use struct scan\n")
	var studentScores1 []StudentScore
	rows, err := global.SqlxDB.Query(selectOrder)
	defer rows.Close()
	if err != nil {
		panic(err)
	}
	// all rows
	err = sqlx.StructScan(rows, &studentScores1)
	if err != nil {
		panic(err)
	}
	for _, studentScore := range studentScores1 {
		fmt.Printf("user info = %v\n", studentScore)
	}
}
