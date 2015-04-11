package main

import "fmt"

func main() {
	i := 100
	a := i * (i + 1) * (2*i + 1) / 6
	b := (i*i + i) / 2
	fmt.Println(b*b - a)
}
