package main

import "fmt"

const limit = 999966663333

func main() {
	sieve := make([]bool, 500000)
	for i := 1; i < 500000; i++ {
		if sieve[i] {
			continue
		}
		num := 2*i + 1
		for j := i + num; j < 500000; j += num {
			sieve[j] = true
		}
	}

	primes := []int{2}
	for key, value := range sieve[1:] {
		if !value {
			primes = append(primes, 2*key+3)
		}
	}
	primes = append(primes, 1000003)

	// Keep track of the sum of all semiprime numbers.
	result := 0

	// Between the square of two primes, we want numbers divisible by one or the
	// other, but not both.
	for i := 0; i < len(primes)-1; i++ {
		a, b, c := primes[i], primes[i+1], primes[i]*primes[i+1]
		start, end := a*a+1, b*b-1

		if start > limit {
			break
		}
		if end > limit {
			end = limit
		}

		// Calculate smallest and largest multiple of a, b, and c in range.
		fA, lA := start-start%a+a, end-end%a
		if start%a == 0 {
			fA -= a
		}
		fB, lB := start-start%b+b, end-end%b
		if start%b == 0 {
			fB -= b
		}
		fC, lC := start-start%c+c, end-end%c
		if start%c == 0 {
			fC -= c
		}

		// Determine number of multiples of a, b, and c in range.
		oA := (lA-fA)/a + 1
		oB := (lB-fB)/b + 1
		oC := (lC-fC)/c + 1

		// Determine the sum of the multiples of a, b, and c in range and do
		// inclusion exclusion.
		count := oA*(lA+fA)/2 + oB*(lB+fB)/2 - oC*(lC+fC)
		result += count
	}

	fmt.Println(result)
}
