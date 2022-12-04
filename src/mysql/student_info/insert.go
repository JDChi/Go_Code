package main

import (
	"Go_Code/src/mysql/global"
	"fmt"
)

// use IGNORE to avoid insert duplicate
const insertStudentInfo = `
   INSERT IGNORE INTO student_info (number, name, sex, id_number, department, major, enrollment_time) 
   VALUES 
       (20220101, 'Jack', 'male', 1000000000000001, 'IT Academy','Computer Science and Engineering', '2022-9-1'),
       (20220102, 'Bob', 'male', 1000000000000002, 'IT Academy','Computer Science and Engineering', '2022-9-1'),
       (20220103, 'Jane', 'female', 1000000000000003, 'IT Academy','Software Engineering', '2022-9-1'),
       (20220104, 'Rock', 'male', 1000000000000004, 'IT Academy','Software Engineering', '2022-9-1'),
       (20220105, 'Alice', 'female', 1000000000000005, 'Aerospace Institute','Aircraft Design', '2022-9-1'),
       (20220106, 'Kobe', 'male', 1000000000000006, 'Aerospace Institute','Digital Information', '2022-9-1')   
       `

const insertStudentScore = `
   INSERT IGNORE INTO student_score (number, subject, score)
   VALUES 
       (20220101, 'English', 78),
       (20220101, 'Physics', 88),
       (20220102, 'English', 100),
       (20220102, 'Physics', 98),
       (20220103, 'English', 59),
       (20220103, 'Physics', 61),
       (20220104, 'English', 55),
       (20220104, 'Physics', 46)
`

func insert() {
	result, err := global.SqlxDB.Exec(insertStudentInfo)
	if err != nil {
		panic(err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Printf("insertStudentInfo rows = %d\n", rows)
	// insert student scores
	result, err = global.SqlxDB.Exec(insertStudentScore)
	if err != nil {
		panic(err)
	}
	rows, err = result.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Printf("insertStudentScore rows = %d\n", rows)
}
