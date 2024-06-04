package main

import "fmt"

type Value float64

// Interface returns v's value as an interface{}.
func (v Value) Interface() interface{} {
	return float64(v)
}

func main() {
	v := Value(10)
	y := v.Interface().(float64) // y will have type float64.
	fmt.Println(y)
}
