package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Now you have %g programs", math.Nextafter(2, 3))
}
