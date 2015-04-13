package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("11.txt")
	if err != nil {
		return
	}
	var data [20][20]int
	reader := bufio.NewReader(f)
	for i := 0; i < 20; i++ {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		line = strings.TrimSpace(line)
		pieces := strings.Split(line, " ")
		for key, value := range pieces {
			data[i][key], _ = strconv.Atoi(value)
		}
	}

	// Store largest product.
	var max int

	// Test horizontal.
	for i := 0; i < 20; i++ {
		for j := 0; j < 17; j++ {
			product := data[i][j] * data[i][j+1] * data[i][j+2] * data[i][j+3]
			if product > max {
				max = product
			}
		}
	}

	// Test vertical.
	for i := 0; i < 20; i++ {
		for j := 0; j < 17; j++ {
			product := data[j][i] * data[j+1][i] * data[j+2][i] * data[j+3][i]
			if product > max {
				max = product
			}
		}
	}

	// Test upper left to lower right diagonal.
	for i := 0; i < 17; i++ {
		for j := 0; j < 17; j++ {
			product := data[i][j] * data[i+1][j+1] * data[i+2][j+2] * data[i+3][j+3]
			if product > max {
				max = product
			}
		}
	}

	// Test upper right to lower left diagonal.
	for i := 0; i < 17; i++ {
		for j := 3; j < 20; j++ {
			product := data[i][j] * data[i+1][j-1] * data[i+2][j-2] * data[i+3][j-3]
			if product > max {
				max = product
			}
		}
	}

	// Print maximum.
	fmt.Println(max)
}
