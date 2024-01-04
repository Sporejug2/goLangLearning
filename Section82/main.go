package main

import (
	"fmt"
	"github.com/Sporejug2/puppy"
)

// go mod init mymodule
// go mod tidy

//$ go get github.com/Sporejug2/puppy
//go: downloading github.com/Sporejug2/puppy v0.0.0-20240104122133-c0093b9b9e78
//go: added github.com/Sporejug2/puppy v0.0.0-20240104122133-c0093b9b9e78

func main() {
	s1 := puppy.Bark()
	s2 := puppy.Barks()
	fmt.Println(s1, s2)

	fmt.Println(puppy.Bark())
	fmt.Println(puppy.Barks())
}
