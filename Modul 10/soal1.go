package 

import "fmt"


const NMAX = 1000

type arrKelinci [NMAX]float64

func main() {
	var berat arrKelinci
	var n int
	var min, max float64

	fmt.Print("Masukkan jumlah anak kelinci (N): ")
	fmt.Scan(&n)

	fmt.Printf("Masukkan berat %d anak kelinci:\n", n)
	for i := 0; i < n; i++ {
		fmt.Scan(&berat[i])
	}

	if n > 0 {
		min = berat[0]
		max = berat[0]

		for i := 1; i < n; i++ {
			if berat[i] < min {
				min = berat[i]
			}
			if berat[i] > max {
				max = berat[i]
			}
		}

		fmt.Println("Berat kelinci terkecil:", min)
		fmt.Println("Berat kelinci terbesar:", max)
	}
}