package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	fmt.Println("this is from the init function")
}

func main() {

	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(400)
	fmt.Printf("The value of x is %v", x)

	switch {
	case x > 0 && x <= 100:
		fmt.Println("x is less than 100", x)
	case x >= 101 && x <= 200:
		fmt.Println("x is more than 100 and less than 200", x)
	case x > 201 && x <= 250:
		fmt.Println("x is more than 200 and less than 250", x)
	default:
		fmt.Println(" x is more than 250", x)

	}
}
