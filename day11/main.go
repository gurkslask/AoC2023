package main

import (
	"fmt"
	"gurkslask/AoC2023/common"
	re "regexp"
	"strings"
	//re "regexp"
)

//github.com/gurkslask/AoC2023
func main() {
	fmt.Println(common.ReadFileSlice("gotemp", false))
	exampleSlice()
	exampleMap()
	exampleConv()
	fmt.Println(simple())
}

func sum(x int, y int) int {return x + y }

type Monkey struct {
	id int
	items []int
	operation_mod string
	operation_num int
	test_num int
	test_false_monkey int
	test_true_monkey int
	times_inspected int
}

func InitMonkey() Monkey {
	var m Monkey
	m.items = []int{}
	return m
}
func (m Monkey) print() {
	fmt.Printf("m.id = %+v\n", m.id)
	fmt.Printf("m.items = %+v\n", m.items)
	fmt.Printf("m.operation_mod = %+v\n", m.operation_mod)
	fmt.Printf("m.operation_num = %+v\n", m.operation_num)
	fmt.Printf("m.test_num = %+v\n", m.test_num)
	fmt.Printf("m.test_false_monkey = %+v\n", m.test_false_monkey)
	fmt.Printf("m.test_true_monkey = %+v\n", m.test_true_monkey)

}
func (m Monkey) printLite() {
	fmt.Println("-----------")
	fmt.Printf("m.id = %+v\n", m.id)
	fmt.Printf("m.items = %+v\n", m.items)
	fmt.Printf("m.times_inspected = %+v\n", m.times_inspected)
	fmt.Println("-----------")
}

func (m *Monkey) AddItem(i int) {
	mm := *m
	mm.items = append(mm.items, i)
	*m = mm
}
func (m Monkey) ClearItems() Monkey {m.items = []int{}; return m}

func (mm *Monkey) Inspect() {
	m := *mm
	for key := range m.items {
		m.times_inspected += 1
		num := m.items[key]
		new_num := 0
		other_num := 0

		if m.operation_num == -42 {other_num = m.items[key]
		} else {other_num = m.operation_num
		}
		//fmt.Printf("---- \n Before Monkey %v, itemnum: %v ----- \n", m.id, m.items[key])

		if m.operation_mod == "+" {new_num = num + other_num
		} else if m.operation_mod == "-" {new_num = num - other_num
		} else if m.operation_mod == "*" {new_num = num * other_num
		} else if m.operation_mod == "/" {new_num = num / other_num
		}
		//fmt.Printf("---- \n After before div Monkey %v, itemnum: %v ----- \n", m.id,new_num)

		new_num = new_num / 3
		//fmt.Printf("---- \n After before div Monkey %v, itemnum: %v ----- \n", m.id,new_num)
		m.items[key] = new_num
	}
	//fmt.Println("------ Monkeys items")
	//fmt.Printf("m.items = %+v\n", m.items)
	//fmt.Println("------ ")
	*mm = m
}

func (m *Monkey) Throw(mm map[int]Monkey) map[int]Monkey {
	for key := range m.items {
		if m.items[key] % m.test_num == 0 {
			//fmt.Printf("-------\n Throw %v to monkey %v \n------\n", m.items[key], m.test_true_monkey)
			monkeydonkey := mm[m.test_true_monkey]
			monkeydonkey.AddItem(m.items[key])
			mm[m.test_true_monkey] = monkeydonkey
		} else {
			monkeydonkey := mm[m.test_false_monkey]
			monkeydonkey.AddItem(m.items[key])
			mm[m.test_false_monkey] = monkeydonkey
			//fmt.Printf("-------\n Throw %v to monkey %v \n------\n", m.items[key], m.test_false_monkey)
		}
	}
	m.items = []int{}
	//fmt.Printf("mm[m.test_true_monkey].items = %+v\n", mm[m.test_true_monkey].items)
	//fmt.Printf("m.test_true_monkey = %+v\n", m.test_true_monkey)
	//fmt.Printf("mm[m.test_false_monkey].items = %+v\n", mm[m.test_false_monkey].items)
	//fmt.Printf("m.test_false_monkey = %+v\n", m.test_false_monkey)
	return mm
}

func GetMonkeyBussiness(mm map[int]Monkey) int {
	max1 := 0
	max2 := 0
	for key := range mm {
		if mm[key].times_inspected > max1 { max1 =mm[key].times_inspected }
	}
	for key := range mm {
		if mm[key].times_inspected > max2 && mm[key].times_inspected != max1 { max2 = mm[key].times_inspected }
	}
	return max1 * max2

}

func simple() int {
	data := common.ReadFileSlice("day11", false)
	re_string := `Monkey (.*):\n .*: (.*)\n.*old (.) (.*)\n.*by (.*)\n.*monkey (\d*)\n.*monkey (\d*)`
	//ss := strings.Join(data[0:5], "\n")
	ss := ""
	//fmt.Println(ss)
	count := 0
	monkeys := map[int]Monkey{}
	// Make monkeys
	for _, line := range data {
		if count == 6 {
			count = 0
			//fmt.Printf("ss = %+v\n", ss)
			rere := re.MustCompile(re_string)
			ll := rere.FindStringSubmatch(ss)
			//fmt.Printf("ll[1] = %+v\n", ll[1])
			m := InitMonkey()
			//fmt.Printf("ll[2] = %+v\n", ll[2])
			m.items = common.ConvertSliceStringToInt(strings.Split(ll[2], ", "))
			m.id = common.StrToInt(ll[1])
			//if ll[3] == "*" {m.operation_mod = "multi" } else if ll[3] == "/" {m.operation_mod = "div" }else if ll[3] == "+" {m.operation_mod = "plus" }else if ll[3] == "-" {m.operation_mod = "minus" }
			m.operation_mod = ll[3]
			if ll[4] == "old" {m.operation_num = -42} else { m.operation_num = common.StrToInt(ll[4])}
			m.test_num = common.StrToInt(ll[5])
			m.test_true_monkey = common.StrToInt(ll[6])
			m.test_false_monkey = common.StrToInt(ll[7])
			monkeys[m.id] = m

			ss = ""
		} else {
			count += 1
			ss += line + "\n"
		}
	}
	// print monkeys
	for _, v := range monkeys {
		//fmt.Println("--------------\nAfter assignments")
		v.print()
	}

	// Run
	i := 0
	for i < 20 {
		//fmt.Printf("i = %+v\n", i)
		m := 0
		for m < len(monkeys ){

			mm := monkeys[m]
			mm.Inspect()
			monkeys[m] = mm
			monkeys[m].printLite()

			mmm := monkeys[m]
			monkeys = mmm.Throw(monkeys)
			monkeys[m] = mmm
			monkeys[m].printLite()
			m += 1
		}
		// print monkeys
		for _, v := range monkeys {
			//fmt.Printf("---------\n after run %v \n", i)
			v.printLite()
			//fmt.Println("---------")
		}
		i += 1
	}
	// print monkeys
	for _, v := range monkeys {
		v.printLite()
	}

	return GetMonkeyBussiness(monkeys)
}
