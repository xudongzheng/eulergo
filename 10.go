package main

import "fmt"

func main() {
	data := make([]bool, 2000000)
	for i := 2; i < 2000000; i++ {
		if data[i] {
			continue
		}
		for j := 2 * i; j < 2000000; j += i {
			data[j] = true
		}
	}
	sum := 0
	for i := 2; i < 2000000; i++ {
		if !data[i] {
			sum += i
		}
	}
	fmt.Println(sum)
}
