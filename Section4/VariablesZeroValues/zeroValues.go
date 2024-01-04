package main

import "fmt"

func main(){

	a := 42
	fmt.Println(a)

	b, c, d, f := 0, 1, 2, "happiness"
	fmt.Println(b, c, d, f)
	
	var h int = 44
	fmt.Println(h)

	adams := 42

	fmt.Println("42 as binary, %b \n", adams)
	fmt.Println("42 as hexadecimal, %x \n", adams)

	//print these values as both binary & hexadecimal
	a, b, c, d, e := 0, 1, 2, 3, 4
// heksadecimal
	fmt.Printf(" %v \t %b \t %#x \n", a, a, a)
	fmt.Printf(" %v \t %b \t %#x \n", b, b, b)
	fmt.Printf(" %v \t %b \t %#x \n", c, c, c)
	fmt.Printf(" %v \t %b \t %#x \n", d, d, d)
	fmt.Printf(" %v \t %b \t %#x \n", e, e, e)

// numeral system: decimal, binary, & hexadecimal

	// conversion nije type casting

	//y := 42
	z := 42.0

	var m float32 = 43.742
	fmt.Printf("%v of type %T \n", m ,m)
	fmt.Printf("%v of type %T \n", z, z)

	z = float64(m)
	fmt.Printf("%v of type %T \n", z, z)

	
}