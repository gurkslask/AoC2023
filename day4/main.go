package main

import (
	"fmt"
	"gurkslask/AoC2023/common"
	"strings"
)

// github.com/gurkslask/AoC2023
func main() {
	data := common.ReadFileSlice("day4", false)
	//fmt.Printf("data = %+v\n", data)
	advanced(data)
}

func sum(x int, y int) int { return x + y }

func advanced(data []string) {
	// 774715 - too low
	// 774814 - ?
	// 774815 - ?
	result := 0
	cards := map[int]int{}
	score_per_card := map[int]int{}
	for i := range data {
		cards[i+1] = 1
		score_per_card[i] = 0
	}
	for _, row := range data {
		if len(row) == 0 {
			break
		}
		winning_num := []string{}
		my_num := []string{}
		points := 0

		splitted := strings.Split(row, "|")
		my_num = strings.Split(splitted[1], " ")
		//fmt.Printf("my_num[0][0] = %+v\n", my_num[0])
		//fmt.Printf("my_num = %+v\n", my_num)

		colon_index := strings.Index(splitted[0], ":")
		winning_num = strings.Split(splitted[0][colon_index+1:], " ")
		//fmt.Printf("winning_num = %+v\n", winning_num)

		game := common.StrToInt(splitted[0][strings.Index(splitted[0], " ")+1 : strings.Index(splitted[0], ":")])
		fmt.Printf("game = %+v\n", game)

		hits := 0

		for _, i := range my_num {
			if common.ContainsString(winning_num, i) && common.CheckIfInt(i) {
				hits += 1
				fmt.Printf("i = %+v\n", i)
			}
		}
		//fmt.Printf("hits = %+v\n", hits)
		if hits >= 1 {
			points = 1
		}
		if hits > 1 {
			k := 1
			for k < hits {
				points = points * 2
				k += 1
			}
		}
		score_per_card[game] = points
		o := 0
		for o < cards[game] {
			u := 0
			for u < hits {
				cards[common.MinInt(game+u+1, len(data)+1)] += 1
				u += 1
			}
			result += 1
			o += 1
		}

		fmt.Printf("points = %+v\n", points)

	}
	sresult := 0
	for _, v := range cards {
		sresult += v
	}
	fmt.Printf("sresult = %+v\n", sresult)
	fmt.Printf("cards = %+v\n", cards)
	fmt.Printf("score_per_card = %+v\n", score_per_card)
	fmt.Printf("result = %+v\n", result)

}
func simple(data []string) {
	result := 0
	for _, row := range data {
		if len(row) == 0 {
			break
		}
		winning_num := []string{}
		my_num := []string{}
		points := 0

		splitted := strings.Split(row, "|")
		my_num = strings.Split(splitted[1], " ")
		fmt.Printf("my_num[0][0] = %+v\n", my_num[0])
		fmt.Printf("my_num = %+v\n", my_num)

		colon_index := strings.Index(splitted[0], ":")
		winning_num = strings.Split(splitted[0][colon_index+1:], " ")
		fmt.Printf("winning_num = %+v\n", winning_num)
		hits := 0

		for _, i := range my_num {
			if common.ContainsString(winning_num, i) && common.CheckIfInt(i) {
				hits += 1
				fmt.Printf("i = %+v\n", i)
			}
		}
		fmt.Printf("hits = %+v\n", hits)
		if hits >= 1 {
			points = 1
		}
		if hits > 1 {
			k := 1
			for k < hits {
				points = points * 2
				k += 1
			}
		}

		result += points
		fmt.Printf("points = %+v\n", points)

	}
	fmt.Printf("result = %+v\n", result)

}
