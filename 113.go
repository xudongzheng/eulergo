package main

import "fmt"

func main() {
	choose1 := int64(1)
	for i := int64(101); i <= int64(109); i++ {
		choose1 *= i
		choose1 /= -100 + i
	}
	choose1 -= 1

	choose2 := int64(1)
	for i := int64(101); i <= int64(110); i++ {
		choose2 *= i
		choose2 /= -100 + i
	}
	choose2 -= 101

	sum := choose1 + choose2 - 9*100
	fmt.Println(sum)
}
