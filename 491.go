package main

import "fmt"

var data [][]int

func handleInt(i int) {
	row := make([]int, 10)
	digits := 0
	for j := 0; j < 10; j++ {
		row[j] = i % 3
		digits += row[j]
		i /= 3
	}

	if digits != 10 {
		return
	}

	sum := 0
	for key, value := range row {
		sum += key * value
	}

	if sum%11 == 1 {
		data = append(data, row)
	}
}

func choose(n, r int) int {
	switch r {
	case 0:
		return 1
	case 1:
		return n
	case 2:
		return n * (n - 1) / 2
	default:
		panic("not implemented")
	}
}

func countA(row []int) int {
	remain := 10
	result := 1

	// zero is an edge case.
	result *= choose(9, row[0])
	remain -= row[0]

	for i := 1; i < 10; i++ {
		result *= choose(remain, row[i])
		remain -= row[i]
	}

	return result
}

func countB(row []int) int {
	remain := 10
	result := 1

	// Calculate the complement.
	for i := 0; i < 10; i++ {
		row[i] = 2 - row[i]
	}

	for i := 0; i < 10; i++ {
		result *= choose(remain, row[i])
		remain -= row[i]
	}

	return result
}

func handleRow(row []int) int {
	return countA(row) * countB(row)
}

func main() {
	// Initialize data slice.
	data = make([][]int, 0)

	for i := 0; i < 59049; i++ {
		handleInt(i)
	}

	result := 0
	for _, value := range data {
		result += handleRow(value)
	}
	fmt.Println(result)
}
