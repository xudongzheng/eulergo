package main

import "fmt"

func main() {
	head := make([]int, 31)
	tail := make([]int, 31)
	head[0], head[1], head[2] = 0, 1, 2
	tail[0], tail[1], tail[2] = 0, 1, 2

	for i := 3; i <= 30; i++ {
		head[i] = tail[i-1] + tail[i-2]
		tail[i] = tail[i-1] + head[i-1]
	}

	total := make([]int, 31)
	total[0] = 1
	for i := 1; i <= 30; i++ {
		total[i] = head[i] + tail[i]
	}

	var result int
	result += total[30]
	for i := 0; i < 30; i++ {
		result += total[i] * total[29-i]
	}

	fmt.Println(result)
}
