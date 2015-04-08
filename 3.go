package main

import (
	"fmt"
	"math"
)

func main() {
	num := 600851475143
	start := 2
	largest := 2
	for {
		end := int(math.Sqrt(float64(num)))
		for i := start; i <= end; i++ {
			if num%i == 0 {
				num /= i
				largest = i
				start = i
			}
		}
		if num == 1 {
			break
		}
	}
	fmt.Println(largest)
}
