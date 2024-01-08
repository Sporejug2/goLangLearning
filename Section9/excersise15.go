package main

import (
	"fmt"
	"github.com/Sporejug2/puppy"
)

func main() {

	fmt.Println("Hello world")

	s1 := puppy.Bark()
	s2 := puppy.Barks()
	s3 := puppy.BigBark()
	s4 := puppy.BigBarks()

	fmt.Println(s1 + "\n" + s2 + "\n" + s3 + "\n" + s4)
	puppy.From13()
	// go build excersise15.go
	// go build ./...
	// go build .
	// ./excersise15.go
	//$ go get github.com/Sporejug2/puppy@v1.3.0

	/*

			bildanje na win, mac, linux
		GOOS=windows go build excersise15.go
		GOOS=darwing go build .
		GOOS=linux go build .

	*/

	cat.Meow()

}
