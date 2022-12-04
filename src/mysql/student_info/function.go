package main

import (
	"Go_Code/src/mysql/global"
	"fmt"
)

const aggregateFunMax = `
   SELECT MAX(score) FROM student_score WHERE subject = 'English'
`

const aggregateFunMulti = `
   SELECT MAX(score) AS max, MIN(score) AS min, AVG(score) AS avg, COUNT(*) AS count FROM student_score
`

func function() {
	var score int
	// using Get will return single result
	err := global.SqlxDB.Get(&score, aggregateFunMax)
	if err != nil {
		panic(err)
	}
	fmt.Printf("aggregateFunMax score = %d\n", score)

	// aggregate multi results
	type AggregateResult struct {
		Max   int     `db:"max"`
		Min   int     `db:"min"`
		AVG   float32 `db:"avg"`
		Count float32 `db:"count"`
	}
	var aggregateMultiResults []AggregateResult
	err = global.SqlxDB.Select(&aggregateMultiResults, aggregateFunMulti)
	if err != nil {
		panic(err)
	}
	fmt.Printf("aggregateMultiResults = %v\n", aggregateMultiResults)

}
