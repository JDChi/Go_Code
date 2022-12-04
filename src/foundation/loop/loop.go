package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	// way 1
	for i := 0; i < len(nums); i++ {
		fmt.Printf("num = %d ", nums[i])
	}
	fmt.Printf("\n")
	// way 2
	for index, num := range nums {
		fmt.Printf("index = %d, num = %d ", index, num)
	}
	fmt.Printf("\n")
	/////////////
	// in for-range, we can ignore key or value to make a better performance
	for _, num := range nums {
		fmt.Printf("num = %d ", num)
	}
	// for-range in map, it will return random result
	kvs := map[string]string{
		"a": "a",
		"b": "b",
		"c": "c",
	}
	fmt.Printf("\n")
	for key, value := range kvs {
		fmt.Printf("key = %s value = %s ", key, value)
	}
	// key = c value = c key = a value = a key = b value = b
	fmt.Printf("\n")

	// for-range in string
	for k, v := range "hello" {
		fmt.Printf("key = %v value = %c ", k, v)
	}
	fmt.Printf("\n")

	// as the value in for-range is "copy" of the value,
	// that means we can't change the value actually
	tmp := []struct {
		int
		string
	}{
		{1, "a"},
		{2, "b"},
	}
	for _, v := range tmp {
		if v.int == 1 {
			v.string = "c"
		}
		fmt.Printf("v = %v ", v)
	}
	fmt.Printf("\n")
	fmt.Printf("actual tmp = %v", tmp)
	fmt.Printf("\n")
	// if we want to change the actual value, we can try to use the index
	for index, v := range tmp {
		if index == 0 {
			tmp[index].string = "c"
		}
		// when we print this, we can see the difference between the actual value and the copy value
		fmt.Printf("v = %v ", v)
	}
	fmt.Printf("\n")
	fmt.Printf("actual tmp = %v", tmp)
}
