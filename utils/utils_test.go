package utils

import "testing"

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
