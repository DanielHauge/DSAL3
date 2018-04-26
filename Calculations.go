package main

import (
	"fmt"
	"strconv"
)

func Calculatefloats(v []*vertex){
	for _,x := range v{
		x.gain = x.late-x.cost-x.early
	}
}

func CalculateCritPath(v *vertex)[]*vertex{
	results := []*vertex{}
	if len(v.edges)==0{
		fmt.Print(v.name+" END!\n")
	} else {
		int := 0
		for i,x := range v.edges{
			if v.late == x.late-x.cost{
				fmt.Print(v.name)
				fmt.Print(" - ")
				int = i
				results = append(results, v)
			}
		}
		CalculateCritPath(v.edges[int])
	}
	fmt.Println(strconv.Itoa(len(results)) + " Length")
	return results
}


func CalculateEarliest(vertex *vertex, earliest int){
	if vertex.early <= earliest{
		vertex.early = earliest
	}
	for _, x := range vertex.edges{
		CalculateEarliest(x, vertex.early+vertex.cost)
	}
}

func CalculateLatest(vertex *vertex){
	for _, x := range vertex.edges{
		CalculateLatest(x)
	}
	vertex.late = vertex.GetLatestForSingleVertex()
}

func (v *vertex) GetLatestForSingleVertex()int{
	if len(v.edges) == 0{
		return v.early+v.cost
	}
	result := 100
	for _, x := range v.edges{
		if x.late-x.cost<result{
			result = x.late-x.cost
		}
	}
	return result
}

func CalculateDrag(v []*vertex){
	for _, x := range v{
		originalcost := x.cost
		drag := 0
		fmt.Println("trying to find drag for: "+x.name+" Where its cost is: "+strconv.Itoa(originalcost))
		for x.IsStillCritPath(originalcost-drag, v){
			fmt.Println("Was still critical path: " + strconv.Itoa(drag))
			drag ++;
			if drag == originalcost {
				break;
			}
		}
		fmt.Println("Found Drag for "+x.name+" - Was: "+strconv.Itoa(drag))
		x.drag = drag
		x.cost = originalcost
	}
}

func (v *vertex) IsStillCritPath(x int,a []*vertex)bool{
	v.cost = x
	list := CalculateCritPath(a[0])
	fmt.Println(len(a))
	for _, a := range list {
		fmt.Println(a.name+" - "+v.name)
		if a.name == v.name {
			return true
		}
	}

	return false
}


