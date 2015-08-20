package main

import "fmt"

const mod = 1e16

func sieve1() []int {
	sieve := make([]bool, 5000)
	for i := 2; i < 5000; i++ {
		if sieve[i] {
			continue
		}
		for j := 2 * i; j < 5000; j += i {
			sieve[j] = true
		}
	}

	data := make([]int, 0)
	for i := 2; i < 5000; i++ {
		if !sieve[i] {
			data = append(data, i)
		}
	}

	return data
}

func sieve2(s1 []int) []int {
	sum := 0
	for _, value := range s1 {
		sum += value
	}

	sieve := make([]bool, sum)
	for i := 2; i < sum; i++ {
		if sieve[i] {
			continue
		}
		for j := 2 * i; j < sum; j += i {
			sieve[j] = true
		}
	}

	data := make([]int, 0)
	for i := 2; i < sum; i++ {
		if !sieve[i] {
			data = append(data, i)
		}
	}

	return data
}

func mult(s1 []int, n int) []int {
	out := make([]int, len(s1)+n)
	for i := 0; i < len(s1); i++ {
		out[i] = (out[i] + s1[i]) % mod
		out[i+n] = (out[i+n] + s1[i]) % mod
	}
	return out
}

func main() {
	s1 := sieve1()
	s2 := sieve2(s1)

	i := []int{1}
	for _, value := range s1 {
		i = mult(i, value)
	}

	sum := 0
	for _, value := range s2 {
		sum = (sum + i[value]) % mod
	}
	fmt.Println(sum)
}
