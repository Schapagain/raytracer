package main

import (
	"fmt"

	"github.com/schapagain/raytracer/canvas"
	"github.com/schapagain/raytracer/tuples"
)

func main() {

	p1 := tuples.Point{X: 1, Z: 1}
	v1 := tuples.Vector{Y: 1}
	c := canvas.NewCanvas(400,400)
	firstColor,_ := c.PixelAt(0,0)
	
	fmt.Printf("Starting location: %s\n", p1.String())
	fmt.Printf("Starting velocity: %s\n", v1.String())
	fmt.Printf("Canvas pixel: %s\n", firstColor)
}
