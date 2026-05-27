<h1 align="center">Laporan Praktikum Modul 1</h1>
<p align="center">[Ferdyan Dhesta Firmansyah] - [109082500109]</p>

## Unguided

### 1. Hercules, preman terkenal seantero ibukota, memiliki kerabat di banyak daerah. Tentunya
Hercules sangat suka mengunjungi semua kerabatnya itu.
Diberikan masukan nomor rumah dari semua kerabatnya di suatu daerah, buatlah program
rumahkerabat yang akan menyusun nomor-nomor rumah kerabatnya secara terurut
membesar menggunakan algoritma selection sort.
Masukan dimulai dengan sebuah integer n (0 < n < 1000), banyaknya daerah kerabat
Hercules tinggal. Isi n baris berikutnya selalu dimulai dengan sebuah integer m (0 < m <
1000000) yang menyatakan banyaknya rumah kerabat di daerah tersebut, diikuti dengan
rangkaian bilangan bulat positif, nomor rumah para kerabat.

Keluaran terdiri dari n baris, yaitu rangkaian rumah kerabatnya terurut membesar di masing-
masing daerah.
Keterangan: Terdapat 3 daerah dalam contoh input, dan di masing-masing daerah
mempunyai 5, 6, dan 3 kerabat.

#### [soal1.go]

```go
package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for k := 0; k < n; k++ {
		var m int
		fmt.Scan(&m)

		arr := make([]int, m)
		for i := 0; i < m; i++ {
			fmt.Scan(&arr[i])
		}

		for i := 0; i < m-1; i++ {
			idxMin := i
			for j := i + 1; j < m; j++ {
				if arr[j] < arr[idxMin] {
					idxMin = j
				}
			}
			temp := arr[idxMin]
			arr[idxMin] = arr[i]
			arr[i] = temp
		}

		for i := 0; i < m; i++ {
			if i > 0 {
				fmt.Print(" ")
			}
			fmt.Print(arr[i])
		}
		fmt.Println()
	}
}
```

## Output Unguided :

#### Output

![Screenshot Output Unguided 1\_1](https://github.com/ferdyandhestafirmansyah-boop/109082500109_Ferdyan-Dhesta-Firmansyah/blob/main/Modul%2010/Output/output-soal1.png)

#### Penjelasan
Jadi program ini mencari banyaknya daerah kerabat Hercules tinggal. Isi n baris berikutnya selalu dimulai dengan sebuah integer m (0 < m <1000000) yang menyatakan banyaknya rumah kerabat di daerah tersebut, diikuti dengan rangkaian bilangan bulat positif, nomor rumah para kerabat.


### 2.Belakangan diketahui ternyata Hercules itu tidak berani menyeberang jalan, maka selalu
diusahakan agar hanya menyeberang jalan sesedikit mungkin, hanya diujung jalan. Karena
nomor rumah sisi kiri jalan selalu ganjil dan sisi kanan jalan selalu genap, maka buatlah
program kerabat dekat yang akan menampilkan nomor rumah mulai dari nomor yang ganjil
lebih dulu terurut membesar dan kemudian menampilkan nomor rumah dengan nomor
genap terurut mengecil.
Format Masukan masih persis sama seperti sebelumnya.
Keluaran terdiri dari n baris, yaitu rangkaian rumah kerabatnya terurut membesar untuk
nomor ganjil, diikuti dengan terurut mengecil untuk nomor genap, di masing-masing daerah.
Keterangan: Terdapat 3 daerah dalam contoh masukan. Baris kedua berisi campuran
bilangan ganjil dan genap. Baris berikutnya hanya berisi bilangan ganjil, dan baris terakhir
hanya berisi bilangan genap.
Petunjuk:
• Waktu pembacaan data, bilangan ganjil dan genap dipisahkan ke dalam dua array
yang berbeda, untuk kemudian masing-masing diurutkan tersendiri.
• Atau, tetap disimpan dalam satu array, diurutkan secara keseluruhan. Tetapi pada
waktu pencetakan, mulai dengan mencetak semua nilai ganjil lebih dulu, kemudian
setelah selesai cetaklah semua nilai genapnya.

#### [soal2.go]

```go

package main

import "fmt"


const NMAX = 1000000

var ganjil [NMAX]int
var genap [NMAX]int

func main() {
	var n int
	fmt.Scan(&n)

	for k := 0; k < n; k++ {
		var m int
		fmt.Scan(&m)

		var nGanjil, nGenap int = 0, 0

		for i := 0; i < m; i++ {
			var val int
			fmt.Scan(&val)
			if val%2 != 0 {
				ganjil[nGanjil] = val
				nGanjil++
			} else {
				genap[nGenap] = val
				nGenap++
			}
		}

		for i := 0; i < nGanjil-1; i++ {
			idxMin := i
			for j := i + 1; j < nGanjil; j++ {
				if ganjil[j] < ganjil[idxMin] {
					idxMin = j
				}
			}
			temp := ganjil[idxMin]
			ganjil[idxMin] = ganjil[i]
			ganjil[i] = temp
		}

		for i := 0; i < nGenap-1; i++ {
			idxMax := i
			for j := i + 1; j < nGenap; j++ {
				if genap[j] > genap[idxMax] {
					idxMax = j
				}
			}
			temp := genap[idxMax]
			genap[idxMax] = genap[i]
			genap[i] = temp
		}

		first := true
		for i := 0; i < nGanjil; i++ {
			if !first {
				fmt.Print(" ")
			}
			fmt.Print(ganjil[i])
			first = false
		}
		for i := 0; i < nGenap; i++ {
			if !first {
				fmt.Print(" ")
			}
			fmt.Print(genap[i])
			first = false
		}
		fmt.Println()
	}
}
```

## Output Unguided :

#### Output

![Screenshot Output Unguided 1\_1](https://github.com/ferdyandhestafirmansyah-boop/109082500109_Ferdyan-Dhesta-Firmansyah/blob/main/Modul%2010/Output/output-soal2.png)

#### Penjelasan
program kerabat dekat yang akan menampilkan nomor rumah mulai dari nomor yang ganjil
lebih dulu terurut membesar dan kemudian menampilkan nomor rumah dengan nomor
genap terurut mengecil.

### 3. PKompetisi pemrograman yang baru saja berlalu diikuti oleh 17 tim dari berbagai perguruan
tinggi ternama. Dalam kompetisi tersebut, setiap tim berlomba untuk menyelesaikan
sebanyak mungkin problem yang diberikan. Dari 13 problem yang diberikan, ada satu
problem yang menarik. Problem tersebut mudah dipahami, hampir semua tim mencoba
untuk menyelesaikannya, tetapi hanya 3 tim yang berhasil. Apa sih problemnya?
"Median adalah nilai tengah dari suatu koleksi data yang sudah terurut. Jika jumlah data
genap, maka nilai median adalah rerata dari kedua nilai tengahnya. Pada problem ini, semua
data merupakan bilangan bulat positif, dan karenanya rerata nilai tengah dibulatkan ke
bawah."
Buatlah program median yang mencetak nilai median terhadap seluruh data yang sudah
terbaca, jika data yang dibaca saat itu adalah 0.
Masukan berbentuk rangkaian bilangan bulat. Masukan tidak akan berisi lebih dari 1000000
data, tidak termasuk bilangan 0. Data 0 merupakan tanda bahwa median harus dicetak, tidak
termasuk data yang dicari mediannya. Data masukan diakhiri dengan bilangan bulat -5313.
Keluaran adalah median yang diminta, satu data per baris.
Keterangan:
Sampai bilangan 0 yang pertama, data terbaca adalah 7 23 11, setelah tersusun: 7 11 23,
maka median saat itu adalah 11.
Sampai bilangan 0 yang kedua, data adalah 7 23 11 5 19 2 29 3 13 17, setelah tersusun
diperoleh: 2 3 5 7 11 13 17 19 23 29. Karena ada 10 data, genap, maka median adalah
(11+13)/2=12.
Petunjuk:
Untuk setiap data bukan 0 (dan bukan marker -5313541) simpan ke dalam array, Dan setiap
kali menemukan bilangan 0, urutkanlah data yang sudah tersimpan dengan menggunakan
metode insertion sort dan ambil mediannya.
#### [soal3.go]
```go
package main

import (
	"fmt"
)

const NMAX = 1000000

var data [NMAX]int

func main() {
	var n int = 0

	for {
		var val int
		fmt.Scan(&val)

		if val == -5313 {
			break
		}

		if val == 0 {
			if n == 0 {
				continue
			}

			for i := 0; i < n-1; i++ {
				idxMin := i
				for j := i + 1; j < n; j++ {
					if data[j] < data[idxMin] {
						idxMin = j
					}
				}
				temp := data[idxMin]
				data[idxMin] = data[i]
				data[i] = temp
			}

			var median int
			if n%2 == 0 {
				median = (data[(n/2)-1] + data[n/2]) / 2
			} else {
				median = data[n/2]
			}
			fmt.Println(median)

		} else {
			data[n] = val
			n++
		}
	}
}
```

## Output Unguided :

#### Output

![Screenshot Output Unguided 1\_1](https://github.com/ferdyandhestafirmansyah-boop/109082500109_Ferdyan-Dhesta-Firmansyah/blob/main/Modul%2010/Output/output-soal3.png)

#### Penjelasan
Jadi program ini median yang mencetak nilai median terhadap seluruh data yang sudah terbaca, jika data yang dibaca saat itu adalah 0.