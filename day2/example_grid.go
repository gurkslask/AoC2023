package main

import ("fmt";
"gurkslask/AoC2023/common";
)
//github.com/gurkslask/AoC2023
func grid() {
	var m common.Marker
	m = common.MakeMarker(3, 1, "K")

	var m2 common.Marker
	m2 = common.MakeMarker(1, 3, "M")

	gg := common.MakeGrid(5,5)
	gg.InitGrid()
	gg.AddMarker(m)
	gg.AddMarker(m2)
	gg.Print()
	
	err := m.MoveMarkerLeft(gg)
	common.CheckErr(err)
	gg.Print()

	t := []int{10, 20 , 2, 10}

	tt := common.SumInts(t)
	fmt.Println(tt)

	tx := common.MaxInts(t)
	fmt.Println(tx)

	tm := common.MinInts(t)
	fmt.Println(tm)
}

