package main

import (
	"fmt"
	"math"
	"reflect"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width, Height float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	c := Circle{Radius: 5}
	r := Rectangle{Width: 3, Height: 4}

	printShapeDetails(c) // Type: main.Circle, Value: {5}
	printShapeDetails(r) // Type: main.Rectangle, Value: {3 4}
}

func printShapeDetails(s Shape) {
	// Get the reflection Value of the interface variable
	value := reflect.ValueOf(s)

	// Get the reflection Type of the interface variable
	typeOf := reflect.TypeOf(s)

	fmt.Printf("Type: %s, Value: %v\n", typeOf, value)
}
