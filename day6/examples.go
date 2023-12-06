package main

import (
	"fmt"
	"gurkslask/AoC2023/common"
	"strconv"

	lop "github.com/samber/lo"
	"golang.org/x/exp/maps"
)
func exampleSlice() {
	s := []string{"hej", "då", "dårå"}
	s = append(s, "tillägg")
	for _, v := range s[:] {
		fmt.Println(v)
	}
}

func exampleMap() {
	m := make(map[string]int)
	m["alex"] = 35
	m["Jeanette"] = 36
	for k, v := range m {
		fmt.Printf("k: %v, v: %v\n", k, v)
	}
	vals := maps.Values(m)
	fmt.Println(vals)
}

func exampleConv() {
	i := 42
	s := strconv.FormatInt(int64(i), 10)
	fmt.Printf("stringconversion: %v", s)

	ss := "42"
	ii, err := strconv.Atoi(ss)
	common.CheckErr(err)
	fmt.Printf("intconversion: %v", ii)

	is := []int64{1,2,3,4,5,6, 1, 2, 3}
	sis := lop.Map(is, func(x int64, index int) string {
    return strconv.FormatInt(x, 10)
	})
	fmt.Printf("\n Maps: %v", sis)

	uniqs := lop.Uniq(is)
	fmt.Printf("\n Uniques: %v", uniqs)

}

