package main

import (
	"Go_Code/src/mysql/global"
	"database/sql"
	"fmt"
)

const leftJoin = `
   SELECT student_info.number, name, major, subject, score FROM student_info LEFT JOIN student_score ON student_info.number = student_score.number
`

func join() {
	type LeftJoinInfo struct {
		Number  int            `db:"number"`
		Name    string         `db:"name"`
		Major   string         `db:"major"`
		Subject sql.NullString `db:"subject"`
		Score   sql.NullInt16  `db:"score"`
	}

	var leftJoinInfos []LeftJoinInfo

	err := global.SqlxDB.Select(&leftJoinInfos, leftJoin)
	if err != nil {
		panic(err)
	}

	for _, value := range leftJoinInfos {

		fmt.Printf("left join info number = %v, name = %v, major = %v, subject = %s score = %d\n ", value.Number, value.Name, value.Major, value.Subject.String, value.Score.Int16)
	}

}
