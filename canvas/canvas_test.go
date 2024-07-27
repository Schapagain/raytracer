package canvas

import (
	"testing"

	"github.com/schapagain/raytracer/utils"
)

func colorsAreEqual(c1, c2 Color) bool {
	return utils.FloatEqual(c1.R, c2.R) &&
		utils.FloatEqual(c1.G, c2.G) &&
		utils.FloatEqual(c1.B, c2.B) &&
		utils.FloatEqual(c1.A, c2.A)
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
		s    float32
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
