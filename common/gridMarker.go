package common

import (
	"errors"
	"fmt"
)

type Marker struct {
	pos []int // row, col
	letter string
}
func MakeMarker(row int, col int, letter string) Marker{
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
func MakeGrid(rows int, cols int) Grid {
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

func (g Grid) CheckLoc(row int, col int) string {
	for _, markers := range g.markers {
		if markers.pos[0] == row && markers.pos[1] == col {return markers.letter}
	}
	return ""

}

func (g Grid) CheckLeft(row int, col int) string { 
	if col == 0 {
		return "ö"
	} else { return g.CheckLoc(row,col - 1) }
}
func (g Grid) CheckRight(row int, col int) string { 
	if col == len(g.g[row]) {
		return "ö"
	} else { return g.CheckLoc(row, col + 1) }
}

func (g Grid) CheckUp(row int, col int) string { 
	if row == 0 {
		return "ö"
	} else { return g.CheckLoc(row - 1, col) }
}
func (g Grid) CheckDown(row int, col int) string { 
	if col == len(g.g[row][col]) {
		return "ö"
	} else { return g.CheckLoc(row + 1, col) }
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

func (g *Grid) GetLenRow() int {return len(g.g)}
func (g *Grid) GetLenCol() int {return len(g.g[0])}
func (g *Grid) AddMarker(marker Marker) {
	// Add markers to the list
	g.markers = append(g.markers, marker)
}

type Alphabet struct {
	a int
	b int
	c int
	d int
	e int
	f int
	g int
	h int
	i int
	j int
	k int
	l int
	m int
	n int
	o int
	p int
	q int
	r int
	s int
	t int
	u int
	v int
	w int
	x int
	y int
	z int
}
func MakeAlphabet() Alphabet {
	var a Alphabet

	a.a = 1
	a.b = 2
	a.c = 3
	a.d = 4
	a.e = 5
	a.f = 6
	a.g = 7
	a.h = 8
	a.i = 9
	a.j = 10
	a.k = 11
	a.l = 12
	a.m = 13
	a.n = 14
	a.o = 15
	a.p = 16
	a.q = 17
	a.r = 18
	a.s = 19
	a.t = 20
	a.u = 21
	a.v = 22
	a.w = 23
	a.x = 24
	a.y = 25
	a.z = 26
	
	return a

}
