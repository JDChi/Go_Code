package main

import (
	"Go_Code/src/mysql/global"
	"fmt"
)

const groupByAvgScore = `
   SELECT subject, AVG(score) AS avg FROM student_score GROUP BY subject
`

const groupByAvgScoreHaving = `
   SELECT subject, AVG(score) AS avg FROM student_score GROUP BY subject HAVING AVG(score) > 73
`

const groupByMultiColumn = `
   SELECT department, major, COUNT(*) AS count FROM student_info GROUP BY department, major
`

func groupBy() {

	type GroupByAvgScore struct {
		Subject string  `db:"subject"`
		Avg     float32 `db:"avg"`
	}

	var groupByAvgScores []GroupByAvgScore

	err := global.SqlxDB.Select(&groupByAvgScores, groupByAvgScore)
	if err != nil {
		panic(err)
	}
	fmt.Printf("groupByAvgScore len of studentInfos = %d\n", len(groupByAvgScores))
	for _, groupByAvgScore := range groupByAvgScores {
		fmt.Printf("result info = %v\n", groupByAvgScore)
	}

	// use HAVING to filter
	err = global.SqlxDB.Select(&groupByAvgScores, groupByAvgScoreHaving)
	if err != nil {
		panic(err)
	}
	fmt.Printf("groupByAvgScoreHaving len of studentInfos = %d\n", len(groupByAvgScores))
	for _, groupByAvgScore := range groupByAvgScores {
		fmt.Printf("result info = %v\n", groupByAvgScore)
	}
	// in some way, we can use HAVING instead of WHERE,
	// but we should know that WHERE can't handle aggregate function which occur in HAVING
	// like this: `SELECT subject, AVG(score), FROM student_score WHERE MAX(score) > 90 GROUP BY subject`, this is wrong
	// as WHERE is used to filter every single record, obviously, the aggregate function is not used to adjust single record

	////////////////////////
	// if a query statement contains nested groups, the aggregate function will act on the last group
	// in this case, it will group by the department firstly, and group by major, then the SUM(*) will act on major
	type GroupByMultiColumn struct {
		Department string `db:"department"`
		Major      string `db:"major"`
		Count      int    `db:"count"`
	}
	var groupByMultiColumns []GroupByMultiColumn
	err = global.SqlxDB.Select(&groupByMultiColumns, groupByMultiColumn)
	if err != nil {
		panic(err)
	}
	fmt.Printf("groupByMultiColumns len of studentInfos = %d\n", len(groupByMultiColumns))
	for _, groupByMultiColumn := range groupByMultiColumns {
		fmt.Printf("result info = %v\n", groupByMultiColumn)
	}

}
