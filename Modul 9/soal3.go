package 

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