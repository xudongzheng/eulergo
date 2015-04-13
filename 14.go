package main

import "fmt"

var data map[int]int

func calculate(i int) int {
	answer, ok := data[i]
	if ok {
		return answer
	}
	if i%2 == 0 {
		answer = calculate(i/2) + 1
	} else {
		answer = calculate(3*i+1) + 1
	}
	data[i] = answer
	return answer
}

func main() {
	data = make(map[int]int)
	data[1] = 1
	for i := 2; i < 1000000; i++ {
		calculate(i)
	}
	loc, size := 0, 0
	for i := 2; i < 1000000; i++ {
		if data[i] > size {
			loc = i
			size = data[i]
		}
	}
	fmt.Println(loc)
}
