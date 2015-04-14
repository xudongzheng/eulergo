package main

import (
	"fmt"
	"math"
)

func main() {
	end := int(math.Ceil(math.Sqrt(50000000)))
	composites := make([]bool, end)
	for i := 2; i < end; i++ {
		if composites[i] {
			continue
		}
		for j := 2 * i; j < end; j += i {
			composites[j] = true
		}
	}
	var prime2, prime3, prime4 []int

	// Look at potential x^2
	for i := 2; i < end; i++ {
		if !composites[i] {
			prime2 = append(prime2, i)
		}
	}

	// Look at potential x^3
	end = int(math.Ceil(math.Pow(50000000, 1.0/3)))
	for _, value := range prime2 {
		if value > end {
			break
		}
		prime3 = append(prime3, value)
	}

	// Look at potential x^4
	end = int(math.Ceil(math.Pow(50000000, 1.0/4)))
	for _, value := range prime3 {
		if value > end {
			break
		}
		prime4 = append(prime4, value)
	}

	// Loop through prime4 first, then prime3, then prime2. Once we exceed 50M,
	// break.
	data := make(map[int]bool)
	for _, value4 := range prime4 {
		for _, value3 := range prime3 {
			for _, value2 := range prime2 {
				product := value4 * value4 * value4 * value4
				product += value3 * value3 * value3
				product += value2 * value2
				if product > 50000000 {
					break
				}
				data[product] = true
			}
		}
	}
	fmt.Println(len(data))
}
