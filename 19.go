package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Date(1901, time.January, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2000, time.December, 31, 0, 0, 0, 0, time.UTC)
	count := 0
	for start.Before(end) {
		if start.Weekday() == time.Sunday {
			count++
		}
		start = start.AddDate(0, 1, 0)
	}
	fmt.Println(count)
}
