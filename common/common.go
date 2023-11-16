package common

import (
	"fmt"
)



func SumInts(ints []int) int {
	var result int
	for _, val := range ints {
		result += val
	}
	return result
}

func MinInts(ints []int) int {
	result := ints[0]
	for _, val := range ints {
		if val < result {result = val}
	}
	return result
}
func MaxInts(ints []int) int {
	result := ints[0]
	for _, val := range ints {
		if val > result {result = val}
	}
	return result
}

func CheckErr(err error) { if err != nil {fmt.Println(err)} }
