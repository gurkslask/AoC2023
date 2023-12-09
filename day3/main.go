package main

import (
	"fmt"
	"gurkslask/AoC2023/common"
)

// 554033 - too high
// 555449
// 553825 - from python
// github.com/gurkslask/AoC2023
func main() {
	data := common.ReadFileSlice("day3", true)
	simple(data)
}

func sum(x int, y int) int { return x + y }

func simple(data []string) {
	result := 0
	rows := len(data)
	cols := len(data[0])
	gg := common.MakeGrid(rows, cols)
	gg.InitGrid()

	for r, row := range data {
		for c := range row {
			gg.AddMarker(common.MakeMarker(r, c, string(data[r][c])))
		}
	}

	// gg.Print()
	r := 0
	for r < rows {
		c := 0
		for c < cols {
			// Check if this position is a digit
			if common.CheckIfInt(gg.CheckLoc(r, c)) {
				how_many_nums := 1
				// Check if next position is a digit
				if common.CheckIfInt(gg.CheckLoc(r, c+1)) && how_many_nums == 1 {
					how_many_nums += 1
				}
				// Check if third position is a digit
				if common.CheckIfInt(gg.CheckLoc(r, c+2)) && how_many_nums == 2 {
					how_many_nums += 1
				}
				//if common.CheckIfInt(gg.CheckLoc(r, c+3)) && how_many_nums == 3 {
				//how_many_nums += 1
				//}

				// Initialize int for loop
				numnum := 0
				// Initialize string for number
				num := ""
				// Reset near variable so we now if we are near a symbol
				near := false

				// Check all the numbers, depending on if its 1, 2 or 3
				for numnum < how_many_nums {
					// Add number so we can use it if near
					num += gg.CheckLoc(r, c+numnum)
					// Check actual number
					if check(r, c+numnum, gg) {
						//fmt.Printf("This is near %v \n", gg.CheckLoc(r, c+numnum))

						// Set bool so we now its a part number
						near = true
					}

					// Increment for next number
					numnum += 1
				}
				// If its a part number, add it to the result
				if near {
					result += common.StrToInt(num)
					// fmt.Printf("Adding: %v to result: %v\n", num, result)
					// fmt.Printf("num = %+v\n", num)
					// // // fmt.Printf("how_many_nums = %+v\n", how_many_nums)
				}

				// Increment with how many numbers we used, minus 1 because we will increment outside of the loop
				c += how_many_nums - 1
			}
			c += 1
		}
		r += 1
	}
	fmt.Printf("result = %+v\n", result)

}

func isSymbol(s string) bool {
	if s != "." && !common.CheckIfInt(s) && s != "ö" {
		// fmt.Printf("s = %+v\n", s)
		return true
	} else {
		return false
	}

}
func check(row, col int, g common.Grid) bool {
	if isSymbol(g.CheckUp(row, col)) {
		return true
	} else if isSymbol(g.CheckLeft(row, col)) {
		return true
	} else if isSymbol(g.CheckRight(row, col)) {
		return true
	} else if isSymbol(g.CheckUpRight(row, col)) {
		return true
	} else if isSymbol(g.CheckDownRight(row, col)) {
		return true
	} else if isSymbol(g.CheckUpLeft(row, col)) {
		return true
	} else if isSymbol(g.CheckDownLeft(row, col)) {
		return true
	} else if isSymbol(g.CheckDown(row, col)) {
		return true
	} else {
		return false
	}

}
