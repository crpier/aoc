package go_tour_8

import (
	"fmt"
	"math"
)

const delta float64 = 0.000000000000001

func Sqrt(x float64) float64 {
	var lastResult float64
	var z float64 = 1
	for {
		z -= (z*z - x) / (2 * z)
		if math.Abs(z-lastResult) < delta {
			break
		}
		lastResult = z
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
  fmt.Println(math.Sqrt(2))
}
