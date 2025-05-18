package main

//interface as a contract 

import(
	"math"
	"fmt"
)

type Shape interface{  // public, to make private use {shape}
  Area() float64
}

type Rectangle struct{
	width, height float64
}

func (r Rectangle) Area() float64{
	return r.width * r.height
}

type Circle struct{
	radius float64
}

func (c Circle) Area() float64{
	return math.pi * c.radius * c.radius
}

func calculateArea(s Shape) float64{
	return s.Area()
}

func main(){
    r := Rectangle{width: 5, height: 4}
	c := Rectangle{width: 2}

	calculateArea(r)
	calculateArea(c)
}