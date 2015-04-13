package main

import (
	"fmt"
	"math/big"
)

func main() {
	// 40 choose 20
	i := big.NewInt(1)
	for j := 21; j <= 40; j++ {
		i.Mul(i, big.NewInt(int64(j)))
	}
	for j := 1; j <= 20; j++ {
		i.Div(i, big.NewInt(int64(j)))
	}
	fmt.Println(i)
}
