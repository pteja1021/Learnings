package main

import (
	"fmt"
	"math"
)
type shape interface{
	area() float64
} 
// this basically means that any type or struct that has all the methods defined in this interface with same return type
// is of type "shape" or implements "shape"
// Hence this interface becomes like a parent to all those structs

type circle struct {
	radius float64
}

type rectangle struct {
	length float64
	breadth float64
}

func (c circle) area() float64{ // since circle struct has method - area which returns float64, it means that it implements the shape interface
	return (math.Pi*c.radius*c.radius)
}

func (r rectangle) area() float64{ // since rectangle struct has method - area which returns float64, it means that it implements the shape interface
	return r.length*r.breadth
}

func Caller(){
	c := circle{4.5}
	r := rectangle{3,4}
	shapes := []shape{c,r} // since c,r both implement shape interface-the type of this slice can be shape
	
	for _, shape := range shapes {
		fmt.Println(shape.area())
		// fmt.Prinltn(shape.radius) - not defined since radius is not defined on shape. It is defined only on circle
	}
}
// objects can implement any number of interfaces
// interfaces can be used in any type now - return type, parameter type, array type- could be anything