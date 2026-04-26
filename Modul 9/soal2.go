package 

import (
	"fmt"
	"math"
)

const NMAX = 1000

type arrayInt [NMAX]int

func main() {
	var A arrayInt
	var n, x, idxHapus, cari, frekuensi int
	var sum, rerata, varians, selisih float64

	fmt.Print("Masukkan Jumlah elemen array (N): ")
	fmt.Scan(&n)

	fmt.Printf("Masukkan %d elemen nilai:\n", n)
	for i := 0; i < n; i++ {
		fmt.Scan(&A[i])
	}

	fmt.Print("\na. Keseluruhan isi dari array:\n")
	for i := 0; i < n; i++ {
		fmt.Print(A[i], " ")
	}
	fmt.Println()

	fmt.Print("\nb. Elemen-elemen array dengan indeks ganjil:\n")
	for i := 1; i < n; i += 2 {
		fmt.Print(A[i], " ")
	}
	fmt.Println()

	fmt.Print("\nc. Elemen-elemen array dengan indeks genap:\n")
	for i := 0; i < n; i += 2 {
		fmt.Print(A[i], " ")
	}
	fmt.Println()

	fmt.Print("\nMasukkan bilangan x (untuk indeks kelipatan): ")
	fmt.Scan(&x)
	
	fmt.Printf("d. Elemen-elemen array dengan indeks kelipatan %d:\n", x)
	for i := 0; i < n; i++ {
		if x != 0 && i%x == 0 {
			fmt.Print(A[i], " ")
		}
	}
	fmt.Println()

	fmt.Print("\nMasukkan indeks elemen yang ingin dihapus: ")
	fmt.Scan(&idxHapus)
	
	if idxHapus >= 0 && idxHapus < n {
		for i := idxHapus; i < n-1; i++ {
			A[i] = A[i+1]
		}
		n--
	}
	
	fmt.Printf("e. Keseluruhan isi dari array setelah indeks %d dihapus:\n", idxHapus)
	for i := 0; i < n; i++ {
		fmt.Print(A[i], " ")
	}
	fmt.Println()

	sum = 0
	for i := 0; i < n; i++ {
		sum += float64(A[i])
	}
	rerata = sum / float64(n)
	fmt.Printf("\nf. Rata-rata dari bilangan di dalam array:\n%.2f\n", rerata)

	varians = 0
	for i := 0; i < n; i++ {
		selisih = float64(A[i]) - rerata
		varians += selisih * selisih
	}
	varians = varians / float64(n)
	fmt.Printf("\ng. Standar deviasi bilangan di dalam array:\n%.2f\n", math.Sqrt(varians))

	fmt.Print("\nMasukkan bilangan yang ingin dihitung frekuensinya: ")
	fmt.Scan(&cari)
	
	frekuensi = 0
	for i := 0; i < n; i++ {
		if A[i] == cari {
			frekuensi++
		}
	}
	fmt.Printf("h. Frekuensi kemunculan bilangan %d di dalam array:\n%d\n", cari, frekuensi)
}