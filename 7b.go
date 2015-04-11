package main

import (
	"fmt"
	"math/big"
)

func main() {
	i := big.NewInt(1)
	for count := 0; count < 10001; {
		i.Add(i, big.NewInt(1))
		if i.ProbablyPrime(16) {
			count++
		}
	}
	fmt.Println(i)
}
