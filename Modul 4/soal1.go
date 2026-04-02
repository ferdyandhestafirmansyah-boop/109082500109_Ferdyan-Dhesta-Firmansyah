package main

import "fmt"

func factorial(n int, hasil *int) {
	*hasil = 1
	for i := 1; i <= n; i++ {
		*hasil *= i
	}
}

func permutation(n, r int, hasil *int) {
	var nFact, nmrFact int
	factorial(n, &nFact)
	factorial(n-r, &nmrFact)
	*hasil = nFact / nmrFact
}

func combination(n, r int, hasil *int) {
	var pVal, rFact int
	permutation(n, r, &pVal) 
	factorial(r, &rFact)
	*hasil = pVal / rFact
}

func main() {
	var a, b, c, d int
	var resP1, resC1, resP2, resC2 int

	fmt.Scan(&a, &b, &c, &d)

	permutation(a, c, &resP1)
	combination(a, c, &resC1)
	fmt.Println(resP1, resC1)

	permutation(b, d, &resP2)
	combination(b, d, &resC2)
	fmt.Println(resP2, resC2)
}