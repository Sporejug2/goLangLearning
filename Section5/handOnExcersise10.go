package main

import "fmt"

var y int

func main() {

	fmt.Println(y)

	s := 42
	fmt.Println(s)

	a, b := 43, "Happiness"
	fmt.Println(a, b)

	var c float32 = 42.42
	fmt.Printf("%v is of this type %T\n", c, c)

	e, f, _ := 45, 46, 47
	fmt.Println(e, f)

}

/*
	var for zero value
	short declaration operator
	multiple initialization
	var when specificity is required
	blank identifier
*/
