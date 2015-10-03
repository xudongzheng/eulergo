package main

import (
	"fmt"
	"math/big"
	"sort"
)

// Initiate arrays to store data.
var pre, post [51][51]*big.Int

func guess(i int) bool {
	rat := big.NewRat(int64(i), 1e11)
	rat.Add(rat, big.NewRat(50, 1))
	result := big.NewRat(0, 1)
	row := big.NewRat(1, 1)
	for q := 0; q <= 50; q++ {
		result.Add(result, new(big.Rat).Mul(row, new(big.Rat).SetInt(post[20][q])))
		row.Mul(row, rat)
	}

	for i := 0; i < 50; i++ {
		result.Quo(result, rat)
	}

	target := big.NewRat(2, 100)
	if result.Cmp(target) == 1 {
		return false
	}

	return true
}

func main() {
	// Initiate initial constant.
	pre[0][0] = big.NewInt(1)

	for i := 1; i <= 50; i++ {
		// Reset temp array to zero.
		for x := 0; x <= i; x++ {
			for q := 0; q <= i; q++ {
				post[x][q] = big.NewInt(0)
			}
		}

		// Do necessary calculations.
		for x := 0; x < i; x++ {
			for q := 0; q < i; q++ {
				n := new(big.Int).Mul(big.NewInt(int64(i)), pre[x][q])
				post[x+1][q+1].Add(post[x+1][q+1], pre[x][q])
				post[x+1][q].Sub(post[x+1][q], n)
				post[x][q].Add(post[x][q], n)
			}
		}

		// Copy temp array to storage.
		pre = post
	}

	// Find the optimal q that gives probability of 0.02.
	i := sort.Search(1e13, guess)
	rat := big.NewRat(int64(i), 1e11)
	rat.Add(rat, big.NewRat(50, 1))
	fmt.Println(rat.FloatString(10))
}
