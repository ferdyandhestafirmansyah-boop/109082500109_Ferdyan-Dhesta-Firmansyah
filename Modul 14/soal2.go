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