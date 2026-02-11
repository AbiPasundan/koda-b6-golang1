package main

// rumus luas lingkaran
// L = phi * r * r
// rumus k
// K = 2 * phi * r

import "fmt"

func keliling(phi float32, r int) int {
	// result float32 := 2 * phi * r
	return int(2 * phi * float32(r))
}

func luas(phi int, r int) int {
	// result := phi * r * r
	return int(phi * r * r)
}
func main() {
	fmt.Println(keliling(3.14, 3))
	fmt.Println(luas(3, 3))

	fmt.Println(3.14 * 3)

}
