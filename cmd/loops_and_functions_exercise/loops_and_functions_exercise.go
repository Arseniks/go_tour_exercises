package main

import "fmt"

func Sqrt(x float64) float64 {
	z := x
	for i := 0; i < 100; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
