package main

import "fmt"

type arrBalita [100]float64

func hitungMinMax(arrBerat arrBalita, n int, bMin, bMax *float64) {
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]

	for i := 1; i < n; i++ {
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]
		}
		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
	}
}

func rerata(arrBerat arrBalita, n int) float64 {
	var sum float64
	for i := 0; i < n; i++ {
		sum += arrBerat[i]
	}
	return sum / float64(n)
}

func main() {
	var n int
	var berat arrBalita
	var min, max float64

	fmt.Print("Masukan banyak data berat balita : ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("Masukan berat balita ke-%d: ", i+1)
		fmt.Scan(&berat[i])
	}

	if n > 0 {
		hitungMinMax(berat, n, &min, &max)
		rata := rerata(berat, n)

		fmt.Printf("Berat balita minimum: %.2f kg\n", min)
		fmt.Printf("Berat balita maksimum: %.2f kg\n", max)
		fmt.Printf("Rerata berat balita: %.2f kg\n", rata)
	}
}