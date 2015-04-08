package main

import (
	"fmt"
	"math/big"
)

func handle(n int64) int64 {
	var max int
	for i := int64(2); i <= int64(20); i++ {
		tmp := i
		count := 0
		for tmp != 1 {
			if tmp%n == 0 {
				tmp /= n
				count++
			} else {
				tmp = 1
			}
		}
		if count > max {
			max = count
		}
	}
	return int64(max)
}

func main() {
	product := big.NewInt(1)
	for i := int64(2); i <= 20; i++ {
		if big.NewInt(int64(i)).ProbablyPrime(10) {
			exp := handle(i)
			product.Mul(product, big.NewInt(0).Exp(big.NewInt(i), big.NewInt(exp), nil))
		}
	}
	fmt.Println(product)
}
