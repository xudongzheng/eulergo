package main

import (
	"fmt"
	"math"
)

func main() {
	// a=m^2-n^2, b=2mn, c=m^2-n^2, a+b+c=1000=2m^2+2mn, 500=m^2+mn,
	// m^2+mn-500=0, m=[-n+sqrt(n^2+2000)]/2, guess n, test if m is integer and
	// greater.
	for n := float64(0); n < 500; n++ {
		m := (-n + math.Sqrt(n*n+2000)) / 2
		if m <= n {
			continue
		}
		if math.Floor(m) != m {
			continue
		}
		a := int64(m*m - n*n)
		b := int64(2 * m * n)
		c := int64(m*m + n*n)
		fmt.Println(a * b * c)
	}
}
