package matrices

import (
	"math"
	"testing"

	"github.com/schapagain/raytracer/tuples"
)

// TestNewTranslation translates a 3d point and
// checks if the destination point is returned as expected
func TestNewTranslation(t *testing.T) {
	testCases := []struct {
		name        string
		srcPt       tuples.Point
		translation Transformation
		expDest     tuples.Point
	}{
		{"zero translation", tuples.NewPoint(1, -4, 3.2), NewTranslation(0, 0, 0), tuples.NewPoint(1, -4, 3.2)},
		{"X-axis translation", tuples.NewPoint(1, -4, 3.2), NewTranslation(45, 0, 0), tuples.NewPoint(46, -4, 3.2)},
		{"negative X-axis translation", tuples.NewPoint(1, -4, 3.2), NewTranslation(-45, 0, 0), tuples.NewPoint(-44, -4, 3.2)},
		{"negative X-axis translation with inverse", tuples.NewPoint(1, -4, 3.2), NewTranslation(45, 0, 0).Inverse(), tuples.NewPoint(-44, -4, 3.2)},
		{"Y-axis translation", tuples.NewPoint(1, -4, 3.2), NewTranslation(0, 2.301, 0), tuples.NewPoint(1, -1.699, 3.2)},
		{"negative Y-axis translation", tuples.NewPoint(1, -4, 3.2), NewTranslation(0, -10, 0), tuples.NewPoint(1, -14, 3.2)},
		{"negative Y-axis translation with inverse", tuples.NewPoint(1, -4, 3.2), NewTranslation(0, 10, 0).Inverse(), tuples.NewPoint(1, -14, 3.2)},
		{"Z-axis translation", tuples.NewPoint(1, -4, 3.2), NewTranslation(0, 0, 1.8), tuples.NewPoint(1, -4, 5)},
		{"negative Z-axis translation", tuples.NewPoint(1, -4, 3.2), NewTranslation(0, 0, -3.2), tuples.NewPoint(1, -4, 0)},
		{"negative Z-axis translation with inverse", tuples.NewPoint(1, -4, 3.2), NewTranslation(0, 0, 3.2).Inverse(), tuples.NewPoint(1, -4, 0)},
		{"all axis translation", tuples.NewPoint(1, -4, 3.2), NewTranslation(1, -4.001, 1.8), tuples.NewPoint(2, -8.001, 5)},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			dest := Transform(testCase.srcPt, testCase.translation)
			if !dest.IsEqualTo(testCase.expDest) {
				t.Fatalf("Expected\n%s\nto translate %s to %s, but was translated to %s instead", testCase.translation, testCase.srcPt, testCase.expDest, dest)
			}
		})
	}
}

// TestNewScaling scales 3d points and vectors and checks if they are
// transformed as expected
func TestNewScaling(t *testing.T) {
	pTestCases := []struct {
		name    string
		srcPt   tuples.Point
		scaling Transformation
		expDest tuples.Point
	}{
		{"zeros", tuples.NewPoint(1, -4, 3.2), NewScaling(0, 0, 0), tuples.NewPoint(0, 0, 0)},
		{"scaling with one", tuples.NewPoint(1, -4, 3.2), NewScaling(1, 1, 1), tuples.NewPoint(1, -4, 3.2)},
		{"positive scaling", tuples.NewPoint(1, -4, 3.2), NewScaling(2.5, 5, 1), tuples.NewPoint(2.5, -20, 3.2)},
		{"negative scaling", tuples.NewPoint(1, -4, 3.2), NewScaling(-2.5, -5, -10), tuples.NewPoint(-2.5, 20, -32)},
		{"reflection over x-axis", tuples.NewPoint(12, -4, 3.2), NewScaling(-1, 1, 1), tuples.NewPoint(-12, -4, 3.2)},
		{"reflection over y-axis", tuples.NewPoint(12, -4, 3.2), NewScaling(1, -1, 1), tuples.NewPoint(12, 4, 3.2)},
		{"reflection over z-axis", tuples.NewPoint(12, -4, 3.2), NewScaling(1, 1, -1), tuples.NewPoint(12, -4, -3.2)},
		{"downscaling", tuples.NewPoint(1, -4, 3.2), NewScaling(0.5, 0.25, -0.1), tuples.NewPoint(0.5, -1, -0.32)},
		{"downscaling with inverse", tuples.NewPoint(1, -4, 3.2), NewScaling(2, 4, -10).Inverse(), tuples.NewPoint(0.5, -1, -0.32)},
	}
	for _, testCase := range pTestCases {
		t.Run(testCase.name, func(t *testing.T) {
			dest := Transform(testCase.srcPt, testCase.scaling)
			if !dest.IsEqualTo(testCase.expDest) {
				t.Fatalf("Expected\n%s\nto scale %s to %s, but was scaled to %s instead", testCase.scaling, testCase.srcPt, testCase.expDest, dest)
			}
		})
	}

	vTestCases := []struct {
		name    string
		srcVec  tuples.Vector
		scaling Transformation
		expDest tuples.Vector
	}{
		{"zeros", tuples.NewVector(1, -4, 3.2), NewScaling(0, 0, 0), tuples.NewVector(0, 0, 0)},
		{"scaling with one", tuples.NewVector(1, -4, 3.2), NewScaling(1, 1, 1), tuples.NewVector(1, -4, 3.2)},
		{"positive scaling", tuples.NewVector(1, -4, 3.2), NewScaling(2.5, 5, 1), tuples.NewVector(2.5, -20, 3.2)},
		{"negative scaling", tuples.NewVector(1, -4, 3.2), NewScaling(-2.5, -5, -10), tuples.NewVector(-2.5, 20, -32)},
		{"reflection over x-axis", tuples.NewVector(12, -4, 3.2), NewScaling(-1, 1, 1), tuples.NewVector(-12, -4, 3.2)},
		{"reflection over y-axis", tuples.NewVector(12, -4, 3.2), NewScaling(1, -1, 1), tuples.NewVector(12, 4, 3.2)},
		{"reflection over z-axis", tuples.NewVector(12, -4, 3.2), NewScaling(1, 1, -1), tuples.NewVector(12, -4, -3.2)},
		{"downscaling", tuples.NewVector(1, -4, 3.2), NewScaling(0.5, 0.25, -0.1), tuples.NewVector(0.5, -1, -0.32)},
		{"downscaling with inverse", tuples.NewVector(1, -4, 3.2), NewScaling(2, 4, -10).Inverse(), tuples.NewVector(0.5, -1, -0.32)},
	}
	for _, testCase := range vTestCases {
		t.Run(testCase.name, func(t *testing.T) {
			dest := Transform(testCase.srcVec, testCase.scaling)
			if !dest.IsEqualTo(testCase.expDest) {
				t.Fatalf("Expected\n%s\nto scale %s to %s, but was scaled to %s instead", testCase.scaling, testCase.srcVec, testCase.expDest, dest)
			}
		})
	}
}

// TestNewRotation rotates 3d points and checks if they are
// transformed as expected
func TestNewRotation(t *testing.T) {
	pTestCases := []struct {
		name     string
		srcPt    tuples.Point
		rotation Transformation
		expDest  tuples.Point
	}{
		{"zero rotation", tuples.NewPoint(1, -4, 3.2), NewRotationX(0), tuples.NewPoint(1, -4, 3.2)},
		{"2pi rotation", tuples.NewPoint(1, -4, 3.2), NewRotationX(0), tuples.NewPoint(1, -4, 3.2)},
		{"y-axis to z-axis", tuples.NewPoint(0, 1, 0), NewRotationX(math.Pi / 2), tuples.NewPoint(0, 0, 1)},
		{"z-axis to negative y-axis", tuples.NewPoint(0, 0, 1), NewRotationX(math.Pi / 2), tuples.NewPoint(0, -1, 0)},
		{"negative y-axis to negative z-axis", tuples.NewPoint(0, -1, 0), NewRotationX(math.Pi / 2), tuples.NewPoint(0, 0, -1)},
		{"negative z-axis to y-axis", tuples.NewPoint(0, 0, -1), NewRotationX(math.Pi / 2), tuples.NewPoint(0, 1, 0)},
		{"negative y-axis to x-axis", tuples.NewPoint(0, -1, 0), NewRotationZ(math.Pi / 2), tuples.NewPoint(1, 0, 0)},
		{"x-axis to y-axis", tuples.NewPoint(1, 0, 0), NewRotationZ(math.Pi / 2), tuples.NewPoint(0, 1, 0)},
		{"x-axis to z-axis", tuples.NewPoint(1, 0, 0), NewRotationY(math.Pi / 2).Inverse(), tuples.NewPoint(0, 0, 1)},
		{"z-axis to negative x-axis", tuples.NewPoint(0, 0, 1), NewRotationY(math.Pi / 2).Inverse(), tuples.NewPoint(-1, 0, 0)},
	}
	for _, testCase := range pTestCases {
		t.Run(testCase.name, func(t *testing.T) {
			dest := Transform(testCase.srcPt, testCase.rotation)
			if !dest.IsEqualTo(testCase.expDest) {
				t.Fatalf("Expected\n%s\nto rotate %s to %s, but was rotated to %s instead", testCase.rotation, testCase.srcPt, testCase.expDest, dest)
			}
		})
	}
}


// TestNewShear applies shear transformation to 3d points and checks if they are
// transformed as expected
func TestNewShear(t *testing.T) {
	pTestCases := []struct {
		name     string
		srcPt    tuples.Point
		shear Transformation
		expDest  tuples.Point
	}{
		{"zero shear", tuples.NewPoint(1, -4, 3.2), NewShear(0,0,0,0,0,0), tuples.NewPoint(1, -4, 3.2)},
		{"xy shear", tuples.NewPoint(1, -4, 3.2), NewShear(1,0,0,0,0,0), tuples.NewPoint(-3, -4, 3.2)},
		{"xz shear", tuples.NewPoint(1, -4, 3.2), NewShear(0,1,0,0,0,0), tuples.NewPoint(4.2, -4, 3.2)},
		{"yx shear", tuples.NewPoint(1, -4, 3.2), NewShear(0,0,1,0,0,0), tuples.NewPoint(1, -3, 3.2)},
		{"yz shear", tuples.NewPoint(1, -4, 3.2), NewShear(0,0,0,1,0,0), tuples.NewPoint(1, -0.8, 3.2)},
		{"zx shear", tuples.NewPoint(1, -4, 3.2), NewShear(0,0,0,0,1,0), tuples.NewPoint(1, -4, 4.2)},
		{"zy shear", tuples.NewPoint(1, -4, 3.2), NewShear(0,0,0,0,0,1), tuples.NewPoint(1, -4, -0.8)},
	}
	for _, testCase := range pTestCases {
		t.Run(testCase.name, func(t *testing.T) {
			dest := Transform(testCase.srcPt, testCase.shear)
			if !dest.IsEqualTo(testCase.expDest) {
				t.Fatalf("Expected\n%s\nto shear %s to %s, but was sheared to %s instead", testCase.shear, testCase.srcPt, testCase.expDest, dest)
			}
		})
	}
}

