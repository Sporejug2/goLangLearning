package Section5

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	/*

		bool
		string
		int int8 int 16 int32 int64
		uint uint8 uint16 uint32 uint64 uintptr
		// unit ide 0 -255
		// int od -128 - 127
		byte // alias for uint8
		rune - char
		float32, float64
		complex64 complex128
	*/

}
