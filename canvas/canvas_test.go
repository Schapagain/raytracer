package canvas

import (
	"strconv"
	"strings"
	"testing"

	"github.com/schapagain/raytracer/utils"
)

func colorsAreEqual(c1, c2 Color) bool {
	return utils.FloatEqual(c1.R, c2.R) &&
		utils.FloatEqual(c1.G, c2.G) &&
		utils.FloatEqual(c1.B, c2.B) &&
		utils.FloatEqual(c1.A, c2.A)
}

// TestColorStringRepr checks if the string representation
// of a color is valid
func TestColorStringRepr(t *testing.T) {
	c := Color{0.5, 1, 0, 1}
	expS := "(0.500,1.000,0.000,1.000)"
	if c.String() != expS {
		t.Fatalf("Expected %s, but got %s", expS, c.String())
	}
}

// TestColorAddition creates two colors
// and checks if adding c1 and c2
// returns the resulting color
func TestColorAddition(t *testing.T) {
	c1 := Color{-1, 2, 4, 1}
	c2 := Color{4, -1, -9.1, 0}
	sumC := c1.Add(c2)
	expC := Color{3, 1, -5.1, 1}
	if !colorsAreEqual(sumC, expC) {
		t.Fatalf("Expected %s, but got %s", expC, sumC)
	}
}

// TestColorSubtraction creates two colors
// and checks if subtracting c2 from c1
// returns the resulting color
func TestColorSubtraction(t *testing.T) {
	c1 := Color{-1, 2, 4, 1}
	c2 := Color{4, -1, -9.1, 0}
	sumC := c1.Subtract(c2)
	expC := Color{-5, 3, 13.1, 1}
	if !colorsAreEqual(sumC, expC) {
		t.Fatalf("Expected %s, but got %s", expC, sumC)
	}
}

// TestColorMultiplication creates two colors
// and checks if multilying them
// returns the resulting color
func TestColorMultiplication(t *testing.T) {
	c1 := Color{-1, 2, 4, 1}
	c2 := Color{4, -1, -9.1, 0}
	sumC := c1.Multiply(c2)
	expC := Color{-4, -2, -36.4, 0}
	if !colorsAreEqual(sumC, expC) {
		t.Fatalf("Expected %s, but got %s", expC, sumC)
	}
}

// TestColorScale creates colors
// and checks if multilying them with a scaler
// returns the resulting color
func TestColorScale(t *testing.T) {
	testCases := []struct {
		name string
		c    Color
		s    float64
		expC Color
	}{
		{"all positive", Color{1, 1, 1, 1}, 3, Color{3, 3, 3, 3}},
		{"all zero", Color{0, 0, 0, 0}, 2, Color{0, 0, 0, 0}},
		{"negative scale", Color{0.5, 1, 0, 0}, -3, Color{-1.5, -3, 0, 0}},
		{"down scale", Color{0.5, 1, 0, 0}, 0.5, Color{0.25, 0.5, 0, 0}},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			resC := testCase.c.Scale(testCase.s)
			if !colorsAreEqual(resC, testCase.expC) {
				t.Fatalf("Expected %s, but got %s", testCase.expC, resC)
			}
		})

	}
}

// TestNewCanvas creates a new canvas and checks
// if the canvas is properly initialized
func TestNewCanvas(t *testing.T) {
	c := NewCanvas(100, 120)

	t.Run("set height and width", func(t *testing.T) {
		if c.Height() != 120 {
			t.Fatalf("Expected canvas height to be %d, but got %d", 100, c.Height())
		}
		if c.Width() != 100 {
			t.Fatalf("Expected canvas width to be %d, but got %d", 100, c.Width())
		}

	})
	t.Run("read valid pixels", func(t *testing.T) {
		firstPixel, err := c.PixelAt(0, 0)
		if err == nil {
			if !colorsAreEqual(firstPixel, Color{0, 0, 0, 0}) {
				t.Fatalf("Expected first pixel to be %s, but got %s", Color{0, 0, 0, 0}, firstPixel)
			}
		} else {
			t.Fatalf("No error expected accessing pixel at (0,0) but received error: %q", err)
		}
	})

	t.Run("access invalid pixel", func(t *testing.T) {
		_, err := c.PixelAt(600, 600)
		if err == nil {
			t.Fatalf("Expected to get error accessing out of bounds pixel, but got none")
		}
	})
	t.Run("set valid pixel", func(t *testing.T) {
		newColor := Color{1, 0, 0, 1}
		err := c.SetPixelAt(0, 0, newColor)
		if err != nil {
			t.Fatalf("No error expected setting pixel at (0,0), but received error: %q", err)
		} else {
			changedPixel, _ := c.PixelAt(0, 0)
			if !colorsAreEqual(changedPixel, newColor) {
				t.Fatalf("Expected color at (0,0) to be %s, but got %s", newColor, changedPixel)
			}
		}
	})
	t.Run("set invalid pixel", func(t *testing.T) {
		err := c.SetPixelAt(600, 600, Color{})
		if err == nil {
			t.Fatalf("Expected to get error setting out of bounds pixel, but got none")
		}
	})
}

// TestCanvasToPPM creates a new canvas and checks if its
// PPM representation is correct
func TestCanvasToPPM(t *testing.T) {
	c := NewCanvas(2, 3)

	c.SetPixelAt(0, 1, Color{1, 0.5, 0, 1})
	c.SetPixelAt(1, 2, Color{0, 0, 0.5, 1})

	expPPMLines := []string{
		PPMMagic,
		"2 3",
		strconv.Itoa(MaxColorValue),
		"0 0 0 0 0 0 255 127 0 0 0 0 0 0 0",
		"0 0 127",
	}
	expPPMString := strings.Join(expPPMLines, "\n")
	ppmString := c.ToPPM().String()
	if ppmString != expPPMString {
		t.Fatalf("Expected ppm string to be:\n%q\nGot:\n%q:\n", expPPMString, ppmString)
	}

}
