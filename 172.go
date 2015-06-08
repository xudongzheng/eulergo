package main

import (
	"fmt"
	"strconv"
)

var factorial []int

func orient(s string, i int) int {
	base := factorial[i]
	for _, value := range s {
		base /= factorial[value-'0']
	}
	return base
}

// Determine all combination of digits 1-9 appearing 3 times or less with
// total length of i. Essentially, look at numbers 0 to 4^9 in base 4, and then
// check the sum of the digits.
func count(i int) int {
	data := make([]string, 0)
	for j := int64(0); j < 1<<18; j++ {
		str := strconv.FormatInt(j, 4)
		sum := 0
		for _, value := range str {
			sum += int(value - '0')
		}
		if sum == i {
			data = append(data, str)
		}
	}

	// Look at all legitimate combination of digits and find number of
	// orientations.
	sum := 0
	for _, value := range data {
		sum += orient(value, i)
	}

	return sum
}

func main() {
	// Calculate factorials from 0 to 18.
	factorial = make([]int, 19)
	factorial[0] = 1
	for i := 1; i <= 18; i++ {
		factorial[i] = factorial[i-1] * i
	}

	// For each of these, we can add between 0 and 3 zeros. As a result, our
	// multipliers are 1, 17, 136, and 680.
	sum := 0
	sum += count(18)
	sum += 17 * count(17)
	sum += 136 * count(16)
	sum += 680 * count(15)
	fmt.Println(sum)
}
