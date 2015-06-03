package main

import "fmt"

const rooms = 30

func handle(i int) int {
	if i > rooms {
		return rooms + 1
	}
	total := rooms + 1
	cards := make([]int, total)
	for j := 0; j < i; j++ {
		cards[rooms-j] = j + 1
	}
	carry := i - 2
	for j := rooms - i; j >= 0; j-- {
		target := cards[j+1] - i + 1
		times := (target + carry - 1) / carry // divide with quick integer ceil()
		cards[j] = cards[j+1] + 1 + times*2
	}

	return cards[0]
}

func main() {
	sum := 0
	for j := 3; j <= 40; j++ {
		sum += handle(j)
	}
	fmt.Println(sum)
}
