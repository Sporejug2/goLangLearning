package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(10)
	y := rand.Intn(10)
	fmt.Printf("x and y are %v and %v\t", x, y)

	if x < 4 && y < 4 {
		fmt.Println("x and y are both less than ", 4)
	} else if x > 6 && y > 6 {
		fmt.Println("x and y are both higher than 6")
	} else if x >= 4 && x <= 6 {
		fmt.Println(" x is higher than 4 and less than 6")
	} else if y != 6 {
		fmt.Println("y is not 5")
	} else {
		fmt.Println("Non of the previous cases where met")
	}

	switch {
	case x < 4 && y < 4:
		fmt.Println("x and y are both less than ", 4)
	case x > 6 && y > 6:
		fmt.Println("x and y are both higher than 6")
	case x >= 4 && x <= 6:
		fmt.Println(" x is higher than 4 and less than 6")
	case y != 6:
		fmt.Println("y is not 5")
	default:
		fmt.Println("Non of the previous cases where met")
	}

}
