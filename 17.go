package main

import (
	"fmt"
)

func main() {
	ones := []int{
		0,
		len("one"),
		len("two"),
		len("three"),
		len("four"),
		len("five"),
		len("six"),
		len("seven"),
		len("eight"),
		len("nine"),
	}
	teens := []int{
		len("ten"),
		len("eleven"),
		len("twelve"),
		len("thirteen"),
		len("fourteen"),
		len("fifteen"),
		len("sixteen"),
		len("seventeen"),
		len("eighteen"),
		len("nineteen"),
	}
	tens := []int{
		0,
		0,
		len("twenty"),
		len("thirty"),
		len("forty"),
		len("fifty"),
		len("sixty"),
		len("seventy"),
		len("eighty"),
		len("ninety"),
	}
	hundreds := []int{
		0,
		len("onehundredand"),
		len("twohundredand"),
		len("threehundredand"),
		len("fourhundredand"),
		len("fivehundredand"),
		len("sixhundredand"),
		len("sevenhundredand"),
		len("eighthundredand"),
		len("ninehundredand"),
	}
	sum := 0
	for i := 1; i < 1000; i++ {
		str := fmt.Sprintf("%03d", i)
		sum += tens[str[1]-'0'] + hundreds[str[0]-'0']
		if str[1] == '1' {
			sum += teens[str[2]-'0']
		} else {
			sum += ones[str[2]-'0']
		}
		// Remove "and" if it's 100, 200, ..., 900.
		if str[1] == '0' && str[2] == '0' {
			sum -= 3
		}
	}
	sum += len("onethousand")
	fmt.Println(sum)
}
