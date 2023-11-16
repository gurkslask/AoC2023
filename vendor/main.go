package main

import ("fmt";
"errors";
"github.com/gurkslask/AoC2023/vendor/common";
)
//github.com/gurkslask/AoC2023
func main() {
	var m Marker
	m = makeMarker(3, 1, "K")

	var m2 Marker
	m2 = makeMarker(1, 3, "M")

	gg := makeGrid(5,5)
	gg.InitGrid()
	gg.AddMarker(m)
	gg.AddMarker(m2)
	gg.Print()
	
	err := m.MoveMarkerLeft(gg)
	CheckErr(err)
	gg.Print()

	t := []int{10, 20 , 2, 10}

	tt := common.SumInts(t)
	fmt.Println(tt)

	tx := MaxInts(t)
	fmt.Println(tx)

	tm := MinInts(t)
	fmt.Println(tm)
}

func CheckErr(err error) { if err != nil {fmt.Println(err)} }


type Marker struct {
	pos []int // row, col
	letter string
}
func makeMarker(row int, col int, letter string) Marker{
	var m Marker
	m.pos = make([]int, 2)
	m.pos[0] = row
	m.pos[1] = col
	m.letter = letter
	return m
}
func (m *Marker) MoveMarkerLeft(g Grid) error { if m.pos[1] - 1 < 0 { return errors.New("Cant move left") } else { m.pos[1] -= 1; return nil } }
func (m *Marker) MoveMarkerRight(g Grid) error { if m.pos[1] + 1 > g.cols - 1  { return errors.New("Cant move right") } else { m.pos[1] += 1; return nil } }
func (m *Marker) MoveMarkerUp(g Grid) error { if m.pos[0] - 1 < 0 { return errors.New("Cant move up") } else { m.pos[0] -= 1; return nil } }
func (m *Marker) MoveMarkerDown(g Grid) error { if m.pos[0] + 1 > g.rows - 1 { return errors.New("Cant move down") } else { m.pos[0] += 1; return nil } }

type Grid struct {
	g [][]string
	markers []Marker
	rows int
	cols int
}
func makeGrid(rows int, cols int) Grid {
	var g Grid
	g.rows = rows
	g.cols = cols
	return g
}

func (g *Grid) InitGrid() {
	g.g = make([][]string, g.rows)
	row := 0
	for g.rows > row {
		col := 0
		for g.cols > col {
			g.g[row] = append(g.g[row], "O")
			col += 1
		}
		row += 1
	}
}

func (g Grid) Print() {
	//Reset grid
	g.InitGrid()

	// Place markers on inited grid
	for _, v := range g.markers {
		g.g[v.pos[0]][v.pos[1]] = v.letter
	}
	// Print grid
	for k := range g.g {
		fmt.Println(g.g[k])
	}
}

func (g *Grid) AddMarker(marker Marker) {
	// Add markers to the list
	g.markers = append(g.markers, marker)
}

