package main

import (
	"Go_Code/src/mysql/global"
	"fmt"
)

const createStudentInfoTable = `
   CREATE TABLE IF NOT EXISTS student_info (
       number INT PRIMARY KEY,
       name VARCHAR(5),
       sex ENUM('male','female'),
       id_number CHAR(18),
       department VARCHAR(30),
       major LONGTEXT,
       enrollment_time DATE,
       UNIQUE KEY (id_number)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
`

const createStudentScoreTable = `
   CREATE TABLE IF NOT EXISTS student_score(
       number INT,
       subject VARCHAR(30),
       score TINYINT,
       PRIMARY KEY (number, subject),
       CONSTRAINT FOREIGN KEY(number) REFERENCES student_info(number)      
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 
`

func createTable() {
	rows, err := global.SqlxDB.MustExec(createStudentInfoTable).RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Printf("createStudentInfoTable rows = %d\n", rows)

	rows, err = global.SqlxDB.MustExec(createStudentScoreTable).RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Printf("createStudentScoreTable rows = %d\n", rows)
}
