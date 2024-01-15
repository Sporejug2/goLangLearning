package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	x := rand.Intn(400)
	fmt.Printf("The value of x is %v\t", x)

	if x <= 100 {
		fmt.Println("number is less than 100 ", x)
	} else if x >= 101 && x <= 200 {
		fmt.Println("number is more than 100 and less than 200", x)
	} else if x >= 201 && x <= 250 {
		fmt.Println("number is more than 201 and less than 250", x)
	} else {
		fmt.Println("this was more than 250")
	}

	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))

}
