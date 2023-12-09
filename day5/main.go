package main

import (
	"fmt"
	"gurkslask/AoC2023/common"
	"strings"
	"sync"
)

// github.com/gurkslask/AoC2023
/*
results = [310461137 1005936853 270062277 506748462 2453327847 105230363 198508730 868993622 99751241 239722700]
ranges = [{from:630335678 to:701491196} {from:260178142 to:385183562} {from:1548082684 to:2067859966} {from:4104586697 to:4135279672} {from:1018893962 to:1429853751} {from:3570781652 to:3615843761} {from:74139777 to:180146500} {from:32
62608046 to:3476068196} {from:3022784256 to:3144777385} {from:2138898608 to:2175668591}]
end_result = 99751241

*/
func main() {
	data := common.ReadFileSlice("day5", false)
	advanced(data)
}

func sum(x int, y int) int { return x + y }

func IsBetweenNumbers(p, source, irange int) bool {
	if p >= source && p <= source+irange {
		return true
	} else {
		return false
	}
}

func RunTheShit(rr r, data []string, c chan int) {
	fmt.Println("Worker started:")
	fmt.Printf("rr = %+v\n", rr)
	map_en := false
	changed := false
	orig_seed := rr.from
	result := 99999999999
	seed := 0

	for orig_seed <= rr.to {
		seed = orig_seed
		for _, row := range data[2:] {
			// If contains :
			if strings.Contains(row, ":") {
				map_en = true
				changed = false
				continue
			} else if len(row) < 1 {
				// If empty
				map_en = false
				continue
			}
			if map_en {
				// fmt.Println(" This is a map")
				// fmt.Printf("row = %+v\n", row)
				instructions := strings.Split(row, " ")
				destination := common.StrToInt(instructions[0])
				source := common.StrToInt(instructions[1])
				irange := common.StrToInt(instructions[2])
				// If seed has changed dont change it again
				if changed {
					continue
				}
				if IsBetweenNumbers(seed, source, irange) {
					delta := source - destination
					seed = seed - delta
					changed = true
					// if key == 1 {
					//if now {
					//fmt.Println("Seed ---------------")
					//fmt.Printf("key = %+v\n", seed)
					//fmt.Printf("seeds[key] = %+v\n", seeds[key])
					//fmt.Printf("destination = %+v\n", destination)
					//fmt.Printf("source = %+v\n", source)
					//fmt.Printf("irange = %+v\n", irange)
					//fmt.Printf("delta = %+v\n", delta)
					//fmt.Printf("seeds_changed = %+v\n", seeds_changed)
					//}
				}
			}

		}
		if seed < result {
			result = seed
		}
		orig_seed += 1
	}
	fmt.Println("Worker done")
	fmt.Printf("rr = %+v\n", rr)
	fmt.Printf("result = %+v\n", result)
	c <- result

}

type r struct {
	from int
	to   int
}

func advanced(data []string) {
	//fmt.Printf("data = %+v\n", data)

	// init seeds
	//seeds := []int{}
	first_split := strings.Split(data[0], ":")
	splits := strings.Split(first_split[1][1:], " ")
	fmt.Printf("splits = %+v\n", splits)
	even := true
	end_result := 999999999
	ranges := []r{}

	for k := range splits {
		if even {
			from := common.StrToInt(splits[k])
			fmt.Printf("from = %+v\n", from)
			iirange := common.StrToInt(splits[k+1])
			fmt.Printf("iirange = %+v\n", iirange)
			//i := 0
			var rr r
			rr.from = from
			rr.to = from + iirange - 1
			ranges = append(ranges, rr)

			/*
				for i < iirange {
					//seeds = append(seeds, from+i)
					result := RunTheShit(from+i, data)
					end_result = common.MinInt(end_result, result)
					i += 1
					if 0%1000 == 0 {
						fmt.Println("1000")
					}
			*/
			even = false

		} else {
			even = true
		}
	}
	fmt.Printf("ranges = %+v\n", ranges)

	c := make(chan int, len(ranges))
	var wg sync.WaitGroup

	var results []int
	fmt.Printf("cap(c) = %+v\n", cap(c))
	for i := 0; i < cap(c); i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			RunTheShit(ranges[i], data, c)
		}()
	}
	wg.Wait()
	close(c)
	for rrr := range c {
		results = append(results, rrr)
	}
	end_result = common.MinInts(results)
	fmt.Printf("results = %+v\n", results)
	fmt.Printf("ranges = %+v\n", ranges)
	fmt.Printf("end_result = %+v\n", end_result)
}
func simple(data []string) {
	//fmt.Printf("data = %+v\n", data)

	// init seeds
	seeds := []int{}
	first_split := strings.Split(data[0], ":")
	for _, num := range strings.Split(first_split[1][1:], " ") {
		seeds = append(seeds, common.StrToInt(num))
	}
	fmt.Printf("seeds = %+v\n", seeds)

	map_en := false
	seeds_changed := []int{}
	for _, row := range data[2:] {
		// If contains :
		if strings.Contains(row, ":") {
			map_en = true
			seeds_changed = []int{}
			continue
		} else if len(row) < 1 {
			// If empty
			map_en = false
			continue
		}
		if map_en {
			// fmt.Println(" This is a map")
			// fmt.Printf("row = %+v\n", row)
			instructions := strings.Split(row, " ")
			destination := common.StrToInt(instructions[0])
			source := common.StrToInt(instructions[1])
			irange := common.StrToInt(instructions[2])
			for key, seed := range seeds {
				// If seed has changed dont change it again
				if common.ContainsInt(seeds_changed, key) {
					continue
				}
				if IsBetweenNumbers(seed, source, irange) {
					delta := source - destination
					seeds[key] = seed - delta
					seeds_changed = append(seeds_changed, key)
					// if key == 1 {
					// fmt.Println("Seed ---------------")
					// fmt.Printf("key = %+v\n", key)
					// fmt.Printf("seeds[key] = %+v\n", seeds[key])
					// fmt.Printf("destination = %+v\n", destination)
					// fmt.Printf("source = %+v\n", source)
					// fmt.Printf("irange = %+v\n", irange)
					// fmt.Printf("delta = %+v\n", delta)
					// fmt.Printf("seeds_changed = %+v\n", seeds_changed)
					// }
				}
			}
		}

	}
	fmt.Printf("seeds = %+v\n", common.MinInts(seeds))
}
