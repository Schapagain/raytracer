package matrices

import (
	"math"

	"github.com/schapagain/raytracer/tuples"
)

type Transformation interface {
	Inverse() Transformation
	String() string
	Operator() Matrix
}

type transformation struct {
	operator Matrix
}

func (t *transformation) Operator() Matrix {
	return t.operator
}

func (t *transformation) Inverse() Transformation {
	invMat, _ := t.operator.Inverse()
	return &transformation{invMat}
}

func (t *transformation) String() string {
	return t.operator.String()
}

// NewTranslation returns a matrix operator that translates
// by the given x,y,z units in x-,y-, and z- axes respectively
func NewTranslation(x, y, z float64) Transformation {
	idenMat, _ := NewIdentityMatrix(4)
	idenMat.Set(0, 3, x)
	idenMat.Set(1, 3, y)
	idenMat.Set(2, 3, z)
	return &transformation{idenMat}
}

// NewScaling returns a matrix operator that scales
// by the given x,y,z factors in x-,y-, and z- axes respectively
func NewScaling(x, y, z float64) Transformation {
	idenMat, _ := NewIdentityMatrix(4)
	idenMat.Set(0, 0, x)
	idenMat.Set(1, 1, y)
	idenMat.Set(2, 2, z)
	return &transformation{idenMat}
}

// NewRotationX returns a matrix operator that rotates
// around the x-axis
func NewRotationX(angle float64) Transformation {
	rotMat, _ := NewMatrixFromSlice([][]float64{
		{1, 0, 0, 0},
		{0, math.Cos(angle), -math.Sin(angle), 0},
		{0, math.Sin(angle), math.Cos(angle), 0},
		{0, 0, 0, 1},
	})
	return &transformation{rotMat}
}

// NewRotationY returns a matrix operator that rotates
// around the Y-axis
func NewRotationY(angle float64) Transformation {
	rotMat, _ := NewMatrixFromSlice([][]float64{
		{math.Cos(angle), 0, math.Sin(angle), 0},
		{0, 1, 0, 0},
		{-math.Sin(angle), 0, math.Cos(angle), 0},
		{0, 0, 0, 1},
	})
	return &transformation{rotMat}
}

// NewRotationZ returns a matrix operator that rotates
// around the z-axis
func NewRotationZ(angle float64) Transformation {
	rotMat, _ := NewMatrixFromSlice([][]float64{
		{math.Cos(angle), -math.Sin(angle), 0, 0},
		{math.Sin(angle), math.Cos(angle), 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	})
	return &transformation{rotMat}
}

// NewShear returns a matrix operator that applies
// shear transformation according to the params provided
func NewShear(xy, xz, yx, yz, zx, zy float64) Transformation {
	rotMat, _ := NewMatrixFromSlice([][]float64{
		{1, xy, xz, 0},
		{yx, 1, yz, 0},
		{zx, zy, 1, 0},
		{0, 0, 0, 1},
	})
	return &transformation{rotMat}
}

// Transform applies the provided transformations to tup in order
func Transform[T tuples.Tuple](tup T, transformations ...Transformation) T {
	operator := NewTranslation(0, 0, 0).Operator()
	for i := len(transformations) - 1; i >= 0; i-- {
		operator, _ = operator.Multiply(transformations[i].Operator())
	}
	switch t := any(tup).(type) {
	case tuples.Vector:
		prod, _ := operator.Multiply(NewMatrixFromVector(t))
		prodColumn, _ := prod.GetCol(0)
		return T(tuples.NewVector(prodColumn[0], prodColumn[1], prodColumn[2]))
	case tuples.Point:
		prod, _ := operator.Multiply(NewMatrixFromPoint(t))
		prodColumn, _ := prod.GetCol(0)
		return T(tuples.NewVector(prodColumn[0], prodColumn[1], prodColumn[2]))
	default:
		return T(tuples.NewPoint(0, 0, 0))
	}
}
