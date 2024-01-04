package Section5

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(11))
	fmt.Println(rand.Intn(14))
	fmt.Println(rand.Intn(7))

	fmt.Printf("Now you have %g problems .\n", math.Sqrt(97))
	fmt.Println(math.Pi)
}
