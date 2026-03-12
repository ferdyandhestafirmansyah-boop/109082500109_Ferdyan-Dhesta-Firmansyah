<h1 align="center">Laporan Praktikum Modul 1</h1>
<p align="center">[Ferdyan Dhesta Firmansyah] - [109082500109]</p>

## Unguided

### 1. Telusuri program berikut dengan cara mengkompilasi dan mengeksekusi program. Silakan
masukan data yang sesuai sebanyak yang diminta program. Perhatikan keluaran yang
diperoleh. Coba terangkan apa sebenarnya yang dilakukan program tersebut?

#### [soal1.go]

```go
package main

import "fmt"

func main() {
	var (
	satu, dua, tiga string
	temp string
	)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&satu)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&dua)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&tiga)
	fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)
	temp = satu
	satu = dua
	dua = tiga
	tiga = temp
	fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
}
```

## Output Unguided :

#### Output

![Screenshot Output Unguided 1\_1](output/output-soal1.png)

#### Penjelasan

jadi program diatas dibuat menerima tiga input string, lalu memunculkan urutan awal input tersebut dan melakukan pertukaran posisi  string, lalu memunculkan urutan setelah ditukar


### 2. Buatlah sebuah program yang menerima input berupa warna dari ke 4 gelas reaksi
sebanyak 5 kali percobaan. Kemudian program akan menampilkan true apabila urutan
warna sesuai dengan informasi yang diberikan pada paragraf sebelumnya, dan false
untuk urutan warna lainnya.

#### [soal2.go]

```go

package main
import "fmt"

func main() {
	var w1, w2, w3, w4 string
	berhasil := true

	for i := 1; i <= 5; i++ {
		fmt.Print("Percobaan ", i, ": ")
		fmt.Scan(&w1, &w2, &w3, &w4)

		if !(w1 == "merah" && w2 == "kuning" && w3 == "hijau" && w4 == "ungu") {
			berhasil = false
		}
	}

	fmt.Println("BERHASIL:", berhasil)
}
```

## Output Unguided :

#### Output

![Screenshot Output Unguided 1\_1](output/output-soal2.png)

#### Penjelasan
jadi code diatas untuk memeriksa 5 percobaan memiliki urutan warna yang sama yaitu merah kuning hijau ungu,saat semua sesuai → true dan saat ada yang berbeda → false.


### 3. Dari berat parsel (dalam gram), harus dihitung total berat dalam kg dan sisanya (dalam
gram). Biaya jasa pengiriman adalah Rp. 10.000,- per kg. Jika sisa berat tidak kurang dari 500
gram, maka tambahan biaya kirim hanya Rp. 5,- per gram saja. Tetapi jika kurang dari 500
gram, maka tambahan biaya akan dibebankan sebesar Rp. 15,- per gram. Sisa berat (yang
kurang dari 1kg) digratiskan biayanya apabila total berat ternyata lebih dari 10kg.

#### [soal3.go]
```go
package main
import "fmt"

func main() {

	var gram int
	var kg, sisa int
	var biayaKg, biayaSisa, total int

	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&gram)

	kg = gram / 1000
	sisa = gram % 1000

	biayaKg = kg * 10000

	if kg > 10 {
		biayaSisa = 0
	} else {
		if sisa >= 500 {
			biayaSisa = sisa * 5
		} else {
			biayaSisa = sisa * 15
		}
	}

	total = biayaKg + biayaSisa

	fmt.Println("Detail berat:", kg, "kg +", sisa, "gr")
	fmt.Println("Detail biaya: Rp.", biayaKg, "+ Rp.", biayaSisa)
	fmt.Println("Total biaya: Rp.", total)
}
```

## Output Unguided :

#### Output

![Screenshot Output Unguided 1\_1](output/output-soal3.png)

#### Penjelasan
Program ini digunakan untuk menghitung biaya pengiriman parsel berdasarkan berat dalam gram. Pertama, pengguna memasukkan berat parsel, lalu program mengubahnya menjadi kilogram dan sisa gram. Biaya dihitung Rp 10.000 per kg. Untuk sisa gram, jika ≥ 500 gram dikenakan Rp 5 per gram, dan jika < 500 gram dikenakan Rp 15 per gram. Namun jika berat total lebih dari 10 kg, maka biaya sisa gram digratiskan. Terakhir program menampilkan detail berat, detail biaya, dan total biaya pengiriman.