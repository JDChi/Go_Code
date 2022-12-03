package main

import (
	"Go_Code/src/mysql/global"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
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

	// insert
	// insert student info
	result, err := global.SqlxDB.Exec(insertStudentInfo)
	if err != nil {
		panic(err)
	}
	rows, err = result.RowsAffected()
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
