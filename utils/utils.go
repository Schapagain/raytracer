package utils

import (
	"errors"
	"math"
)

type float interface {
	float32 | float64
}

var (
	ErrLengthMismatch = errors.New("utils: input lengths do not match")
)

// FloatEqual returns whether floats a and b are approximately equal
//
// Typical threshold used is 1e-6
func FloatEqual[T float](a, b T) bool {
	return math.Abs(float64(a)-float64(b)) < FloatDiffThreshold
}

// FloatSlicesEqual returns whether slices a and b are equal.
// Two float slices are equal if they have equal length, and
// all elements are equal up to the FloatDiffThreshold
func FloatSlicesEqual[T float](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !FloatEqual(a[i], b[i]) {
			return false
		}
	}
	return true
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

// Dot returns the dot product between two float slices
//
// It returns an error if slices are of unequal lengths
func Dot[T float](a, b []T) (float64, error) {
	if len(a) != len(b) {
		return 0, ErrLengthMismatch
	}
	dot := 0.0
	for i := range a {
		dot += float64(a[i] * b[i])
	}
	return dot, nil
}
