package main

// rumus luas lingkaran
// L = phi * r * r
// rumus k
// K = 2 * phi * r

import "fmt"

func keliling(r int) int {
	// result float32 := 2 * phi * r
	return int(2 * 3.14 * float32(r))
}

func luas(r int) int {
	// result := phi * r * r
	return int(3.14 * float32(r) * float32(r))
}
func main() {
	// fmt.Println(keliling(3.14, 3))
	// fmt.Println(luas(3.14, 3))

	var input int
	fmt.Println("Masukan Jari-Jari Untuk Mengetahui Luas dan Keliling Lingkaran")
	fmt.Scanln(&input)
	fmt.Print("Hasil dari keliling ", input, " adalah ", keliling(3), "\n")
	fmt.Print("Hasil dari Luas ", input, " adalah ", luas(3), "\n")
}
