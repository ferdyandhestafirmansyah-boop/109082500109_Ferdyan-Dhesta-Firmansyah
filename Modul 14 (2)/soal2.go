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