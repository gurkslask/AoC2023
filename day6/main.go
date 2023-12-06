package main

import (
	"fmt"
	"gurkslask/AoC2023/common"
	"strings"
)

// github.com/gurkslask/AoC2023
func main() {
	data := common.ReadFileSlice("day6", false)
	advanced(data)
}

func sum(x int, y int) int { return x + y }

type rec struct {
	time   int
	record int
}

func advanced(data []string) {
	fmt.Printf("data = %+v\n", data)
	times := strings.Replace(strings.Trim(strings.Split(data[0], ":")[1], " "), " ", "", -1)
	itimes := common.StrToInt(times)
	fmt.Printf("times = %+v\n", itimes)

	dist := strings.Replace(strings.Trim(strings.Split(data[1], ":")[1], " "), " ", "", -1)
	idist := common.StrToInt(dist)
	fmt.Printf("idist = %+v\n", idist)

	records := []rec{}
	var r rec
	r.time = itimes
	r.record = idist
	records = append(records, r)

	result := 1
	for _, r := range records {
		i := 0
		distance := 0
		success := 0
		for i < r.time {
			distance = i * (r.time - i)
			if distance > r.record {
				success += 1
			}
			i += 1
		}
		fmt.Printf("success = %+v\n", success)
		result *= success
	}
	fmt.Printf("result = %+v\n", result)
}
func simple(data []string) {
	fmt.Printf("data = %+v\n", data)
	times := strings.Split(strings.Trim(strings.Split(data[0], ":")[1], " "), " ")
	itimes := common.StringsToInt(times)
	fmt.Printf("times = %+v\n", itimes)

	dist := strings.Split(strings.Trim(strings.Split(data[1], ":")[1], " "), " ")
	idist := common.StringsToInt(dist)
	fmt.Printf("idist = %+v\n", idist)

	records := []rec{}
	for k := range itimes {
		var r rec
		r.time = itimes[k]
		r.record = idist[k]
		records = append(records, r)
	}

	result := 1
	for _, r := range records {
		i := 0
		distance := 0
		success := 0
		for i < r.time {
			distance = i * (r.time - i)
			if distance > r.record {
				success += 1
			}
			i += 1
		}
		fmt.Printf("success = %+v\n", success)
		result *= success
	}
	fmt.Printf("result = %+v\n", result)
}
