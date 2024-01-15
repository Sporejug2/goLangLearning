package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("this is the first statement to run")
	fmt.Println("this is the second statement to run")
	x := 40
	y := 5
	fmt.Printf(" x=%v \n y=%v\n", x, y)

	if x < 42 {
		fmt.Println("Less then the meaning of life")
	}

	if x < 42 {
		fmt.Println("Less then the meaning of life")
	} else {
		fmt.Println("equal to, or greater than, the meaning of life")
	}

	if x < 42 {
		fmt.Println("Less then meaning of life")
	} else if x == 42 {
		fmt.Println("equal to the meaning of life")
	} else {
		fmt.Println("greater than the meaning of life")
	}

	//Logical operators

	if (x < 42) && (y < 42) {
		fmt.Println("both are less than the meaning of life")
	}

	if x > 30 || x < 42 {
		fmt.Println("x is getting close to the meaning of life")
	}

	if x != 42 && y != 10 {
		fmt.Println("x is not 42 & y is not 10")
	}

	// statement idium
	// limitira scope

	if z := 2 * rand.Intn(x); z >= x {
		fmt.Printf("z is %v and that is greater than or equal x which is %v\n", z, x)
	} else {
		fmt.Printf("z is %v and that is Less than x wich is %v\n", z, x)
	}

	// comma ok idiom

	/*func offset(tz string) int {
		if seconds, ok := timeZone[tz]; ok {
			return seconds
	}
		Println("Unknown time zone:", tz)
		retrun 0
	}*/

	switch {
	case x < 42:
		fmt.Println("x is less than 42")
	case x == 42:
		fmt.Println("x is equal to 42")
	case x > 42:
		fmt.Println("x is greater than 42")
	default:
		fmt.Println("This is the default case for x")
	}

	switch x {
	case 40:

		fmt.Println(" x is 40")
	case 41:

		fmt.Println("x is 41")
	case 42:

		fmt.Println("x is 42")
	case 43:
		fmt.Println("x is 43")
	default:
		fmt.Println("this is the default case for x")
	}

	switch x {
	case 40:
		fmt.Println("x is 40")
		fallthrough
	case 41:
		fmt.Println("printed because of fallthrough: x is 41")
	case 42:
		fmt.Println("printed because of fallthrough: x is 42")
	case 43:
		fmt.Println("printed because of fallthrough: x is 43")
	default:
		fmt.Println("printed because of fallthrugh: this is the default case for x")
	}

	switch x {
	case 40:
		fmt.Println("x is 40")
		fallthrough
	case 41:
		fmt.Println("printed because of all of the fallthrought statements: x is 41")
		fallthrough
	case 42:
		fmt.Println("printed because of all of the fallthrought statements: x is 42")
		fallthrough
	case 43:
		fmt.Println("printed because of all of the fallthrought statements: x is 43")
		fallthrough
	default:
		fmt.Println("printed because of all of the fallthrought statements: this is the default case for x")
	}

	// Select statement choses which of a set of possible send or recieve operations will proceed.
	// It looks similar to a "switch" statement but with the cases all referring to communication operations.
	// https://go.dev/ref/spec#Select_statement
	select {
	case v1 := <-ch1:
		fmt.Println("Value from channel 1", v1)
	case v2 := <-ch2:
		fmt.Println("value from channel 2", v2)
	}

}
