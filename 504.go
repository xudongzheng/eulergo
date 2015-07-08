package main

import "fmt"

var squares [200 * 100]bool
var precomp [101][101]int

func under(a, b, x, y int) bool {
	return a*y < b*(a-x)
}

func square(a, b, c, d int) bool {
	i := precomp[a][b] + precomp[b][c] + precomp[c][d] + precomp[d][a]
	i += a + b + c + d - 3
	return squares[i]
}

func main() {
	for i := 0; ; i++ {
		square := i * i
		if square > 200*100 {
			break
		}
		squares[square] = true
	}

	for a := 1; a <= 100; a++ {
		for b := 1; b <= 100; b++ {
			for x := 1; x <= 100; x++ {
				for y := 1; y <= 100; y++ {
					if under(a, b, x, y) {
						precomp[a][b]++
					} else {
						y = 101
					}
				}
			}
		}
	}

	count := 0
	for a := 1; a <= 100; a++ {
		for b := 1; b <= 100; b++ {
			for c := 1; c <= 100; c++ {
				for d := 1; d <= 100; d++ {
					if square(a, b, c, d) {
						count++
					}
				}
			}
		}
	}
	fmt.Println(count)
}
