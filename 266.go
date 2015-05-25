package main

import (
	"fmt"
	"math"
	"math/big"
	"sort"
)

type item struct {
	key   int
	value float64
}

type sorter struct {
	data []item
}

func (s sorter) Len() int {
	return 1 << 21
}

func (s sorter) Less(i, j int) bool {
	return s.data[i].value < s.data[j].value
}

func (s sorter) Swap(i, j int) {
	s.data[i], s.data[j] = s.data[j], s.data[i]
}

func handle(list []float64, i int) item {
	product := float64(1)
	for j := uint8(0); j < 21; j++ {
		if i&(1<<j) != 0 {
			product *= list[j]
		}
	}
	return item{i, product}
}

func main() {
	sieve := make([]bool, 190)
	for i := 2; i < 190; i++ {
		if sieve[i] {
			continue
		}
		for j := 2 * i; j < 190; j += i {
			sieve[j] = true
		}
	}

	data := make([]float64, 0)
	sqrt := float64(1)
	for i := 2; i < 190; i++ {
		if !sieve[i] {
			data = append(data, float64(i))
			sqrt *= math.Sqrt(float64(i))
		}
	}

	// data should have size 42. list1 will hold half and list2 will hold half.
	list1 := make([]float64, 0, 21)
	for i := 0; i < 42; i += 2 {
		list1 = append(list1, data[i])
	}
	list2 := make([]float64, 0, 21)
	for i := 1; i < 42; i += 2 {
		list2 = append(list2, data[i])
	}

	// Find all possible combinations in both lists.
	data1 := make([]item, 0, 1<<21)
	for i := 0; i < 1<<21; i++ {
		data1 = append(data1, handle(list1, i))
	}
	data2 := make([]item, 0, 1<<21)
	for i := 0; i < 1<<21; i++ {
		data2 = append(data2, handle(list2, i))
	}

	sort.Sort(sorter{data: data1})
	sort.Sort(sorter{data: data2})

	maxI := float64(0)
	max1, max2 := 0, 0
	for i := 0; i < 1<<21; i++ {
		// Given a product in list1. Find the maximum product in list2 such that
		// their combined product is less than sqrt.
		item1 := data1[i]

		// sort.Search() will return the smallest i for which the function
		// returns true.
		j := sort.Search(1<<21, func(j int) bool {
			item2 := data2[1<<21-1-j]
			return (item1.value * item2.value) < sqrt
		})

		if j == 1<<21 {
			break
		}
		item2 := data2[1<<21-1-j]
		product := item1.value * item2.value
		if product > maxI {
			max1, max2 = item1.key, item2.key
			maxI = product
		}
	}

	out := big.NewInt(1)
	for i := uint8(0); i < 21; i++ {
		if max1&(1<<i) != 0 {
			out.Mul(out, big.NewInt(int64(list1[i])))
		}
	}
	for i := uint8(0); i < 21; i++ {
		if max2&(1<<i) != 0 {
			out.Mul(out, big.NewInt(int64(list2[i])))
		}
	}

	// Print final answer.
	fmt.Println(new(big.Int).Mod(out, big.NewInt(int64(1e16))))
}
