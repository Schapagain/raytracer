package utils

import "testing"

// TestFloatEqual checks if approximately equal floats are
//
// correctly deemed equal, and floats that have a larger difference
//
// than the threshold are deemed not equal
func TestFloatEqual(t *testing.T) {
	var f1 float64 = 3
	f2 := 2.9999999999
	f3 := 3.00
	f4 := 30.0
	f5 := 2.99
	t.Run("approximately equal floats", func(t *testing.T) {
		if !FloatEqual(f1, f2) || !FloatEqual(f2, f1) {
			t.Fatalf("Expected %f to equal %f", f1, f2)
		}
	})
	t.Run("floats with small difference", func(t *testing.T) {
		if FloatEqual(f2, f5) || FloatEqual(f5, f2) {
			t.Fatalf("Expected %f to not equal %f", f2, f5)
		}
	})
	t.Run("floats with large difference", func(t *testing.T) {
		if FloatEqual(f3, f4) || FloatEqual(f4, f3) {
			t.Fatalf("Expected %f to not equal %f", f3, f4)
		}
	})
}

// TestMinInt checks if the right minimum is returned
// out of all given integers
func TestMinInt(t *testing.T) {
	testCases := []struct {
		name   string
		args   []int
		expMax int
	}{
		{"all positive", []int{5, 1, 6, 2, 1, 5}, 1},
		{"all negative", []int{-5, -1, -6, -2, -1}, -6},
		{"mixed sign", []int{5, -1, 6, -2, -1}, -2},
	}

	for _, testCase := range testCases {

		t.Run(testCase.name, func(t *testing.T) {
			minInt := MinInt(testCase.args...)
			if minInt != testCase.expMax {
				t.Fatalf("Expected min of %v to be %d, but got %d", testCase.args, testCase.expMax, minInt)
			}
		})
	}
}

// TestMaxInt checks if the right maximum is returned
// out of all given integers
func TestMaxInt(t *testing.T) {
	testCases := []struct {
		name   string
		args   []int
		expMax int
	}{
		{"all positive", []int{5, 1, 6, 2, 1, 5}, 6},
		{"all negative", []int{-5, -1, -6, -2, -1}, -1},
		{"mixed sign", []int{5, -1, 6, -2, -1}, 6},
	}

	for _, testCase := range testCases {

		t.Run(testCase.name, func(t *testing.T) {
			minInt := MaxInt(testCase.args...)
			if minInt != testCase.expMax {
				t.Fatalf("Expected max of %v to be %d, but got %d", testCase.args, testCase.expMax, minInt)
			}
		})
	}
}

// TestFloatSlicesEqual checks if slices with float values
// are considered equal iff they have the same length and
// all corresponding floats are equal
func TestFloatSlicesEqual(t *testing.T) {

	testCases := []struct {
		name     string
		a        []float64
		b        []float64
		areEqual bool
	}{
		{"unequal length", []float64{5.01, 1, 6}, []float64{5, 1}, false},
		{"empty slices", []float64{}, []float64{}, true},
		{"one empty slice", []float64{1, 2}, []float64{}, false},
		{"equal elements", []float64{1.00, 2.45}, []float64{1, 2.45}, true},
	}

	for _, testCase := range testCases {

		t.Run(testCase.name, func(t *testing.T) {
			equal := FloatSlicesEqual(testCase.a, testCase.b)
			if equal != testCase.areEqual {
				t.Fatalf("Expected slices %v and %v to be equal: %t, but got: %t", testCase.a, testCase.b, testCase.areEqual, equal)
			}
		})
	}
}

// TestDotProduct checks whether dot products between two float slices
// are calculated correctly
func TestDotProduct(t *testing.T) {
	testCases := []struct {
		name     string
		a        []float64
		b        []float64
		expDot   float64
		expError bool
	}{
		{"unequal length", []float64{5.01, 1, 6}, []float64{5, 1}, 0, true},
		{"empty slices", []float64{}, []float64{}, 0, false},
		{"all positive", []float64{2, 3, 1}, []float64{1, 2, 3}, 11, false},
	}

	for _, testCase := range testCases {

		t.Run(testCase.name, func(t *testing.T) {
			dot, err := Dot(testCase.a, testCase.b)
			if !testCase.expError {
				if err != nil {
					t.Fatalf("No error expected for Dot(%v,%v), but got one: %q", testCase.a, testCase.b, err)
				} else {
					if dot != testCase.expDot {
						t.Fatalf("Expected Dot(%v,%v) to be %f, but got %f", testCase.a, testCase.b, testCase.expDot, dot)
					}
				}
			}
			if testCase.expError && err == nil {
				t.Fatalf("Error expected for Dot(%v,%v), but got none", testCase.a, testCase.b)
			}
		})
	}

}
