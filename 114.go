package main

import "fmt"

func main() {
	// The number of ways you can end with black.
	db := make([]int, 51)

	// The number of ways you can end with red.
	dr := make([]int, 51)

	db[0] = 1
	db[1] = 1
	db[2] = 1

	for i := 3; i <= 50; i++ {
		db[i] = dr[i-1] + dr[i-2] + dr[i-3] + db[i-3]
		for j := 0; j <= i-3; j++ {
			dr[i] += db[j]
		}
	}

	fmt.Println(db[50] + dr[50])
}
