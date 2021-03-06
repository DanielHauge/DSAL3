package main

import (
	"fmt"

)

func main() {

	e := CreateVertex(nil, 20, "E", 0,0,0,0)
	g := CreateVertex([]*vertex{e}, 5, "G", 0, 0, 0,0)
	h := CreateVertex([]*vertex{e}, 15, "H", 0, 0, 0,0)
	d := CreateVertex([]*vertex{e}, 10, "D", 0, 0, 0,0)
	c := CreateVertex([]*vertex{d, g}, 5, "C", 0, 0, 0,0)
	f := CreateVertex([]*vertex{g}, 15, "F", 0, 0, 0,0)
	b := CreateVertex([]*vertex{c}, 20, "B", 0, 0, 0,0)
	a := CreateVertex([]*vertex{f,b,h}, 10, "A", 0, 0, 0,0)

	CalculateEarliest(a, 1)
	CalculateLatest(a)
	fmt.Print("\nCRITICAL PATH: \n")
	fmt.Println(StringCrit(a,""))
	Fullgraph := []*vertex{a,b,c,d,e,f,g,h}
	Calculatefloats(Fullgraph)
	CalculateDrag(a,a)
	for _,x := range Fullgraph{
		x.PrintMe()
	}



}