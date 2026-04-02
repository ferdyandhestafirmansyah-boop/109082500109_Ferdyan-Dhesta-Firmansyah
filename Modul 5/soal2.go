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