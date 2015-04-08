package main

import (
	"fmt"
	"strconv"
)

func palindrome(i int) bool {
	str := strconv.Itoa(i)
	l := len(str)
	if str[0] != str[l-1] {
		return false
	}
	if str[1] != str[l-2] {
		return false
	}
	if str[2] != str[l-3] {
		return false
	}
	return true
}

func main() {
	var largest int
	for a := 100; a < 1000; a++ {
		for b := 100; b < 1000; b++ {
			product := a * b
			if palindrome(product) && product > largest {
				largest = product
			}
		}
	}
	fmt.Println(largest)
}
