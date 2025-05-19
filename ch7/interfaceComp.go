package main

//interface inside interface (embeddings)

import(
	"math"
	"fmt"
)

type Shape interface{  // public, to make private use {shape}
  Area() float64
}

type Measurable interface{
	Perimeter() float64 
}

type Geometry interface{
	Shape,
	Measurable
}

type Rectangle struct{
	width, height float64
}

func (r Rectangle) Perimeter() float64{
	return 2*(r.width + r.height)
}

func (r Rectangle) Area() float64{
	return r.width * r.height
}

func describeShape(g Geometry) float64{
	fmt.Println("Area:", g.Area())
	fmt.Println("Perimeter:", g.Perimeter())

}

func main(){
    r := Rectangle{width: 5, height: 4}

	describeShape(r)
}