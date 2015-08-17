package main

import (
	"fmt"
	"math"
	"math/big"
)

func calc(f float64) float64 {
	num := math.Log(1e9) - 1000*math.Log1p(-f)
	den := math.Log1p(2*f) - math.Log1p(-f)
	return num / den
}

func main() {
	min := 1000
	for f := 0.0001; f < 1; f += 0.0001 {
		v := int(math.Ceil(calc(f)))
		if v < min {
			min = v
		}
	}

	num := big.NewInt(0)
	for i := int64(min); i <= 1000; i++ {
		num.Add(num, new(big.Int).Binomial(1000, i))
	}
	den := new(big.Int).Exp(big.NewInt(2), big.NewInt(1000), nil)

	result := new(big.Rat).SetInt(num)
	result.Quo(result, new(big.Rat).SetInt(den))
	fmt.Println(result.FloatString(12))
}
