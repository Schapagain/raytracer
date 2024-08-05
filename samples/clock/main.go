package main

import (
	"math"

	"github.com/schapagain/raytracer/canvas"
	"github.com/schapagain/raytracer/matrices"
	"github.com/schapagain/raytracer/tuples"
)

func drawMark(c canvas.Canvas, p tuples.Point) {
	block_length := 20
	block_height := 20
	centerX := c.Width() / 2
	centerY := c.Height() / 2
	for i := 0; i < block_height; i++ {
		for j := 0; j < block_length; j++ {
			c.SetPixelAt(int(p.X)+centerX-j+block_length/2, int(p.Y)+centerY-i+block_height/2, canvas.Color{B: 1})
		}
	}
}

func drawHourMarks(c canvas.Canvas) {
	clock_size := int(float64(c.Width()) * 0.4)
	currentMark := tuples.NewPoint(-float64(clock_size/2), 0, 0)
	rotation := matrices.NewRotationZ(2 * math.Pi / 12)
	for i := 0; i < 12; i++ {
		drawMark(c, currentMark)
		currentMark = matrices.Transform(currentMark, rotation)
	}
}

func main() {
	c := canvas.NewCanvas(500, 500)
	drawHourMarks(c)
	c.ToPPM().Save("clock.ppm")
}
