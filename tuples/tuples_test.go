package tuples

import (
	"errors"
	"math"
	"testing"

	"github.com/schapagain/raytracer/utils"
)

var ErrStringToFloatConversion = errors.New("error converting strings to floats")

// TestPointStringRepr creates a point and
// checks if its String() method returns
// the expected string format
func TestPointStringRepr(t *testing.T) {
	p := Point{-3, 1, 9}
	pStringExpected := "(-3.000,1.000,9.000)"
	if p.String() != pStringExpected {
		t.Fatalf("Expected %s, but got %s", pStringExpected, p)
	}
}

// TestVectorStringRepr creates a vector and
// checks if its String() method returns
// the expected string format
func TestVectorStringRepr(t *testing.T) {
	v := Vector{1, 2.0, 3}
	vStringExpected := "<1.000,2.000,3.000>"
	if v.String() != vStringExpected {
		t.Fatalf("Expected %s, but got %s", vStringExpected, v)
	}
}

// TestPointEquality creates points and
// checks if points with same coordinates
func TestPointEquality(t *testing.T) {
	var testCases = []struct {
		name        string
		p1          Point
		p2          Point
		expectEqual bool
	}{
		{
			"varying decimal places", Point{1, 2, 3}, Point{1.00, 2.0, 3}, true,
		}, {
			"containing negative coords", Point{-1, 2, -3}, Point{-1.00, 2.0, -3}, true,
		}, {
			"differing X coord", Point{0.99, 2, 3}, Point{1.00, 2.0, 3}, false,
		}, {
			"differing Y coord", Point{1, 2.11, 3}, Point{1.00, 2.0, 3}, false,
		}, {
			"differing Z coord", Point{1, 2, 3.001}, Point{1.00, 2.0, 3}, false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			if testCase.p1.IsEqualTo(testCase.p2) && !testCase.expectEqual {
				t.Fatalf("Not expected %s to be equal to %s", testCase.p1, testCase.p2)
			}
			if !testCase.p1.IsEqualTo(testCase.p2) && testCase.expectEqual {
				t.Fatalf("Expected %s to be equal to %s", testCase.p1, testCase.p2)
			}
		})
	}
}

// TestVectorEquality creates vectors and
// checks if points with same components
func TestVectorEquality(t *testing.T) {
	var testCases = []struct {
		name        string
		v1          Vector
		v2          Vector
		expectEqual bool
	}{
		{
			"varying decimal places", Vector{1, 2, 3}, Vector{1.00, 2.0, 3}, true,
		}, {
			"containing negative coords", Vector{-1, 2, -3}, Vector{-1.00, 2.0, -3}, true,
		}, {
			"differing X coord", Vector{0.99, 2, 3}, Vector{1.00, 2.0, 3}, false,
		}, {
			"differing Y coord", Vector{1, 2.11, 3}, Vector{1.00, 2.0, 3}, false,
		}, {
			"differing Z coord", Vector{1, 2, 3.001}, Vector{1.00, 2.0, 3}, false,
		}, {
			"approx. equal coords", Vector{1, 2, 3}, Vector{1.0000001, 1.9999999, 3.0000000001}, true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			if testCase.v1.IsEqualTo(testCase.v2) && !testCase.expectEqual {
				t.Fatalf("Not expected %s to be equal to %s", testCase.v1, testCase.v2)
			}
			if !testCase.v1.IsEqualTo(testCase.v2) && testCase.expectEqual {
				t.Fatalf("Expected %s to be equal to %s", testCase.v1, testCase.v2)
			}
		})
	}
}

// TestPointSubtraction creates two points
// and checks if subtracting p1 from p2
// returns a direction vector p2 -> p1
func TestPointSubtraction(t *testing.T) {
	p1 := Point{-1, 2, 4}
	p2 := Point{4, -1, -9.1}
	diffV := p1.Subtract(p2)
	expV := Vector{-5, 3, 13.1}
	if !diffV.IsEqualTo(expV) {
		t.Fatalf("Expected %s, but got %s", expV, diffV)
	}
}

// TestPointMove creates a point and a vector
// and checks if moving the point along the vector
// returns the correct destination point
func TestPointMove(t *testing.T) {
	p := Point{-1, 2, 4}
	v := Vector{4, -1, -9.1}
	destP := p.Move(v)
	expDestP := Point{-1 + 4, 2 - 1, 4 - 9.1}
	if !destP.IsEqualTo(expDestP) {
		t.Fatalf("Expected %s, but got %s", expDestP, destP)
	}
}

// TestPointMoveBack creates a point and a vector
// and checks if moving the point backwards along the vector
// returns the correct destination point
func TestPointMoveBack(t *testing.T) {
	p := Point{-1, 2, 4}
	v := Vector{4, -1, -9.1}
	destP := p.MoveBack(v)
	expDestP := Point{-1 - 4, 2 + 1, 4 + 9.1}
	if !destP.IsEqualTo(expDestP) {
		t.Fatalf("Expected %s, but got %s", expDestP, destP)
	}
}

// TestVectorAddition creates two vectors
// and checks if adding v1 from v2
// returns a resultant vector
func TestVectorAddition(t *testing.T) {
	v1 := Vector{-1, 2, 4}
	v2 := Vector{4, -1, -9.1}
	sumV := v1.Add(v2)
	expV := Vector{3, 1, -5.1}
	if !sumV.IsEqualTo(expV) {
		t.Fatalf("Expected %s, but got %s", expV, sumV)
	}
}

// TestVectorSubtraction creates two vectors
// and checks if subtracting v1 from v2
// returns a difference vector
func TestVectorSubtraction(t *testing.T) {
	v1 := Vector{-1, 2, 4}
	v2 := Vector{4, -1, -9.1}
	diffV := v1.Subtract(v2)
	expV := Vector{-5, 3, 13.1}
	if !diffV.IsEqualTo(expV) {
		t.Fatalf("Expected %s, but got %s", expV, diffV)
	}
}

// TestVectorNegation checks if vectors
// can be negated properly
func TestVectorNegation(t *testing.T) {
	var testCases = []struct {
		name       string
		X, Y, Z    float64
		eX, eY, eZ float64
	}{
		{"all positive whole", 2, 3, 1, -2, -3, -1},
		{"all negative whole", -2, -3, -1, 2, 3, 1},
		{"mixed sign", -2, 3, 1, 2, -3, -1},
		{"mixed sign floats", 2.03, -1.6, 1, -2.03, 1.6, -1},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			v := Vector{c.X, c.Y, c.Z}
			negV := v.Negated()
			expV := Vector{c.eX, c.eY, c.eZ}
			if !negV.IsEqualTo(expV) {
				t.Fatalf("Expected %s, but got %s", expV, negV)
			}
		})
	}

}

// TestVectorScalerMultiplication multiplies vectors
// with scaler values and checks if the vectors are scaled appropriately
func TestVectorScalerMultiplication(t *testing.T) {
	var testCases = []struct {
		name       string
		X, Y, Z    float64
		s          float64
		eX, eY, eZ float64
	}{
		{"all positive with positive", 2, 3, 1, 2, 4, 6, 2},
		{"all positive with negative", 2, 3, 1, -2, -4, -6, -2},
		{"mixed floats with positive", -0.5, 2.33, 3, 5, -2.5, 11.65, 15},
		{"mixed floats with negative", -0.5, 2.33, 3, -5, 2.5, -11.65, -15},
		{"mixed floats with downscale", -0.5, 2.33, 3, 0.5, -0.25, 1.165, 1.5},
		{"mixed floats with zero", -0.5, 2.33, 3, 0, 0, 0, 0},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			v := Vector{c.X, c.Y, c.Z}
			scaledV := v.Multiply(c.s)
			expV := Vector{c.eX, c.eY, c.eZ}
			if !scaledV.IsEqualTo(expV) {
				t.Fatalf("Expected %s, but got %s", expV, scaledV)
			}
		})
	}

}

// TestVectorScalerDivision multiplies vectors
// with scaler values and checks if the vectors are scaled appropriately
func TestVectorScalerDivision(t *testing.T) {
	var testCases = []struct {
		name          string
		X, Y, Z       float64
		d             float64
		eX, eY, eZ    float64
		errorExpected bool
	}{
		{"all positive with positive", 2, 3, 1, 2, 1, 1.5, 0.5, false},
		{"all positive with negative", 2, 3, 1, -2, -1, -1.5, -0.5, false},
		{"mixed floats with positive", -0.5, 2.33, 3, 5, -0.1, 0.466, 0.6, false},
		{"mixed floats with negative", -0.5, 2.33, 3, -5, 0.1, -0.466, -0.6, false},
		{"mixed floats with scaling", -0.5, 2.33, 3, 0.5, -1, 4.66, 6, false},
		{"division by zero", 2, 3, 1, 0, 1, 1.5, 0.5, true},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			v := Vector{c.X, c.Y, c.Z}
			scaledV, err := v.Divide(c.d)
			expV := Vector{c.eX, c.eY, c.eZ}
			if c.errorExpected && err == nil {
				t.Fatalf("Error expected but received none")
			} else if !c.errorExpected {
				if err != nil {
					t.Fatalf("No error expected, but got: %q", err)
				} else {
					if !scaledV.IsEqualTo(expV) {
						t.Fatalf("Expected %s, but got %s", expV, scaledV)
					}
				}
			}

		})
	}

}

// TestVectorMagnitude creates vectors and checks if their
// magnitudes are computed properly
func TestVectorMagnitude(t *testing.T) {
	var testCases = []struct {
		name    string
		X, Y, Z float64
		expM    float64
	}{
		{"all positive coords", 2, 3, 1, math.Sqrt(4 + 9 + 1)},
		{"mixed sign coords", -1.7, -5, 1, math.Sqrt(math.Pow(-1.7, 2) + 25 + 1)},
		{"zero vector", 0, 0, 0.00, 0},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			v := Vector{c.X, c.Y, c.Z}
			magV := v.Magnitude()
			if !utils.FloatEqual(magV, c.expM) {
				t.Fatalf("Expected magnitude of %s to be %f, but got %f", v, c.expM, magV)
			}

		})
	}

}

// TestVectorNormalization creates vectors and checks if their
// normals are computed correctly
func TestVectorNormalization(t *testing.T) {
	var testCases = []struct {
		name             string
		X, Y, Z          float64
		expX, expY, expZ float64
		errorExpected    bool
	}{
		{"all positive coords", 1, 1, 1, 1 / math.Sqrt(3), 1 / math.Sqrt(3), 1 / math.Sqrt(3), false},
		{"zero vector", 0, 0, 0, 0, 0, 0, true},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			v := Vector{c.X, c.Y, c.Z}
			vNorm, err := v.Normalized()
			vExpNorm := Vector{c.expX, c.expY, c.expZ}
			if c.errorExpected && err == nil {
				t.Fatalf("Error expected but received none")
			} else if !c.errorExpected {
				if err != nil {
					t.Fatalf("No error expected, but got: %q", err)
				} else {
					if !vNorm.IsEqualTo(vExpNorm) {
						t.Fatalf("Expected normalized %s to be %s, but got %s", v, vExpNorm, vNorm)
					}
					if vNorm.Magnitude() != 1 {
						t.Fatalf("Expected magnitude of normalized vector %s to be 1, but got %f", vNorm, vNorm.Magnitude())
					}
				}
			}

		})
	}

}

// TestDotProduct creates vectors and
// checks if their dot products are calculated correctly
func TestDotProduct(t *testing.T) {
	var testCases = []struct {
		name string
		v1   Vector
		v2   Vector
		expD float64
	}{
		{"all positive coords", Vector{1, 2, 5}, Vector{3, 1, 4}, 25},
		{"containing negative coords", Vector{-1, 2, -3}, Vector{-1.00, 2.0, -3}, 14},
		{"zero vector", Vector{0, 0, 0}, Vector{4, -3.4, 1}, 0},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			dot := testCase.v1.Dot(testCase.v2)
			if dot != testCase.expD {
				t.Fatalf("Expected %s.%s to be %f, but got %f", testCase.v1, testCase.v2, testCase.expD, dot)
			}
		})
	}
}

// TestCrossProduct creates vectors and
// checks if their cross products are calculated correctly
func TestCrossProduct(t *testing.T) {
	var testCases = []struct {
		name     string
		v1       Vector
		v2       Vector
		expCross Vector
	}{
		{"zero vector", Vector{0, 0, 0}, Vector{4, -3.4, 1}, Vector{0, 0, 0}},
		{"unit-x cross unit-y", Vector{1, 0, 0}, Vector{0, 1, 0}, Vector{0, 0, 1}},
		{"unit-y cross unit-x", Vector{0, 1, 0}, Vector{1, 0, 0}, Vector{0, 0, -1}},
		{"unit-y cross unit-z", Vector{0, 1, 0}, Vector{0, 0, 1}, Vector{1, 0, 0}},
		{"unit-z cross unit-y", Vector{0, 0, 1}, Vector{0, 1, 0}, Vector{-1, 0, 0}},
		{"unit-z cross unit-x", Vector{0, 0, 1}, Vector{1, 0, 0}, Vector{0, 1, 0}},
		{"unit-x cross unit-z", Vector{1, 0, 0}, Vector{0, 0, 1}, Vector{0, -1, 0}},
		{"non unit vectors", Vector{3, -3, 1}, Vector{4, 9, 2}, Vector{-15, -2, 39}},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			cross := testCase.v1.Cross(testCase.v2)
			if !testCase.expCross.IsEqualTo(cross) {
				t.Fatalf("Expected %s X %s to be %s, but got %s", testCase.v1, testCase.v2, testCase.expCross, cross)
			}
		})
	}
}
