package main

import "fmt"

const NMAX = 1000

type arrIkan [NMAX]float64
type arrWadah [NMAX]float64

func main() {
	var berat arrIkan
	var wadah arrWadah
	var x, y, numWadah int
	var totalBerat, rerata float64

	fmt.Print("Masukkan x (jumlah ikan) dan y (kapasitas wadah): ")
	fmt.Scan(&x, &y)

	fmt.Printf("Masukkan berat %d ikan:\n", x)
	for i := 0; i < x; i++ {
		fmt.Scan(&berat[i])
	}

	for i := 0; i < x; i++ {
		idx := i / y
		wadah[idx] += berat[i]
		totalBerat += berat[i]
	}

	if x%y == 0 {
		numWadah = x / y
	} else {
		numWadah = (x / y) + 1
	}

	fmt.Println("Total berat ikan di setiap wadah:")
	for i := 0; i < numWadah; i++ {
		fmt.Printf("%.2f ", wadah[i])
	}
	fmt.Println()

	if numWadah > 0 {
		rerata = totalBerat / float64(numWadah)
		fmt.Println("Berat rata-rata ikan di setiap wadah:")
		fmt.Printf("%.2f\n", rerata)
	}
}