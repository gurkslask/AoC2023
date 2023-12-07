package main

import (
	"fmt"
	"gurkslask/AoC2023/common"
	"sort"
	"strings"
)

// github.com/gurkslask/AoC2023
// 246410927 - Too high
// 246409899

// Advanced
// 244848487
func main() {
	data := common.ReadFileSlice("day7_adv", false)
	simple(data)
}
func simple(data []string) {
	fmt.Printf("data = %+v\n", data)
	h := []Hand{}
	for _, row := range data {
		d := strings.Split(row, " ")
		var sh Hand
		sh.bid = common.StrToInt(d[1])
		for _, s := range d[0] {
			sh.cards = append(sh.cards, string(s))
		}
		sh.GetValue()

		h = append(h, sh)
	}
	fmt.Printf("len(h) = %+v\n", len(h))
	sort.Sort(Hands(h))
	fmt.Printf("h = %+v\n", h)
	result := 0
	iter := 1
	for _, hand := range h {
		iiter := len(h) - iter + 1
		result += hand.bid * iiter
		//fmt.Println("--------")
		//fmt.Printf("hand.bid = %+v\n", hand.bid)
		//fmt.Printf("iter = %+v\n", iiter)
		//fmt.Println("--------")
		iter += 1
	}
	fmt.Printf("result = %+v\n", result)
}

func sum(x int, y int) int { return x + y }

type Hand struct {
	cards     []string
	bid       int
	cardvalue int
}

type Hands []Hand

func (h Hands) Len() int { return len(h) }

func (h Hands) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h Hands) Less(i, j int) bool {
	// Evaluate card values
	if h[i].cardvalue < h[j].cardvalue {
		return true
	} else if h[i].cardvalue > h[j].cardvalue {
		return false
	} else {
		// if same card values
		for ii := 0; ii < 5; ii++ {
			if alphaValues[h[i].cards[ii]] > alphaValues[h[j].cards[ii]] {
				return true
			} else if alphaValues[h[i].cards[ii]] < alphaValues[h[j].cards[ii]] {
				return false
			}
		}
	}
	return false
}

var values = []int{
	1, //Full house
	2, //Fyrpar
	3, //kåk
	4, //triss
	5, //tvåpar
	6, //ett par
	7, //högst kort
}

func (h *Hand) GetValue() {
	cards := map[string]int{}
	joker := 0
	for _, c := range h.cards {
		cards[c] += 1
		if c == "J" {
			joker += 1
		}
	}

	if joker == 0 {
		if len(cards) == 1 {
			// fempar
			h.cardvalue = 1
		} else if len(cards) == 2 {
			for _, number := range cards {
				if number == 1 || number == 4 {
					//fyrpar
					h.cardvalue = 2
				} else {
					//kåk
					h.cardvalue = 3
				}
			}
		} else if len(cards) == 3 {
			for _, number := range cards {
				if number == 3 {
					//triss
					h.cardvalue = 4
				}
			}
			if h.cardvalue != 4 {
				// tvåpar
				h.cardvalue = 5
			}
		} else if len(cards) == 4 {
			// par
			h.cardvalue = 6
		} else {
			// högt kort
			h.cardvalue = 7
		}
	} else {

		fmt.Printf("cards = %+v\n", cards)
		if len(cards) >= 2 {
			keymax := GetMaxWOJokers(cards)
			cards[keymax] += joker
			delete(cards, "J")
		}
		fmt.Printf("cards = %+v\n", cards)

		if len(cards) == 1 {
			// fempar
			h.cardvalue = 1
		} else if len(cards) == 2 {
			for _, number := range cards {
				if number == 1 || number == 4 {
					//fyrpar
					h.cardvalue = 2
				} else {
					//kåk
					h.cardvalue = 3
				}
			}
		} else if len(cards) == 3 {
			for _, number := range cards {
				if number == 3 {
					//triss
					h.cardvalue = 4
				}
			}
			if h.cardvalue != 4 {
				// tvåpar
				h.cardvalue = 5
			}
		} else if len(cards) == 4 {
			// par
			h.cardvalue = 6
		} else {
			// högt kort
			h.cardvalue = 7
		}
	}
}

func GetMaxWOJokers(inmap map[string]int) string {
	res := 0
	reskey := ""
	for k, v := range inmap {
		if k == "J" {
			continue
		} else if v > res {
			res = v
			reskey = k
		}
	}
	return reskey
}

var alphaValues = map[string]int{
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"T": 10,
	"J": 1,
	"Q": 12,
	"K": 13,
	"A": 14,
}
