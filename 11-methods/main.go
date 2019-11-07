package main

import (
	"fmt"
	"math"
)

type Coord struct {
	X, Y float64
}

func (v Coord) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Coord) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Coord{3, 4}
	fmt.Println(v.Magnitude())
	v.Scale(10)
	fmt.Println(v.Magnitude())
}
