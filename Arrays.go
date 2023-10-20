package main

import (
	"fmt"
	"math"
)

type point struct {
	X float64
	Y float64
}

func (p *point) SetXY(x, y float64) {
	p.X = x
	p.Y = y
}
type Round struct{
	Radius float64
}
type Kvadrat struct{
	A float64
}
func (f Kvadrat) Square() float64{
	return f.A*f.A
} 
func (f Round) Square() float64{
	return f.Radius *f.Radius * math.Pi
}
type Fidure interface{
	Square() float64
}
func culcSquare(f Fidure) float64{
	return f.Square()
}
func main() {
	p := point{X: 2, Y: 2}
	p.SetXY(4,4)
	fmt.Println(p)
	k := Kvadrat{A:4}
	s := Round{Radius: 1}
	fmt.Println(culcSquare(k), culcSquare(s))
}
