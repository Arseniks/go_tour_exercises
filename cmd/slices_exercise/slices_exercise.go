package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	a := make([][]uint8, dy)

	for y := range a {
		a[y] = make([]uint8, dx)
		for x := range a[y] {
			a[y][x] = uint8(x ^ y)
		}
	}
	return a
}

func main() {
	pic.Show(Pic)
}
