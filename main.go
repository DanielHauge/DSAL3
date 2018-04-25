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
	graph := []*vertex{a,b,c,d,e,f,g,h}


	CalculateEarliest(a, 1)
	CalculateLatest(a)

	fmt.Print("\n\n CRITICAL PATH: \n")
	CalculateCritPath(a)
	fmt.Print("\n")

	Calculatefloats(graph)
	CalculateDrag(graph)
	for _,x := range graph{
		x.PrintMe()
	}

}