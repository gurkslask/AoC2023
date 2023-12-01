package main

import (
	"fmt"
	"gurkslask/AoC2023/common"
	"strconv"
	"strings"
	"unicode"
)

//github.com/gurkslask/AoC2023
func main() {
	fmt.Println(common.ReadFileSlice("day1", true))
	exampleSlice()
	exampleMap()
	exampleConv()
	data := common.ReadFileSlice("day1", false)
	result := simple(data)
	//fmt.Printf("result = %+v\n", result)
	result = adv(data)
	fmt.Printf("Advanced: result = %+v\n", result)
}

func sum(x int, y int) int {return x + y }


func adv(intext []string) int {
	res := 0
	numbers := map[string]int {
		"one":1,
		"two":2,
		"three":3,
		"four":4,
		"five":5,
		"six":6,
		"seven":7,
		"eight":8,
		"nine":9, 
	}
	keys := []string{}
	for key := range numbers {keys = append(keys, key)}
	fmt.Printf("numbers = %+v\n", numbers)

	for _, row := range intext {
		digits := []int{}
		fmt.Printf("row = %+v\n", row)
		// for numtext, num := range numbers {
			// row = strings.Replace(row, numtext, num,-1)
		// }

		for i, c := range row {
			if unicode.IsDigit(c) {
				ci, err := strconv.Atoi(string(c))
				common.CheckErr(err)
				digits = append(digits, ci)
			} else {
				for _, key := range keys {
					if strings.Contains(row[i:common.MinInt(i+len(key), len(row))], key) {digits = append(digits, numbers[key])}
				}
			}
		}
		fmt.Printf("digits = %+v\n", digits)
		if len(digits) >= 1 {
			res += (digits[0] * 10) + digits[len(digits) - 1]
			//##fmt.Printf("res = %+v\n", res)
			//fmt.Println("___________")
		}
	}
	return res

}

func simple(intext []string) int {
	res := 0

	for _, row := range intext {
		digits := []int{}
		for _, c := range row {
			if unicode.IsDigit(c) {
				ci, err := strconv.Atoi(string(c))
				common.CheckErr(err)
				digits = append(digits, ci)
			}
		}
		fmt.Printf("digits = %+v\n", digits)
		if len(digits) >= 1 {
			res += (digits[0] * 10) + digits[len(digits) - 1]
			//##fmt.Printf("res = %+v\n", res)
			//fmt.Println("___________")
		}
	}
	return res

}

