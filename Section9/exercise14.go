package main

import "fmt"

var a1 = 10

const b1 = 3.14

var x int = 41

const y = 40

func main() {

	c := "decleration"

	fmt.Println(a1, b1, c)

	fmt.Printf("the value of x is %v and the type of x is %T\n", x, x)
	fmt.Printf("the value of y is %v and the type of y is %T\n", y, y)
	z := 42
	fmt.Printf("the value of z is %v and the type of z is %T\n", z, z)

}
