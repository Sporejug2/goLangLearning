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

/*

etominac@etominac-ntb MINGW64 ~/Software/goLang/goLangLearning/Section83 (main)
$ go get github.com/Sporejug2/puppy@b6d4d5e
go: downloading github.com/Sporejug2/puppy v0.0.0-20240104130407-b6d4d5e42a2b
go: added github.com/Sporejug2/Dog v0.0.0-20240104125923-7b8204ba61c5
go: upgraded github.com/Sporejug2/puppy v0.0.0-20240104122133-c0093b9b9e78 => v0.0.0-20240104130407-b6d4d5e42a2b


*/

func main() {
	s1 := puppy.Bark()
	s2 := puppy.Barks()
	fmt.Println(s1, s2)

	s3 := puppy.BigBark()
	s4 := puppy.BigBarks()
	fmt.Println(s3)
	fmt.Println(s4)

	fmt.Println(puppy.Bark())
	fmt.Println(puppy.Barks())
}
