package main

import (
	"fmt"
	"gurkslask/AoC2023/common"
	re "regexp"
	"strings"
)

// github.com/gurkslask/AoC2023
func main() {
	data := common.ReadFileSlice("day2", false)
	//simple(data)
	adv(data)
}

func sum(x int, y int) int { return x + y }

func simple(data []string) int {
	max_red := 12
	max_green := 13
	max_blue := 14

	result := 0
	re_string := ` *(\d*) (\w*)`
	rere := re.MustCompile(re_string)
	//re_string := `Monkey (.*):\n .*: (.*)\n.*old (.) (.*)\n.*by (.*)\n.*monkey (\d*)\n.*monkey (\d*)`
	//rere := re.MustCompile(re_string)
	//ll := rere.FindStringSubmatch(ss)

	// Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
	/*

		sets := strings.Split(balls, ";")
		fmt.Printf("game = %+v\n", game)
		fmt.Printf("sets = %+v\n", sets)
		for _, set := range sets {
			balls := strings.Split(set, ",")
			for _, ball := range balls {
				fmt.Printf("ball = %+v\n", ball)
				ll := rere.FindStringSubmatch(ball)
				info[ll[0]] = common.StrToInt(ll[1])
			}
		}
	*/
	for _, row := range data {
		info := map[string]int{"red": 0, "blue": 0, "green": 0}
		splitted := strings.Split(row, ":")
		if len(splitted) == 1 {
			break
		}
		game := common.StrToInt(strings.Split(splitted[0], " ")[1])
		fmt.Printf("game = %+v\n", game)
		balls := splitted[1]
		sets := strings.Split(balls, ";")
		for _, set := range sets {

			ballss := strings.Split(set, ",")
			fmt.Printf("ballss = %+v\n", ballss)
			for _, ball := range ballss {
				ll := rere.FindStringSubmatch(ball)
				info[ll[2]] = common.MaxInt(common.StrToInt(ll[1]), info[ll[2]])
			}
		}
		fmt.Printf("info = %+v\n", info)
		if info["red"] <= max_red && info["green"] <= max_green && info["blue"] <= max_blue {
			result += game
			fmt.Printf("Game is valid %v\n", game)
		}

	}
	fmt.Printf("result = %+v\n", result)

	return 0
}

func adv(data []string) int {

	result := 0
	re_string := ` *(\d*) (\w*)`
	rere := re.MustCompile(re_string)

	for _, row := range data {
		info := map[string]int{"red": 0, "blue": 0, "green": 0}
		splitted := strings.Split(row, ":")
		if len(splitted) == 1 {
			break
		}
		game := common.StrToInt(strings.Split(splitted[0], " ")[1])
		fmt.Printf("game = %+v\n", game)
		balls := splitted[1]
		sets := strings.Split(balls, ";")
		for _, set := range sets {

			ballss := strings.Split(set, ",")
			fmt.Printf("ballss = %+v\n", ballss)
			for _, ball := range ballss {
				ll := rere.FindStringSubmatch(ball)
				info[ll[2]] = common.MaxInt(common.StrToInt(ll[1]), info[ll[2]])
			}
		}
		fmt.Printf("info = %+v\n", info)
		this_res := info["red"] * info["blue"] * info["green"]
		fmt.Printf("this_res = %+v\n", this_res)
		result += info["red"] * info["blue"] * info["green"]

	}
	fmt.Printf("result = %+v\n", result)

	return 0
}
