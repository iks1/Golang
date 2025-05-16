package geometry 

import "math"

type Point struct{ X, Y float64 }

//traditional function
func Distance(p,q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

//same thing, but as method of the point type

func (p Point) Distance(q Point) float64{
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

p := Point{1,2}
q := Point{4,5}

fmt.Println(Distance(p,q)) // function call 
fmt.Println(p.Distance(q)) // method

type Path []Point 


// distance traveled along the path 
func (path Path) Distance() float64 {
	sum := 0.0 
	for i := range path {
		if i>0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

path := Path{
	{1,1},
	{5,1},
	{5,4},
	{1,1},
}

fmt.Println(perim.Distance()) // 

