package main

import "fmt"

func main() {
	data := make([]int, 1e7)
	for i := 2; i < 1e7; i++ {
		for j := i; j < 1e7; j += i {
			data[j]++
		}
	}
	count := 0
	for i := 2; i < 1e7-1; i++ {
		if data[i] == data[i+1] {
			count++
		}
	}
	fmt.Println(count)
}
