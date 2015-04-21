package main

import "fmt"

func legit(str string) bool {
	d1 := str[0] - '0'
	d2 := str[1] - '0'
	d3 := str[2] - '0'
	if d1+d2+d3 > 9 {
		return false
	}
	return true
}

func main() {
	var list []string
	for i := 0; i < 1000; i++ {
		str := fmt.Sprintf("%03d", i)
		if legit(str) {
			list = append(list, str)
		}
	}

	index := make(map[string]int64)
	for _, value := range list {
		index[value[:2]]++
	}

	data := make(map[string]int64)
	for i := 4; i < 20; i++ {
		for key, value := range index {
			for j := 0; j < 10; j++ {
				str := string('0'+j) + key
				if !legit(str) {
					break
				}
				data[str[:2]] += value
			}
		}

		index = make(map[string]int64)
		for key, value := range data {
			index[key] = value
		}

		data = make(map[string]int64)
	}

	var result int64
	for key, value := range index {
		for i := 1; i < 10; i++ {
			str := string('0'+i) + key
			if !legit(str) {
				break
			}
			result += value
		}
	}
	fmt.Println(result)
}
