package canvas

import (
	"fmt"

	"github.com/schapagain/raytracer/errors"
)

// Width returns the width of the canvas c
func (c *canvas) Width() int {
	return c.width
}

// Height returns the height of the canvas c
func (c *canvas) Height() int {
	return c.height
}

// PixelAt returns the color of pixel at (x,y) on canvas c
//
// It returns an error if location (x,y) is out of bounds
func (c *canvas) PixelAt(x, y int) (Color, error) {
	if x >= c.width || y >= c.height {
		return Color{}, &errors.OutOfBoundsError{Details: fmt.Sprintf("Invalid location (%d,%d)", x, y)}
	}
	return c.buffer[y*c.width+x], nil
}

// SetPixelAt sets the color of the pixel at (x,y) on the canvas c
//
// It returns an error if location (x,y) is out of bounds
func (c *canvas) SetPixelAt(x, y int, color Color) (bool, error) {

	if x >= c.width || y >= c.height {
		return false, &errors.OutOfBoundsError{Details: fmt.Sprintf("Invalid location (%d,%d)", x, y)}
	}
	c.buffer[y*c.width+x] = color
	return true, nil
}

// NewCanvas returns an empty canvas initialized with Color{0,0,0,0}
func NewCanvas(width, height int) Canvas {
	newCanvas := canvas{width, height, make([]Color, width*height)}
	return &newCanvas
}

// String returns the string representation
// of c
func (c Color) String() string {
	return fmt.Sprintf("(%f,%f,%f,%f)", c.R, c.G, c.B, c.A)
}

// Add add each component of c1 and c2 and
// returns the resulting color
func (c1 Color) Add(c2 Color) Color {
	return Color{c1.R + c2.R, c1.G + c2.G, c1.B + c2.B, c1.A + c2.A}
}

// Subtract subtracts each component of c2 from c1 and
// returns the resulting color
func (c1 Color) Subtract(c2 Color) Color {
	return Color{c1.R - c2.R, c1.G - c2.G, c1.B - c2.B, c1.A - c2.A}
}

// Multiply multiplies each component of c1 and c2 and
// returns the resulting color
func (c1 Color) Multiply(c2 Color) Color {
	return Color{c1.R * c2.R, c1.G * c2.G, c1.B * c2.B, c1.A * c2.A}
}

// Multiply multiplies each component of c1 with s and
// returns the resulting color
func (c1 Color) Scale(s float32) Color {
	return Color{c1.R * s, c1.G * s, c1.B * s, c1.A * s}
}
