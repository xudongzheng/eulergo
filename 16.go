package main

import (
	"fmt"
	"math/big"
)

func main() {
	result := new(big.Int).Exp(big.NewInt(2), big.NewInt(1000), nil)
	str := result.String()
	var sum int
	for _, value := range str {
		sum += int(value - '0')
	}
	fmt.Println(sum)
}
