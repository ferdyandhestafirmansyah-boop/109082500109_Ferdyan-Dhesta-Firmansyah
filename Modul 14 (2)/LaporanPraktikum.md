<h1 align="center">Laporan Praktikum Modul 1</h1>
<p align="center">[Ferdyan Dhesta Firmansyah] - [109082500109]</p>

## Unguided

### 1.Buatlah sebuah program yang digunakan untuk membaca data integer seperti contoh yang
diberikan di bawah ini, kemudian diurutkan (menggunakan metoda insertion sort), dan
memeriksa apakah data yang terurut berjarak sama terhadap data sebelumnya.
Masukan terdiri dari sekumpulan bilangan bulat yang diakhiri oleh bilangan negatif. Hanya
bilangan non negatif saja yang disimpan ke dalam array.
Keluaran terdiri dari dua baris. Baris pertama adalah isi dari array setelah dilakukan
pengurutan, sedangkan baris kedua adalah status jarak setiap bilangan yang ada di dalam
array. "Data berjarak x" atau "data berjarak tidak tetap".

#### [soal1.go]

```go
package main

import "fmt"

const NMAX = 1000000

var arr [NMAX]int

func main() {
	var n int = 0
	var val int

	for {
		fmt.Scan(&val)
		if val < 0 {
			break
		}
		arr[n] = val
		n++
	}

	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}

	for i := 0; i < n; i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(arr[i])
	}
	fmt.Println()

	if n > 1 {
		diff := arr[1] - arr[0]
		tetap := true
		for i := 1; i < n-1; i++ {
			if arr[i+1]-arr[i] != diff {
				tetap = false
				break
			}
		}

		if tetap {
			fmt.Printf("Data berjarak %d\n", diff)
		} else {
			fmt.Println("Data berjarak tidak tetap")
		}
	}
}
```

## Output Unguided :

#### Output

![Screenshot Output Unguided 1\_1](https://github.com/ferdyandhestafirmansyah-boop/109082500109_Ferdyan-Dhesta-Firmansyah/blob/main/Modul%2014/Output/output-soal1.png)

#### Penjelasan
Jadi program ini program yang digunakan untuk membaca data integer seperti contoh yang
diberikan di bawah ini, kemudian diurutkan,dan memeriksa apakah data yang terurut berjarak sama terhadap data sebelumnya.


### 2.Sebuah program perpustakaan digunakan untuk mengelola data buku di dalam suatu
perpustakaan. Misalnya terdefinisi struct dan array seperti berikut ini:
const nMax : integer = 7919
type Buku = <
id, judul, penulis, penerbit : string
eksemplar, tahun, rating : integer >
type DaftarBuku = array [ 1..nMax] of Buku
Pustaka : DaftarBuku
nPustaka: integer

Masukan terdiri dari beberapa baris. Baris pertama adalah bilangan bulat N yang
menyatakan banyaknya data buku yang ada di dalam perpustakaan. N baris berikutnya,
masing-masingnya adalah data buku sesuai dengan atribut atau field pada struct. Baris
terakhir adalah bilangan bulat yang menyatakan rating buku yang akan dicari.
Keluaran terdiri dari beberapa baris. Baris pertama adalah data buku terfavorit, baris kedua
adalah lima judul buku dengan rating tertinggi, selanjutnya baris terakhir adalah data buku
yang dicari sesuai rating yang diberikan pada masukan baris terakhir.

Lengkapi subprogram-subprogram dibawah ini, sesuai dengan I.S. dan F.S yang diberikan.
procedure DaftarkanBuku(in/out pustaka : DaftarBuku, n : integer)
{I.S. sejumlah n data buku telah siap para piranti masukan
F.S. n berisi sebuah nilai, dan pustaka berisi sejumlah n data buku}
procedure CetakTerfavorit(in pustaka : DaftarBuku, in n : integer)
{I.S. array pustaka berisi n buah data buku dan belum terurut
F.S. Tampilan data buku (judul, penulis, penerbit, tahun)
terfavorit, yaitu memiliki rating tertinggi}
procedure UrutBuku( in/out pustaka : DaftarBuku, n : integer )
{I.S. Array pustaka berisi n data buku
F.S. Array pustaka terurut menurun/mengecil terhadap rating.
Catatan: Gunakan metoda Insertion sort}
procedure Cetak5Terbaru( in pustaka : DaftarBuku, n integer )
{I.S. pustaka berisi n data buku yang sudah terurut menurut rating
F.S. Laporan 5 judul buku dengan rating tertinggi
Catatan: Isi pustaka mungkin saja kurang dari 5}
procedure CariBuku(in pustaka : DaftarBuku, n : integer, r : integer )
{I.S. pustaka berisi n data buku yang sudah terurut menurut rating
F.S. Laporan salah satu buku (judul, penulis, penerbit, tahun,
eksemplar, rating) dengan rating yang diberikan. Jika tidak ada buku
dengan rating yang ditanyakan, cukup tuliskan “Tidak ada buku dengan
rating seperti itu”. Catatan: Gunakan pencarian biner/belah dua.}



#### [soal2.go]

```go

package main

import "fmt"


const nMax = 7919

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating     int
}

type DaftarBuku [nMax]Buku

func DaftarkanBuku(pustaka *DaftarBuku, n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&pustaka[i].id, &pustaka[i].judul, &pustaka[i].penulis, &pustaka[i].penerbit, &pustaka[i].eksemplar, &pustaka[i].tahun, &pustaka[i].rating)
	}
}

func CetakTerfavorit(pustaka DaftarBuku, n int) {
	if n == 0 {
		return
	}
	maxIdx := 0
	for i := 1; i < n; i++ {
		if pustaka[i].rating > pustaka[maxIdx].rating {
			maxIdx = i
		}
	}
	fmt.Println(pustaka[maxIdx].judul, pustaka[maxIdx].penulis, pustaka[maxIdx].penerbit, pustaka[maxIdx].tahun)
}

func UrutBuku(pustaka *DaftarBuku, n int) {
	for i := 1; i < n; i++ {
		key := pustaka[i]
		j := i - 1
		for j >= 0 && pustaka[j].rating < key.rating {
			pustaka[j+1] = pustaka[j]
			j--
		}
		pustaka[j+1] = key
	}
}

func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	limit := 5
	if n < limit {
		limit = n
	}
	for i := 0; i < limit; i++ {
		fmt.Println(pustaka[i].judul)
	}
}

func CariBuku(pustaka DaftarBuku, n int, r int) {
	kr := 0
	kn := n - 1
	found := -1

	for kr <= kn {
		med := (kr + kn) / 2
		if pustaka[med].rating == r {
			found = med
			break
		} else if pustaka[med].rating < r {
			kn = med - 1
		} else {
			kr = med + 1
		}
	}

	if found != -1 {
		fmt.Println(pustaka[found].judul, pustaka[found].penulis, pustaka[found].penerbit, pustaka[found].tahun, pustaka[found].eksemplar, pustaka[found].rating)
	} else {
		fmt.Println("Tidak ada buku dengan rating seperti itu")
	}
}

func main() {
	var pustaka DaftarBuku
	var n, r int

	fmt.Scan(&n)
	DaftarkanBuku(&pustaka, n)

	fmt.Scan(&r)

	CetakTerfavorit(pustaka, n)
	UrutBuku(&pustaka, n)
	Cetak5Terbaru(pustaka, n)
	CariBuku(pustaka, n, r)
}
```

## Output Unguided :

#### Output

![Screenshot Output Unguided 1\_1](https://github.com/ferdyandhestafirmansyah-boop/109082500109_Ferdyan-Dhesta-Firmansyah/blob/main/Modul%2014/Output/output-soal2.png)

#### Penjelasan
program ini digunakan untuk Baris pertama adalah bilangan bulat N yang
menyatakan banyaknya data buku yang ada di dalam perpustakaan. N baris berikutnya,
masing-masingnya adalah data buku sesuai dengan atribut atau field pada struct. Baris
terakhir adalah bilangan bulat yang menyatakan rating buku yang akan dicari.

Baris pertama adalah data buku terfavorit, baris kedua
adalah lima judul buku dengan rating tertinggi, selanjutnya baris terakhir adalah data buku
yang dicari sesuai rating yang diberikan pada masukan baris terakhir.