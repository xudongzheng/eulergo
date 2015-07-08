package main

import (
	"bytes"
	"fmt"
	"sort"
)

var rows [][]uint8

type sorter struct{}

func (s sorter) Len() int {
	return len(rows)
}

func (s sorter) Less(i, j int) bool {
	return bytes.Compare(rows[i], rows[j]) <= 0
}

func (s sorter) Swap(i, j int) {
	rows[i], rows[j] = rows[j], rows[i]
}

func generate23() {
	rows = make([][]uint8, 0)
	for i := 0; i < 1<<16; i++ {
		sum := 0
		row := make([]uint8, 0)
		for j := 0; j < 16; j++ {
			if i&(1<<uint(j)) == 0 {
				sum += 2
				row = append(row, 2)
			} else {
				sum += 3
				row = append(row, 3)
			}
			if sum == 32 {
				rows = append(rows, row)
				break
			} else if sum > 32 {
				break
			}
		}
	}

	sort.Sort(sorter{})

	last := make([]byte, 0)
	for i := 0; i < len(rows); i++ {
		if bytes.Equal(last, rows[i]) {
			rows = append(rows[:i], rows[i+1:]...)
			i--
		}
		last = make([]byte, len(rows[i]))
		copy(last, rows[i])
	}
}

func main() {
	// Generate permutations of 2 and 3 that sum of to 32 (and order matters).
	// This isn't the most efficient way of doing so but is very easy to write.
	generate23()

	// Keep track of which rows are allowed to be stacked together.
	allows := make([][]int, len(rows))

	// Keep track of the sums reached. Make size 33 to include 32 (the final
	// one).
	cache := make([]bool, 33)

	for i := 0; i < len(rows); i++ {
		sum := uint8(0)
		for _, value := range rows[i] {
			sum += value
			cache[sum] = true
		}

		for j := i + 1; j < len(rows); j++ {
			crack := false
			sum = 0
			for _, value := range rows[j] {
				sum += value
				if cache[sum] && sum != 32 {
					crack = true
					break
				}
			}

			if !crack {
				allows[i] = append(allows[i], j)
				allows[j] = append(allows[j], i)
			}
		}

		// Reset cache to false.
		cache = make([]bool, 33)
	}

	data := make([][]int, 10)
	data[0] = make([]int, len(rows))
	for i := 0; i < len(rows); i++ {
		data[0][i] = 1
	}

	for i := 1; i <= 9; i++ {
		data[i] = make([]int, len(rows))
		for j := 0; j < len(rows); j++ {
			sum := 0
			for _, value := range allows[j] {
				sum += data[i-1][value]
			}
			data[i][j] = sum
		}
	}

	result := 0
	for _, value := range data[9] {
		result += value
	}
	fmt.Println(result)
}
