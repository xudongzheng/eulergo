package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("13.txt")
	if err != nil {
		return
	}
	sum := big.NewInt(0)
	reader := bufio.NewReader(f)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		line = strings.TrimSpace(line)
		i, ok := new(big.Int).SetString(line, 10)
		if !ok {
			return
		}
		sum.Add(sum, i)
	}
	fmt.Println(sum.String()[:10])
}
