package 

import "fmt"

func didalam(c [3]int, p [2]int) bool {
	dx := p[0] - c[0]
	dy := p[1] - c[1]
	
	jarakKuadrat := (dx * dx) + (dy * dy)
	radiusKuadrat := c[2] * c[2]

	return jarakKuadrat <= radiusKuadrat
}

func () {
	var c1, c2 [3]int
	var p [2]int

	fmt.Scan(&c1[0], &c1[1], &c1[2])

	fmt.Scan(&c2[0], &c2[1], &c2[2])

	fmt.Scan(&p[0], &p[1])

	inC1 := didalam(c1, p)
	inC2 := didalam(c2, p)


	if inC1 && inC2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if inC1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if inC2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}