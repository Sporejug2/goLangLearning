package main

import "fmt"

func main() {

	xi := []int{42, 43, 44, 45, 46, 47}
	for i, v := range xi {
		fmt.Println("ranging over a slice", i, v)
	}

	// for range loop
	// data structure - map
	m := map[string]int{
		"James":      42,
		"MoneyPenny": 32,
	}
	for k, v := range m {
		fmt.Println("raging over a map", k, v)
	}

}
