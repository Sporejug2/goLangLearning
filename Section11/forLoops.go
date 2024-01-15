package main

import "fmt"

func main() {

	x := 40
	y := 5
	fmt.Printf(" x=%v \n y=%v\n ", x, y)

	for i := 0; i < 5; i++ {
		fmt.Printf("y is %v \t\t\t second for loop\n", y)
		y++
	}

	for y < 10 {
		fmt.Printf("y is %v \t\t\t second for loop\n", y)
		y++
	}

	//breaks
	for {
		fmt.Printf("y is %v \t\t third for loop\n", y)
		if y > 20 {
			break
		}
		y++
	}

	for i := 0; i < 20; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println("counting even numbers: ", i)
	}

}
