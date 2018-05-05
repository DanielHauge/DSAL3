# 2. Compulsory algorithm Assignment - graphs
This repository is for a hand-in for Software development (PBA) Datastructure and Algorithms course. Daniel (cph-dh136)

## Description
This assignment is about graphs and the critical path in one. The assignment is based on this ressource: [Here](https://i.gyazo.com/0ff3ad9cf15ed2679dab93dc3f544200.png)

## Assignment
First is to create the datastructure (Graph). I've done this with Adjacency list, ie. Having each node have an array with each its connections. Reasoning for this, is that i wanted to have more properties per node. This will make it quite handy. This is how i create nodes:
```go
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
```
See [Vertex.go](https://github.com/DanielHauge/DSAL3/blob/master/vertex.go)


Now to be able to calculate on a graph, for allmost all calculations i've used recursion to traverse the graph, this ensures all paths are explored and calculated correctly. Here is a snippet of one of the calculations:
```go
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
```
This is a snippet of CalculateLatest. It is traversing all the way to the end as long as the given node has edges forward. 
If not, it will then set the latest for the last node in the graph, and then move backwards and using the "next" nodes latest property to calculate it's own. From there it will go all the way back to the start and setting late properties on the way back to the start.
It is using a helper method to make it more easier to write the code. It is the GetLatestForSingleVertex, it is a function that looks at all it's next nodes and finds critical latest node, and basing it's own latest on that.

There is also functions for:
- Calculating earliest start
- Calculating critical path
- Calculating total float
- Calculating drag

See [Calculations.go](https://github.com/DanielHauge/DSAL3/blob/master/Calculations.go)

With all the Algorithms and Datastructures in place, we can now Initialize and run a test example. I've used the same graph example from the picture in the assignment.
```go
e := CreateVertex(nil, 20, "E", 0,0,0,0)
	g := CreateVertex([]*vertex{e}, 5, "G", 0, 0, 0,0)
	h := CreateVertex([]*vertex{e}, 15, "H", 0, 0, 0,0)
	d := CreateVertex([]*vertex{e}, 10, "D", 0, 0, 0,0)
	c := CreateVertex([]*vertex{d, g}, 5, "C", 0, 0, 0,0)
	f := CreateVertex([]*vertex{g}, 15, "F", 0, 0, 0,0)
	b := CreateVertex([]*vertex{c}, 20, "B", 0, 0, 0,0)
	a := CreateVertex([]*vertex{f,b,h}, 10, "A", 0, 0, 0,0)
``` 
Note: That we do not want to cheat and set the properties, hence the algorithms would be meaningless. So therefor i just set all properties to 0, as should represent not having been calculated yet.

Now we can calculate all the properties:
```go
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
```
This will go through the graph and calculate all the properties. Note: that the order in which the properties are calculated are very important. It is impossible to calculate drag without knowing earliest and latest. It is also impossible to calculate drag with having correct earliest and latest, and so on.

This is the output of the program, with given graph:

```
CRITICAL PATH: 
ABCDE

A
 Gain:  0
 Earliest!   1 - 10
 Latest!   1 - 10
 Drag!   10
 Cost!   10
B
 Gain:  0
 Earliest!   11 - 30
 Latest!   11 - 30
 Drag!   16
 Cost!   20
C
 Gain:  0
 Earliest!   31 - 35
 Latest!   31 - 35
 Drag!   5
 Cost!   5
D
 Gain:  0
 Earliest!   36 - 45
 Latest!   36 - 45
 Drag!   5
 Cost!   10
E
 Gain:  0
 Earliest!   46 - 65
 Latest!   46 - 65
 Drag!   20
 Cost!   20
F
 Gain:  15
 Earliest!   11 - 25
 Latest!   26 - 40
 Drag!   0
 Cost!   15
G
 Gain:  5
 Earliest!   36 - 40
 Latest!   41 - 45
 Drag!   0
 Cost!   5
H
 Gain:  20
 Earliest!   11 - 25
 Latest!   31 - 45
 Drag!   0
 Cost!   15
Process finished with exit code 0
```

## Comments and notes to results
The values are mostly the same as in the picture of the assignment, which indicates the algorithms are correctly calculating the different properties. It should be noted, that for the calculations of the drag, it is very dependend of which criticalpath it finds, as the algorithm is taking only one path as the criticalpath, even it there exists more.
Then it might believe that it is still the criticalpath even though there is more critical paths, then depend on which position in the edges of the previus node it is, be chosen or not to be criticalpath, therefor it can depend on it's positon to calculate the actual drag or it's drag+1. This is why B has a drag of 16, because it's position on A's edges is the last valid criticalpath, even if F path would be valid too, it is then taking the B path. So B needs to be lowered one more time before its forcing F as the only criticalpath, and then recognizing that B is no longer the critical path.

See StringCrit function in [Calculations.go](https://github.com/DanielHauge/DSAL3/blob/master/Calculations.go)

These algorithms is theorized to work on any graph that is acyclic and has 1 finish node and 1 entry node.