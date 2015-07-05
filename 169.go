package main

import (
	"fmt"
	"math/big"
	"strconv"
)

var str string
var ones []int

func handle(mask int) int {
	product := 1
	for i := 0; i < len(ones); i++ {
		// Determine if the 1 will be split.
		if mask&(1<<uint(i)) > 0 {
			// Yes it will be split. Determine if it is the last 1 of all the
			// 1's. We have to determine the multiplier differently.
			if i == len(ones)-1 {
				// Yes, it is the last 1. Handle and break.
				multiplier := len(str) - ones[i] - 1
				product *= multiplier
				break
			}

			var multiplier int

			// Determine if the next 1 is split. If so, we get one additional
			// possibility.
			if mask&(1<<uint(i+1)) > 0 {
				multiplier = ones[i+1] - ones[i]
			} else {
				multiplier = ones[i+1] - ones[i] - 1
			}

			// Multiply product by the number of ways the specific 1 can be
			// split.
			product *= multiplier
		}
	}

	return product
}

func main() {
	num := new(big.Int).Exp(big.NewInt(10), big.NewInt(25), nil)
	data := num.Bytes()
	for _, value := range data {
		str += fmt.Sprintf("%08s", strconv.FormatUint(uint64(value), 2))
	}

	ones = make([]int, 0)
	for i := 0; i < len(str); i++ {
		if str[i] == '1' {
			ones = append(ones, i)
		}
	}

	// Keep track of possibilities.
	sum := 0

	// This will determine which of the 1's will be split.
	for i := 0; i < 1<<uint(len(ones)); i++ {
		sum += handle(i)
	}

	fmt.Println(sum)
}
