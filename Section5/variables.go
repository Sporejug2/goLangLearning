package main

import "fmt"

var c, python, java bool

var i, j int = 1, 2

const d int = 42

var k, x = 2, 2.0

func main() {

	var i int
	fmt.Println(i, c, python, java, d)

	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
	fmt.Printf("%T \t %T \t %T \t %T \t %T \t", k, x, c, python, java)

	// unutar funkcije moze se koristiti := short assignment statement
	// a := implicitan tip
}
