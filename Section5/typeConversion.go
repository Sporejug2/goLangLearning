package Section5

import (
	"fmt"
	"math"
)

const Pi float64 = 3.14

func main() {
	var i int = 42
	var f float64 = float64(i)
	var u uint = uint(f)

	// simpley
	var y = 42
	var k = float64(y)
	var l = uint(f)
	fmt.Println(y, k, l, u)

	var n, c int = 3, 4
	var m float64 = math.Sqrt(float64(c*c + n*n))
	var z uint = uint(m)
	fmt.Println(c, m, z)

	var v1 int = 42
	v := 42
	fmt.Printf("v is type of %T\n", v)
}
