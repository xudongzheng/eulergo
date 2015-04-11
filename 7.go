package main

import "fmt"

func main() {
	data := make([]int64, 10001)
	queue := make(chan int64, 10001)
	top := int64(10003)
	for i := 0; i < 10001; i++ {
		queue <- int64(i + 2)
	}
	for count := 0; count < 10001; {
		num := <-queue
		for i := 0; i < 10001; i++ {
			if data[i] == 0 {
				// num is prime.
				data[count] = num
				count++
				queue <- top
				top++
				break
			} else if num%data[i] == 0 {
				// num is composite.
				queue <- top
				top++
				break
			}
		}
	}
	fmt.Println(data[10000])
}
