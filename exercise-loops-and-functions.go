package main

import (
    "fmt"
	"math"
)

func Sqrt(x float64) float64 {
	var z float64 = 1
	for i := 0; i < 10; i++ {
		z = z - (math.Pow(z, 2) - x) / (2*z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
