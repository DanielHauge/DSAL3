package main

import "fmt"

func CreateVertex(edges []*vertex, cost int, name string, gain int, early int, late int, drag int) *vertex{
	result := new(vertex)
	result.edges = edges
	result.cost = cost
	result.name = name
	result.gain = gain
	result.early = early
	result.late = late
	result.drag = drag
	return result

}

// Adjacency linked list
type vertex struct {
	edges []*vertex
	cost int
	name string
	gain int
	early int
	late int
	drag int
}

func (v vertex) PrintMe(){
	fmt.Print("\n")
	fmt.Print(v.name)
	fmt.Print("\n Gain:  ")
	fmt.Print(v.gain)
	fmt.Print("\n Earliest!   ")
	fmt.Print(v.early)
	fmt.Print(" - ")
	fmt.Print(v.early+v.cost-1)
	fmt.Print("\n Latest!   ")
	fmt.Print(v.late-v.cost)
	fmt.Print(" - ")
	fmt.Print(v.late-1)
	fmt.Print("\n Drag!   ")
	fmt.Print(v.drag)
	fmt.Print("\n Cost!   ")
	fmt.Print(v.cost)

}
