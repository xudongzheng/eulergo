package main

import "fmt"

func exp(b, e int) int {
	r := 1
	for i := 0; i < e; i++ {
		r *= b
	}
	return r
}

func handle(i int) int {
	a := exp(16, i) - 3*exp(15, i) + 3*exp(14, i) - exp(13, i)
	b := exp(16, i) - 2*exp(15, i) + exp(14, i)
	return 13*a + 2*b
}

func main() {
	sum := 0
	for i := 3; i <= 16; i++ {
		sum += handle(i - 1)
	}
	fmt.Printf("%X\n", sum)
}
