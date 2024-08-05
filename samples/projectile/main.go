package main

import (
	"github.com/schapagain/raytracer/canvas"
	"github.com/schapagain/raytracer/tuples"
)

type Projectile struct {
	Position tuples.Point
	Velocity tuples.Vector
	Color    canvas.Color
}

func projectileSim(projectiles []Projectile, gravity, wind tuples.Vector) {
	c := canvas.NewCanvas(1000, 1000)
	deltaT := 0.001
	tick := func(p *Projectile) {
		p.Position = p.Position.Move(p.Velocity.Multiply(deltaT))
		p.Velocity = p.Velocity.Add(gravity.Multiply(deltaT)).Add(wind.Multiply(deltaT))
	}
	for _, p := range projectiles {
		for {
			if p.Position.X < 0 || p.Position.Y < 0 || p.Position.X > float64(c.Width()) || p.Position.Y > float64(c.Height()) {
				break
			}
			c.SetPixelAt(int(p.Position.X), c.Height()-int(p.Position.Y), p.Color)
			tick(&p)
		}
	}
	c.ToPPM().Save("projectile_path.ppm")
}

func main() {
	projectile1 := Projectile{
		Position: tuples.Point{X: 0, Y: 0},
		Color:    canvas.Color{R: 1},
		Velocity: tuples.Vector{X: 1, Y: 10},
	}

	projectile2 := Projectile{
		Position: tuples.Point{X: 0, Y: 0},
		Color:    canvas.Color{G: 1},
		Velocity: tuples.Vector{X: 3, Y: 10},
	}
	projectile3 := Projectile{
		Position: tuples.Point{X: 0, Y: 0},
		Color:    canvas.Color{B: 1},
		Velocity: tuples.Vector{X: 7, Y: 10},
	}
	gravity := tuples.Vector{Y: -0.1}
	wind := tuples.Vector{X: 0}
	projectileSim([]Projectile{projectile1, projectile2, projectile3}, gravity, wind)
}
