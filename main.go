package main

// rumus luas lingkaran
// L = phi * r * r
// rumus k
// K = 2 * phi * r

import "fmt"

func keliling(phi int, r int) int {
	result := 2 * phi * r
	return result
}

func luas(phi int, r int) int {
	result := phi * r * r
	return result
}
func main() {
	fmt.Println(keliling(3, 3))
	fmt.Println(luas(3, 3))

	fmt.Println(3.14 * 3)
}
