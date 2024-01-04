package Section5

import "fmt"

const (
	// create a huge number shifting a 1 bit left 100 places
	// in other words the binary numbers that is 1 followed by 100 zeroes
	Big = 1 << 100
	// shift it right again 99 places, so we end up with 1 << 1, of 2
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

	fmt.Println("%d \t %b\n", 1, 1)
	fmt.Println("%d \t %b\n", 1<<1, 1<<1)
	fmt.Println("%d \t %b\n", 1<<2, 1<<2)
	fmt.Println("%d \t %b\n", 1<<3, 1<<3)
	fmt.Println("%d \t %b\n", 1<<4, 1<<4)
	fmt.Println("%d \t %b\n", 1<<5, 1<<5)
}
