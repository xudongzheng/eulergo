package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	f, err := os.Open("18.txt")
	if err != nil {
		return
	}

	data := make([][]int, 0)
	reader := bufio.NewReader(f)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		line = strings.TrimSpace(line)
		row := make([]int, 0)
		pieces := strings.Split(line, " ")
		for _, value := range pieces {
			i, err := strconv.Atoi(value)
			if err != nil {
				return
			}
			row = append(row, i)
		}
		data = append(data, row)
	}

	rows := len(data)
	for i := rows - 2; i >= 0; i-- {
		cmns := len(data[i])
		for j := 0; j < cmns; j++ {
			data[i][j] += max(data[i+1][j], data[i+1][j+1])
		}
	}

	fmt.Println(data[0][0])
}
