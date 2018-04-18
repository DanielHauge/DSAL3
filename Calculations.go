package main

import "fmt"

func Calculatefloats(v []*vertex){
	for _,x := range v{
		x.gain = x.late-x.cost-x.early
	}
}

func CalculateCritPath(v *vertex){

	if len(v.edges)==0{
		fmt.Print(v.name+" END!")
	} else {
		int := 0
		for i,x := range v.edges{
			if v.late == x.late-x.cost{
				fmt.Print(v.name)
				fmt.Print(" - ")
				int = i
			}
		}
		CalculateCritPath(v.edges[int])
	}
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
