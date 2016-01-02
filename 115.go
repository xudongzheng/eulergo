package main

import "fmt"

func main() {
	m := 50

	count := make([]int, m)
	for i := 0; i < m; i++ {
		count[i] = 1
	}
	for i := m; ; i++ {
		count = append(count, 0)
		count[i] += count[i-1] + count[0]
		for j := 0; j < i-m; j++ {
			count[i] += count[j]
		}
		if count[i] > 1000000 {
			fmt.Println(i)
			return
		}
	}
}
