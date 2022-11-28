package main

import (
	"fmt"
	"strconv"
)

func main() {
	/////////////////// type conversion
	var numString = "111"
	// string to int
	numInt, err := strconv.Atoi(numString)
	if err != nil {
		fmt.Printf("atoi err = %v\n", err)
	} else {
		fmt.Printf("numInt = %d\n", numInt)
	}
	// int to string
	numString = strconv.FormatInt(int64(numInt), 10)
	fmt.Printf("numString = %s\n", numString)
}
