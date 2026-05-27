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