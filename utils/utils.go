package utils

import "math"

type float interface {
	float32 | float64
}

func FloatEqual[T float](a, b T) bool {
	return math.Abs(float64(a)-float64(b)) < FloatDiffThreshold
}