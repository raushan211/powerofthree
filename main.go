package main

import (
	"fmt"
	"math"
)

func main() {
	n := 27

	fmt.Println(isPowerOfThree(n))

}
func isPowerOfThree(n int) bool {

	for x := 0; x < 31; x++ {
		y := float64(x)
		if float64(n) == math.Pow(3, y) {

			return true

		}
	}
	return false
}
