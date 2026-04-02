<h1 align="center">Laporan Praktikum Modul 1</h1>
<p align="center">[Ferdyan Dhesta Firmansyah] - [109082500109]</p>

## Unguided

### 1. Minggu ini, mahasiswa Fakultas Informatika mendapatkan tugas dari mata kuliah matematika
diskrit untuk mempelajari kombinasi dan permutasi. Jonas salah seorang mahasiswa, iseng
untuk mengimplementasikannya ke dalam suatu program. Oleh karena itu bersediakah kalian
membantu Jonas? (tidak tentunya ya :p)
Masukan terdiri dari empat buah bilangan asli a, b, c, dan d yang dipisahkan oleh spasi,
dengan syarat a ≥ c dan b ≥ d.
Keluaran terdiri dari dua baris. Baris pertama adalah hasil permutasi dan kombinasi a
terhadap c, sedangkan baris kedua adalah hasil permutasi dan kombinasi b terhadap d.
Catatan: permutasi (P) dan kombinasi (C) dari n terhadap r (n ≥ r) dapat dihitung dengan
menggunakan persamaan berikut!

#### [soal1.go]

```go
package main

import "fmt"

func factorial(n int, hasil *int) {
	*hasil = 1
	for i := 1; i <= n; i++ {
		*hasil *= i
	}
}

func permutation(n, r int, hasil *int) {
	var nFact, nmrFact int
	factorial(n, &nFact)
	factorial(n-r, &nmrFact)
	*hasil = nFact / nmrFact
}

func combination(n, r int, hasil *int) {
	var pVal, rFact int
	permutation(n, r, &pVal) 
	factorial(r, &rFact)
	*hasil = pVal / rFact
}

func main() {
	var a, b, c, d int
	var resP1, resC1, resP2, resC2 int

	fmt.Scan(&a, &b, &c, &d)

	permutation(a, c, &resP1)
	combination(a, c, &resC1)
	fmt.Println(resP1, resC1)

	permutation(b, d, &resP2)
	combination(b, d, &resC2)
	fmt.Println(resP2, resC2)
}
```

## Output Unguided :

#### Output

![Screenshot Output Unguided 1\_1](https://github.com/ferdyandhestafirmansyah-boop/109082500109_Ferdyan-Dhesta-Firmansyah/blob/main/Modul%204/Output/output-soal1.png)

#### Penjelasan

Jadi program ini membaca 4 angka yaitu a, b, c, dan d,program menghitung permutasi dan kombinasi menggunakan rumus,dijadikan factorial untuk menghitung faktorial angka.
Kemudian fungsi permutasi untuk menghitung P(n,r) dan fungsi combinasi untuk menghitung C(n,r).


### 2. Kompetisi pemrograman tingkat nasional berlangsung ketat. Setiap peserta diberikan 8 soal
yang harus dapat diselesaikan dalam waktu 5 jam saja. Peserta yang berhasil menyelesaikan
soal paling banyak dalam waktu paling singkat adalah pemenangnya.
Buat program gema yang mencari pemenang dari daftar peserta yang diberikan. Program
harus dibuat modular, yaitu dengan membuat prosedur hitungSkor yang mengembalikan total
soal dan total skor yang dikerjakan oleh seorang peserta, melalui parameter formal.
Pembacaan nama peserta dilakukan di program utama, sedangkan waktu pengerjaan dibaca
di dalam prosedur.
prosedure hitungSkor(in/out soal, skor : integer)
Setiap baris masukan dimulai dengan satu string nama peserta tersebut diikuti dengan adalah
8 integer yang menyatakan berapa lama (dalam menit) peserta tersebut menyelesaikan soal.
Jika tidak berhasil atau tidak mengirimkan jawaban maka otomatis dianggap menyelesaikan
dalam waktu 5 jam 1 menit (301 menit).
Satu baris keluaran berisi nama pemenang, jumlah soal yang diselesaikan, dan nilai yang
diperoleh. Nilai adalah total waktu yang dibutuhkan untuk menyelesaikan soal yang berhasil
diselesaikan.
#### [soal2.go]

```go

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
```

## Output Unguided :

#### Output

![Screenshot Output Unguided 1\_1](https://github.com/ferdyandhestafirmansyah-boop/109082500109_Ferdyan-Dhesta-Firmansyah/blob/main/Modul%204/Output/output-soal2.png)

#### Penjelasan
Program ini membaca nama peserta satu per satu sampai bertemu kata "Selesai".Untuk setiap peserta, program membaca 8 angka waktu. Jika angkanya di bawah 301, itu dihitung sebagai soal benar dan waktunya dijumlahkan.Program membandingkan semuanya: siapa yang soal benarnya paling banyak akan menjadi pemenang dan dicetak di akhir.
