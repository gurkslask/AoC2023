package common

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)



func SumInts(ints []int) int {
	var result int
	for _, val := range ints {
		result += val
	}
	return result
}

func MinInt(i, j int) int {
	if i < j {return i} else {return j}
}
func MaxInt(i, j int) int {
	if i > j {return i} else {return j}
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

func ReadFileSlice(day string, test bool) []string {
	fmt.Println(os.Getwd())
	path := ""
	if test {
		path = fmt.Sprintf("../%v/test_input.txt", day)
	} else {
		path = fmt.Sprintf("../%v/input.txt", day)
	}

	file, err := os.ReadFile(path)
	CheckErr(err)

	s := strings.Split(string(file), "\n")

	return s
}

func ContainsInt(slice []int, i int) bool {
	res := false
	for _, num := range slice {
		if num == i {res = true; break}
	}
	return res
	

}

func StrToInt(s string) int {
	i, err := strconv.Atoi(s)
	CheckErr(err)
	return i
}
func StrToInt64(s string) int64 {
	i, err := strconv.Atoi(s)
	CheckErr(err)
	return int64(i)
}

func ConvertSliceStringToInt(slice []string) []int {
	si := []int{}
	fmt.Printf("slice = %+v\n", slice)
	for _, num := range slice {
		nn, err := strconv.Atoi(num)
		CheckErr(err)
		si = append(si, nn)
	}
	return si
}
func ConvertSliceStringToInt64(slice []string) []int64 {
	si := []int64{}
	fmt.Printf("slice = %+v\n", slice)
	for _, num := range slice {
		nn, err := strconv.Atoi(num)
		CheckErr(err)
		si = append(si, int64(nn))
	}
	return si
}
