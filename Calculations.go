package main

import (
	"strings"
)

var CritPath []*vertex
var OriginalCritPath []*vertex
func Calculatefloats(v []*vertex){
	for _,x := range v{
		x.gain = x.late-x.cost-x.early
	}
}
/*
func CalculateCritPath(v *vertex){
	if len(v.edges)==0{
		fmt.Print(v.name+" END!\n")
		CritPath = append(CritPath, v)
	} else {
		int := 0
		for i,x := range v.edges{
			if v.late == x.late-x.cost{
				fmt.Print(v.name)
				fmt.Print(" - ")
				int = i
				CritPath = append(CritPath, v)


			}
		}
		CalculateCritPath(v.edges[int])
	}
}
*/

func StringCrit(v *vertex, res string)string{
	res+=v.name
	if len(v.edges)==0{
		return res
	} else {
		crit_index := 0
		for i,x := range v.edges{
			if v.late == x.late-x.cost{
				crit_index = i
			}
		}
		return StringCrit(v.edges[crit_index], res)
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

func CalculateDrag(todrag *vertex, start *vertex){
	if len(todrag.edges)!=0{
		int := 0
		for i,x := range todrag.edges{
			if todrag.late == x.late-x.cost{
				int = i
			}
		}
		//fmt.Print(todrag.name+"->")
		CalculateDrag(todrag.edges[int],start)
		//fmt.Println("->"+todrag.name)
		drag := 0
		Orig := todrag.cost
		//fmt.Println("Setting orig: "+strconv.Itoa(Orig))
		for IsStillCrit(todrag, start){
			todrag.cost--
			drag++
			if todrag.cost==0{
				break;
			}
		}
		todrag.drag = drag
		todrag.cost = Orig
	} else {
		//fmt.Print(todrag.name)
		//fmt.Println("\nSetting "+todrag.name+" To :"+strconv.Itoa(todrag.cost))
		todrag.drag = todrag.cost
	}
}

func IsStillCrit(v *vertex, start *vertex)bool{
	CalculateEarliest(start, 1)
	CalculateLatest(start)
	Crit := StringCrit(start, "")
	bool := strings.Contains(Crit, v.name)
	//if bool{fmt.Println(v.name+" Was still crit with cost: "+strconv.Itoa(v.cost)+" Path: "+Crit)} else {fmt.Println(v.name+" Was not crit with cost: "+strconv.Itoa(v.cost)+" Path: "+Crit)}
	return bool

}











