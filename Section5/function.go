package Section5

import "fmt"

func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return x, y
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	// naked return
	return
}

func main() {
	fmt.Println(add(5, 4))
	sayHello()
	a := add(100, 200)
	fmt.Println(a)
	fmt.Println(add(42, 13))

	b, c := swap("Hello", "Hello")
	fmt.Println(b, c)
	fmt.Println("--------")
	fmt.Println(split(18))
}

func sayHello() {
	fmt.Println("Hello")
}
