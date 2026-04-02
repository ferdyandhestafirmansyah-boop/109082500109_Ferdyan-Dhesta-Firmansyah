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