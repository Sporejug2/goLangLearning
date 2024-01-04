package Section5

import (
	"fmt"
)

//var x = 42

func main() {

	fmt.Println(x)

	printMe()

	y := 43
	fmt.Println(y)

	p1 := person{
		"James",
	}

	p1.sayHello()

	x := 32
	fmt.Println(x)

	printMe()

	if z := 82; z > 50 {
		fmt.Println(z)
	}

	futh

	futherexplorer.Fascinating()

	futherexplorer.fascinating()

	futherexplorer.planetSpeed
}

func printMe() {
	fmt.Println(x)
}

type person struct {
	first string
}

func (p person) sayHello() {
	fmt.Println("Hi my name is ", p.first)
}
