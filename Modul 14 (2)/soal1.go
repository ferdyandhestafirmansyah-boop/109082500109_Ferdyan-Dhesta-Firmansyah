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