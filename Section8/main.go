package main

import (
	"fmt"
	"mymodule/puppy"
)

// go mod init mymodule
// go mod tidy
func main() {
	s1 := puppy.Bark()
	s2 := puppy.Barks()
	fmt.Println(s1, s2)
}
