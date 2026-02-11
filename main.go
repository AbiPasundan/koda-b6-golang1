package main

import "fmt"

func keliling(r int, phi int) int {
	// Menggunakan Diameter (\(d\)): \(K=\pi \times d\)
	// Menggunakan Jari-jari (\(r\)): \(K=2\times \pi \times r\)(Catatan: \(d=2\times r\))
	// d = 2 * r
	// phi := 3
	diameter := 2 * phi * r
	// jarijari := 2 * phi * r
	return diameter
}

func luas(r int, phi int) int {
	result := phi * (r * r)
	return result
}
func main() {
	fmt.Println(keliling(14, 3))
	fmt.Println(luas(14, 3))
}
