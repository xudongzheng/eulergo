package main

import (
	"fmt"
	"math/big"
)

var primes []int64
var primeCount int
var result int64
var alloc []int64
var multiplier []int64

func handle(i []int64) {
	product := int64(1)
	for _, value := range i {
		product *= int64(value)
		if product > 10000000000000000 {
			return
		}
	}
	row := 10000000000000000 / product
	result += multiplier[len(i)] * row
}

func choose(i, j int, c []int64) {
	if i == 1 {
		for k := j; k < primeCount; k++ {
			c[0] = primes[k]
			handle(alloc)
		}
		return
	}
	for k := j; k < primeCount; k++ {
		c[0] = primes[k]
		choose(i-1, k+1, c[1:])
	}
}

func main() {
	for i := 4; i <= 25; i++ {
		alloc = make([]int64, i)
		choose(i, 0, alloc)
	}

	// Finally print the result.
	fmt.Println(result)
}

func init() {
	// Lazy prime sieve.
	for i := int64(2); i < 100; i++ {
		if big.NewInt(i).ProbablyPrime(10) {
			primes = append(primes, i)
		}
	}

	// We'll need to quickly determine number of primes (25).
	primeCount = len(primes)

	// We need to calculate the multiplier for inclusion exclusion.
	multiplier = make([]int64, 26)
	for i := int64(1); i <= 21; i++ {
		multiplier[i+3] = i * (i + 1) * (i + 2) / 6
		if i%2 == 0 {
			multiplier[i+3] *= -1
		}
	}
}
