<h1 align="center">Laporan Praktikum Modul 1</h1>
<p align="center">[Ferdyan Dhesta Firmansyah] - [109082500109]</p>

## Unguided

### 1. Suatu lingkaran didefinisikan dengan koordinat titik pusat (cx, cy) dengan radius r. Apabila
diberikan dua buah lingkaran, maka tentukan posisi sebuah titik sembarang (x, y)
berdasarkan dua lingkaran tersebut. Gunakan tipe bentukan titik untuk menyimpan
koordinat, dan tipe bentukan lingkaran untuk menyimpan titik pusat lingkaran dan
radiusnya.
Masukan terdiri dari beberapa tiga baris. Baris pertama dan kedua adalah koordinat titik pusat
dan radius dari lingkaran 1 dan lingkaran 2, sedangkan baris ketiga adalah koordinat titik
sembarang. Asumsi sumbu x dan y dari semua titik dan juga radius direpresentasikan dengan
bilangan bulat.
Keluaran berupa string yang menyatakan posisi titik "Titik di dalam lingkaran 1 dan 2", "Titik
di dalam lingkaran 1", "Titik di dalam lingkaran 2", atau "Titik di luar lingkaran 1 dan 2".

#### [soal1.go]

```go
package main

import "fmt"

func didalam(c [3]int, p [2]int) bool {
	dx := p[0] - c[0]
	dy := p[1] - c[1]
	
	jarakKuadrat := (dx * dx) + (dy * dy)
	radiusKuadrat := c[2] * c[2]

	return jarakKuadrat <= radiusKuadrat
}

func main() {
	var c1, c2 [3]int
	var p [2]int

	fmt.Scan(&c1[0], &c1[1], &c1[2])

	fmt.Scan(&c2[0], &c2[1], &c2[2])

	fmt.Scan(&p[0], &p[1])

	inC1 := didalam(c1, p)
	inC2 := didalam(c2, p)


	if inC1 && inC2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if inC1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if inC2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
```

## Output Unguided :

#### Output

![Screenshot Output Unguided 1\_1](https://github.com/ferdyandhestafirmansyah-boop/109082500109_Ferdyan-Dhesta-Firmansyah/blob/main/Modul%209/Output/ouput-soal1.png)

####Penjelasan

Jadi program ini membaca data berupa koordinat titik pusat dan jari-jari untuk dua buah lingkaran, serta koordinat untuk satu titik sembarang. Program menghitung posisi titik tersebut menggunakan fungsi didalam dengan cara mencari nilai kuadrat jarak antara titik ke pusat lingkaran, lalu membandingkannya dengan nilai kuadrat jari-jarinya. program utama menggunakan hasil dari fungsi tersebut untuk mengecek dan mencetak kalimat yang menyatakan apakah titik itu berada di dalam lingkaran 1, lingkaran 2, di dalam keduanya, atau di luar kedua lingkaran tersebut.


### 2.Sebuah array digunakan untuk menampung sekumpulan bilangan bulat. Buatlah program
yang digunakan untuk mengisi array tersebut sebanyak N elemen nilai. Asumsikan array
memiliki kapasitas penyimpanan data sejumlah elemen tertentu. Program dapat
menampilkan beberapa informasi berikut:
a. Menampilkan keseluruhan isi dari array.
b. Menampilkan elemen-elemen array dengan indeks ganjil saja.
c. Menampilkan elemen-elemen array dengan indeks genap saja (asumsi indek ke-0 adalah
genap).
d. Menampilkan elemen-elemen array dengan indeks kelipatan bilangan x. x bisa diperoleh
dari masukan pengguna.
e. Menghapus elemen array pada indeks tertentu, asumsi indeks yang hapus selalu valid.
Tampilkan keseluruhan isi dari arraynya, pastikan data yang dihapus tidak tampil
f. Menampilkan rata-rata dari bilangan yang ada di dalam array.
g. Menampilkan standar deviasi atau simpangan baku dari bilangan yang ada di dalam array
tersebut.
h. Menampilkan frekuensi dari suatu bilangan tertentu di dalam array yang telah diisi
tersebut.
#### [soal2.go]

```go

package main

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
```

## Output Unguided :

#### Output

![Screenshot Output Unguided 1\_1](https://github.com/ferdyandhestafirmansyah-boop/109082500109_Ferdyan-Dhesta-Firmansyah/blob/main/Modul%209/Output/output-soal2.png)

#### Penjelasan
Jadi program ini membaca ukuran array sejumlah n dan mengisi array tersebut dengan input nilai dari pengguna. Program kemudian menampilkan berbagai informasi secara berurutan keseluruhan isi array, nilai pada indeks ganjil, elemen pada indeks genap, serta elemen pada indeks kelipatan bilangan x yang diinputkan. program meminta input indeks tertentu, lalu menghapus indeks dengan cara menggeser sisa nilai di sebelah kanannya ke kiri dan menampilkan kembali isi array yang baru.  program menghitung dan menampilkan nilai rata-rata, standar deviasi dan terakhir menghitung frekuensi atau seberapa banyak kemunculan suatu bilangan tertentu yang dicari oleh pengguna di dalam array tersebut.

### 3. Sebuah program digunakan untuk menyimpan dan menampilkan nama-nama klub yang
memenangkan pertandingan bola pada suatu grup pertandingan. Buatlah program yang
digunakan untuk merekap skor pertandingan bola 2 buah klub bola yang berlaga.
Pertama-tama program meminta masukan nama-nama klub yang bertanding, kemudian
program meminta masukan skor hasil pertandingan kedua klub tersebut. Yang disimpan
dalam array adalah nama-nama klub yang menang saja.
Proses input skor berhenti ketika skor salah satu atau kedua klub tidak valid (negatif). Di akhir
program, tampilkan daftar klub yang memenangkan pertandingan.
#### [soal3.go]
```go
package main

import (
	"fmt"
)

const NMAX = 1000

type arrString [NMAX]string

func main() {
	var klubA, klubB string
	var hasil arrString
	var skorA, skorB int
	var count int = 0
	var match int = 1

	fmt.Print("Klub A : ")
	fmt.Scan(&klubA)
	fmt.Print("Klub B : ")
	fmt.Scan(&klubB)

	for {
		fmt.Printf("Pertandingan %d : ", match)
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			hasil[count] = klubA
		} else if skorA < skorB {
			hasil[count] = klubB
		} else {
			hasil[count] = "Draw"
		}
		count++
		match++
	}

	for i := 0; i < count; i++ {
		fmt.Printf("Hasil %d : %s\n", i+1, hasil[i])
	}
	fmt.Println("Pertandingan selesai")
}
```

## Output Unguided :

#### Output

![Screenshot Output Unguided 1\_1](https://github.com/ferdyandhestafirmansyah-boop/109082500109_Ferdyan-Dhesta-Firmansyah/blob/main/Modul%209/Output/output-soal3.png)

#### Penjelasan
Jadi program ini membaca nama dua buah klub bola yang bertanding. program akan terus-menerus meminta input skor hasil pertandingan dari pengguna secara berulang. Di dalam perulangan tersebut, program membandingkan skor untuk menentukan siapa pemenangnya dan menyimpan nama pemenang tersebut ke dalam sebuah array statis.input skor ini akan terus berjalan dan baru akan terhenti apabila salah satu atau kedua skor yang dimasukkan bernilai negatif.lalu proses input selesai, program akan menampilkan kembali daftar seluruh hasil pertandingan di dalam array tersebut.

### 4. Sebuah array digunakan untuk menampung sekumpulan karakter, Anda diminta untuk
membuat sebuah subprogram untuk melakukan membalikkan urutan isi array dan memeriksa
apakah membentuk palindrom.
#### [soal3.go]
```go
package main

import (
	"fmt"
)

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var input string
	*n = 0
	for {
		fmt.Scan(&input)
		if input == "." {
			break
		}
		if *n < NMAX {
			(*t)[*n] = []rune(input)[0]
			*n++
		}
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c ", t[i])
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		temp := (*t)[i]
		(*t)[i] = (*t)[n-1-i]
		(*t)[n-1-i] = temp
	}
}

func palindrom(t tabel, n int) bool {
	var tempTabel tabel = t
	balikanArray(&tempTabel, n)
	
	for i := 0; i < n; i++ {
		if t[i] != tempTabel[i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int

	fmt.Print("Teks         : ")
	isiArray(&tab, &m)

	var reversedTab tabel = tab
	balikanArray(&reversedTab, m)

	fmt.Print("Reverse teks : ")
	cetakArray(reversedTab, m)

	fmt.Printf("Palindrom    ? %t\n", palindrom(tab, m))
}
```

## Output Unguided :

#### Output

![Screenshot Output Unguided 1\_1](https://github.com/ferdyandhestafirmansyah-boop/109082500109_Ferdyan-Dhesta-Firmansyah/blob/main/Modul%209/Output/output-soal4.png)

#### Penjelasan
Jadi program ini membaca karakter yang dimasukkan oleh pengguna satu per satu ke dalam sebuah array hingga pengguna mengetikkan tanda titik (.). Setelah array terisi, program akan membuat salinan dari array tersebut dan membalikkan urutan isinya untuk ditampilkan ke layar.program menggunakan sebuah fungsi untuk mengecek apakah kumpulan karakter tersebut membentuk palindrom dengan cara membandingkan isi array asli dengan array salinan yang sudah dibalik. Hasil akhir dari pengecekan tersebut akan ditampilkan ke layar dalam bentuk nilai boolean.