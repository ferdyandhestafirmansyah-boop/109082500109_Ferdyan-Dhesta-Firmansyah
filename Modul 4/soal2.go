package main

import "fmt"

func hitungSkor(soal *int, skor *int) {
	*soal = 0
	*skor = 0
	var waktu int

	for i := 0; i < 8; i++ {
		fmt.Scan(&waktu)
		if waktu < 301 {
			*soal++          
			*skor += waktu   
		}
	}
}

func main() {
	var nama string

	var pemenangNama string
	pemenangSoal := -1       
	pemenangSkor := 9999999  

	for {
		fmt.Scan(&nama)

		if nama == "Selesai" {
			break
		}

		var soalPeserta, skorPeserta int

		hitungSkor(&soalPeserta, &skorPeserta)

		if soalPeserta > pemenangSoal {
			pemenangNama = nama
			pemenangSoal = soalPeserta
			pemenangSkor = skorPeserta
		} else if soalPeserta == pemenangSoal {
			if skorPeserta < pemenangSkor {
				pemenangNama = nama
				pemenangSoal = soalPeserta
				pemenangSkor = skorPeserta
			}
		}
	}
	fmt.Println(pemenangNama, pemenangSoal, pemenangSkor)
}