package canvas

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/schapagain/raytracer/errors"
)

type Color struct {
	R, G, B, A float64
}

type Canvas interface {
	Width() int
	Height() int
	PixelAt(int, int) (Color, error)
	SetPixelAt(int, int, Color) error
	ToPPM() PPM
}

type canvas struct {
	width, height int
	buffer        []Color
}

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
func (c *canvas) SetPixelAt(x, y int, color Color) error {
	if x >= c.width || y >= c.height {
		return &errors.OutOfBoundsError{Details: fmt.Sprintf("Invalid location (%d,%d)", x, y)}
	}
	c.buffer[y*c.width+x] = color
	return nil
}

// ToPPM creates a new PPM format from the current canvas pixel data
func (c *canvas) ToPPM() PPM {

	maxPPMLineLength := 70
	maxPixelCharLength := len(strconv.Itoa(MaxColorValue))*3 + 3
	numPixelsInPPMLine := int(math.Floor(float64(maxPPMLineLength / maxPixelCharLength)))
	numPPMLines := int(math.Ceil(float64(len(c.buffer) / int(numPixelsInPPMLine))))

	ppmLines := make([]string, numPPMLines)
	idx := 0
	for {
		if idx >= numPPMLines {
			break
		}
		currPPMLinePixels := c.buffer[idx*numPixelsInPPMLine : idx*numPixelsInPPMLine+numPixelsInPPMLine]
		currPPMLine := strings.Builder{}
		for i := range currPPMLinePixels {
			color := currPPMLinePixels[i]
			r := int(math.Max(math.Min(float64(MaxColorValue), color.R*float64(MaxColorValue)), 0))
			g := int(math.Max(math.Min(float64(MaxColorValue), color.G*float64(MaxColorValue)), 0))
			b := int(math.Max(math.Min(float64(MaxColorValue), color.B*float64(MaxColorValue)), 0))
			currPPMLine.WriteString(fmt.Sprintf("%d %d %d ", r, g, b))
		}
		ppmLines = append(ppmLines, currPPMLine.String())
		idx++
	}
	p := NewPPM(c.height, c.width, MaxColorValue, &ppmLines)
	return p
}

// NewCanvas returns an empty canvas initialized with Color{0,0,0,0}
func NewCanvas(width, height int) Canvas {
	newCanvas := canvas{width, height, make([]Color, width*height)}
	return &newCanvas
}

// String returns the string representation
// of c
func (c Color) String() string {
	return fmt.Sprintf("(%.3f,%.3f,%.3f,%.3f)", c.R, c.G, c.B, c.A)
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
func (c1 Color) Scale(s float64) Color {
	return Color{c1.R * s, c1.G * s, c1.B * s, c1.A * s}
}
