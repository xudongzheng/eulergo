package main

import "fmt"

func main() {
	sieve := make([]bool, 20000001)
	for i := 2; i <= 20000000; i++ {
		if sieve[i] {
			continue
		}
		for j := 2 * i; j <= 20000000; j += i {
			sieve[j] = true
		}
	}

	primes := make([]int, 0)
	for i := 2; i <= 20000000; i++ {
		if !sieve[i] {
			primes = append(primes, i)
		}
	}

	result := 0
	for _, prime := range primes {
		for i := prime; i <= 5000000; i += prime {
			for j := i; j%prime == 0; j /= prime {
				result -= prime
			}
		}
		for i := 15000000 - (15000000 % prime) + prime; i <= 20000000; i += prime {
			for j := i; j%prime == 0; j /= prime {
				result += prime
			}
		}
	}

	fmt.Println(result)
}
