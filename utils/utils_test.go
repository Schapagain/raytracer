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
