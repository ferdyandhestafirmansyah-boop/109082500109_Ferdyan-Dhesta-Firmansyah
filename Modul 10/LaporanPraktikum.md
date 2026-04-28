<h1 align="center">Laporan Praktikum Modul 1</h1>
<p align="center">[Ferdyan Dhesta Firmansyah] - [109082500109]</p>

## Unguided

### 1. Sebuah program digunakan untuk mendata berat anak kelinci yang akan dijual ke pasar.
Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat anak
kelinci yang akan dijual.
Masukan terdiri dari sekumpulan bilangan, yang mana bilangan pertama adalah bilangan
bulat N yang menyatakan banyaknya anak kelinci yang akan ditimbang beratnya. Selanjutnya
N bilangan riil berikutnya adalah berat dari anak kelinci yang akan dijual.
Keluaran terdiri dari dua buah bilangan riil yang menyatakan berat kelinci terkecil dan
terbesar.

#### [soal1.go]

```go
package main

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
```

## Output Unguided :

#### Output

![Screenshot Output Unguided 1\_1](https://github.com/ferdyandhestafirmansyah-boop/109082500109_Ferdyan-Dhesta-Firmansyah/blob/main/Modul%2010/Output/output-soal1.png)

#### Penjelasan

Jadi program ini terdiri dari sekumpulan bilangan, yang mana bilangan pertama adalah bilanganbulat N yang menyatakan banyaknya anak kelinci yang akan ditimbang beratnya.Selanjutnya N bilangan riil berikutnya adalah berat dari anak kelinci yang akan dijual.


### 2.Sebuah program digunakan untuk menentukan tarif ikan yang akan dijual ke pasar. Program
ini menggunakan array dengan kapasitas 1000 untuk menampung data berat ikan yang akan
dijual.
Masukan terdiri dari dua baris, yang mana baris pertama terdiri dari dua bilangan bulat x dan
y. Bilangan x menyatakan banyaknya ikan yang akan dijual, sedangkan y adalah banyaknya
ikan yang akan dimasukan ke dalam wadah. Baris kedua terdiri dari sejumlah x bilangan riil
yang menyatakan banyaknya ikan yang akan dijual.
Keluaran terdiri dari dua baris. Baris pertama adalah kumpulan bilangan riil yang menyatakan
total berat ikan di setiap wadah (jumlah wadah tergantung pada nilai x dan y, urutan ikan yang
dimasukan ke dalam wadah sesuai urutan pada masukan baris ke-2). Baris kedua adalah
sebuah bilangan riil yang menyatakan berat rata-rata ikan di setiap wadah.
#### [soal2.go]

```go

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
```

## Output Unguided :

#### Output

![Screenshot Output Unguided 1\_1](https://github.com/ferdyandhestafirmansyah-boop/109082500109_Ferdyan-Dhesta-Firmansyah/blob/main/Modul%2010/Output/output-soal2.png)

#### Penjelasan
Jadi program ini terdiri dari dua baris, yang mana baris pertama terdiri dari dua bilangan bulat x dan
y. Bilangan x menyatakan banyaknya ikan yang akan dijual, sedangkan y adalah banyaknya
ikan yang akan dimasukan ke dalam wadah. Baris kedua terdiri dari sejumlah x bilangan riil
yang menyatakan banyaknya ikan yang akan dijual.

### 3. Pos Pelayanan Terpadu (posyandu) sebagai tempat pelayanan kesehatan perlu mencatat data
berat balita (dalam kg). Petugas akan memasukkan data tersebut ke dalam array. Dari data
yang diperoleh akan dicari berat balita terkecil, terbesar, dan reratanya.
Buatlah program dengan spesifikasi subprogram sebagai berikut:
#### [soal3.go]
```go
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
```

## Output Unguided :

#### Output

![Screenshot Output Unguided 1\_1](https://github.com/ferdyandhestafirmansyah-boop/109082500109_Ferdyan-Dhesta-Firmansyah/blob/main/Modul%2010/Output/output-soal3.png)

#### Penjelasan
Jadi program ini menyimpan deretan input berat balita ke dalam array. program mencari berat minimum dan maksimum sekaligus menggunakan pointer pada prosedur, menghitung nilai rata-ratanya menggunakan fungsi, lalu mencetak ketiga hasil tersebut.
