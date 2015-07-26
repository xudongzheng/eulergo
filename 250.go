package main

import (
	"fmt"
	"math/big"
	"time"
)

func main() {
	t := time.Now()
	defer func() {
		fmt.Println(time.Since(t))
	}()

	data := make([]int, 250251)
	mod := big.NewInt(250)
	for i := 0; i <= 250250; i++ {
		num := big.NewInt(int64(i))
		result := new(big.Int).Exp(num, num, mod)
		data[i] = int(result.Int64())
	}

	sum := 0
	count := make([]int, 250)
	for _, value := range data {
		count[value]++
		sum += value
	}
	fmt.Println(count)
	fmt.Println(sum, sum/250)
}
