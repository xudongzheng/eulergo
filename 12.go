package main

import (
	"fmt"
	"math"
)

func main() {
	for i := 1; ; i++ {
		sum := i * (i + 1) / 2
		end := int(math.Floor(math.Sqrt(float64(sum))))
		var count int
		for j := 1; j <= end; j++ {
			if sum%j == 0 {
				count++
				if count > 250 {
					fmt.Println(sum)
					return
				}
			}
		}
	}
}
