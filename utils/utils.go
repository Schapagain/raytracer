package utils

import (
	"math"
)

type float interface {
	float32 | float64
}

// FloatEqual returns whether floats a and b are approximately equal
//
// Typical threshold used is 1e-6
func FloatEqual[T float](a, b T) bool {
	return math.Abs(float64(a)-float64(b)) < FloatDiffThreshold
}

// MinInt returns the minimum integer argument
func MinInt(ints ...int) int {
	minVal := ints[0]
	for _, num := range ints {
		if num < minVal {
			minVal = num
		}
	}
	return minVal
}


// MaxInt returns the maximum integer argument
func MaxInt(ints ...int) int {
	maxVal := ints[0]
	for _, num := range ints {
		if num > maxVal {
			maxVal = num
		}
	}
	return maxVal
}
