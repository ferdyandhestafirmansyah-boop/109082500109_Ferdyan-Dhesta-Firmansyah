<h1 align="center">Laporan Praktikum Modul 1</h1>
<p align="center">[Ferdyan Dhesta Firmansyah] - [109082500109]</p>

## Unguided

### 1. Deret fibonacci adalah sebuah deret dengan nilai suku ke-0 dan ke-1 adalah 0 dan 1, dan nilai
suku ke-n selanjutnya adalah hasil penjumlahan dua suku sebelumnya. Secara umum dapat
diformulasikan Sn = Sn−1 + Sn−2 . Berikut ini adalah contoh nilai deret fibonacci hingga suku
ke-10. Buatlah program yang mengimplementasikan fungsi rekursif pada deret fibonacci
tersebut.

#### [soal1.go]

```go
package 

import "fmt"


func fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}

func main() {
	var n int
	
	fmt.Scan(&n)
	fmt.Println(fibonacci(n))
}
```

## Output Unguided :

#### Output

![Screenshot Output Unguided 1\_1](https://github.com/ferdyandhestafirmansyah-boop/109082500109_Ferdyan-Dhesta-Firmansyah/blob/main/Modul%205/Output/output-soal1.png)

#### Penjelasan

Jadi program ini membaca 1 angka yaitu n, program menghitung deret bilangan menggunakan penjumlahan dua angka sebelumnya, dijadikan batasan awal (angka 0 dan 1) untuk menghentikan hitungan mundurnya. Kemudian fungsi fibonacci untuk menghitung dan mengembalikan nilai suku ke-n tersebut.


### 2.Buatlah sebuah program yang digunakan untuk menampilkan pola bintang berikut ini dengan
menggunakan fungsi rekursif. N adalah masukan dari user.
#### [soal2.go]

```go

package main

import "fmt"

func cetakPola(n int) {
	if n == 0 {
		return
	}

	cetakPola(n - 1)

	for i := 0; i < n; i++ {
		fmt.Print("*")
	}

	fmt.Println()
}

func main() {
	var n int
	fmt.Scan(&n)
	cetakPola(n)
}
```

## Output Unguided :

#### Output

![Screenshot Output Unguided 1\_1](https://github.com/ferdyandhestafirmansyah-boop/109082500109_Ferdyan-Dhesta-Firmansyah/blob/main/Modul%205/Output/output-soal2.png)

#### Penjelasan
Jadi program ini membaca 1 angka yaitu N, program mencetak simbol bintang menggunakan perulangan, dijadikan teknik penundaan cetak untuk membalik urutan agar bintang terbentuk dari kecil ke besar. Kemudian fungsi cetakPola untuk menampilkan barisan bintang tersebut ke layar.

### 3. Buatlah program yang mengimplementasikan rekursif untuk menampilkan faktor bilangan dari
suatu N, atau bilangan yang apa saja yang habis membagi N.
Masukan terdiri dari sebuah bilangan bulat positif N.
Keluaran terdiri dari barisan bilangan yang menjadi faktor dari N (terurut dari 1 hingga N ya).
#### [soal3.go]
```go
package main

import "fmt"
func cetakFaktor(n int, i int) {
	if i == 0 {
		return
	}
	cetakFaktor(n, i-1)
	if n%i == 0 {
		fmt.Print(i, " ")
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	cetakFaktor(n, n)
	fmt.Println() 
}
```

## Output Unguided :

#### Output

![Screenshot Output Unguided 1\_1](https://github.com/ferdyandhestafirmansyah-boop/109082500109_Ferdyan-Dhesta-Firmansyah/blob/main/Modul%205/Output/output-soal3.png)

#### Penjelasan
Jadi program ini membaca 1 angka yaitu N, program mencari angka pembagi yang pas menggunakan sisa bagi (modulo), dijadikan teknik penundaan agar hasil pembaginya bisa dicetak urut dari yang terkecil hingga terbesar. Kemudian fungsi cetakFaktor untuk mengecek dan mencetak angka apa saja yang berhasil membagi N.
