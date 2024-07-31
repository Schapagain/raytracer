package canvas

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/schapagain/raytracer/errors"
	"github.com/schapagain/raytracer/utils"
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

// clipColor clips R,G,B channels in color to fit in the low -> high range (inclusive)
func clipColor(color *Color, lo, hi int) {
	color.R = float64(utils.MaxInt(utils.MinInt((hi), int(color.R)), lo))
	color.G = float64(utils.MaxInt(utils.MinInt((hi), int(color.G)), lo))
	color.B = float64(utils.MaxInt(utils.MinInt((hi), int(color.B)), lo))
}

// ToPPM creates a new PPM format from the current canvas pixel data
func (c *canvas) ToPPM() PPM {
	maxPixelCharLength := len(strconv.Itoa(MaxColorValue))*3 + 3
	numPixelsInPPMLine := int(math.Floor(float64(MaxPPMLineLength / maxPixelCharLength)))
	numPPMLines := int(math.Ceil(float64(len(c.buffer)) / float64(numPixelsInPPMLine)))
	ppmLines := make([]string, numPPMLines)
	for idx := 0; idx < numPPMLines; idx++ {
		currPPMLinePixels := c.buffer[idx*numPixelsInPPMLine : utils.MinInt(idx*numPixelsInPPMLine+numPixelsInPPMLine, len(c.buffer))]
		currPPMLine := strings.Builder{}
		for i := range currPPMLinePixels {
			color := currPPMLinePixels[i].Scale(float64(MaxColorValue))
			clipColor(&color, 0, MaxColorValue)
			pixelTemplate := " %d %d %d"
			if i == 0 {
				pixelTemplate = "%d %d %d"
			}
			currPPMLine.WriteString(fmt.Sprintf(pixelTemplate, int(color.R), int(color.G), int(color.B)))
		}
		ppmLines[idx] = currPPMLine.String()
	}
	return NewPPM(c.width, c.height, MaxColorValue, &ppmLines)
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
