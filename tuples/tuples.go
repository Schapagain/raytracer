package tuples

import (
	"fmt"
	"math"

	"github.com/schapagain/raytracer/errors"
	"github.com/schapagain/raytracer/utils"
)

// String returns the string representation of p
func (p Point) String() string {
	formatString := "(%.3f,%.3f,%.3f)"
	return fmt.Sprintf(formatString, p.X, p.Y, p.Z)
}

// IsEqualTo reports whether p1 and p2 are equal
// by performing element-wise float comparison
func (p1 Point) IsEqualTo(p2 Point) bool {
	return utils.FloatEqual(p1.X, p2.X) &&
		utils.FloatEqual(p1.Y, p2.Y) &&
		utils.FloatEqual(p1.Z, p2.Z)
}

// Move moves the point forwards along direction of v
// by |v| units
func (p Point) Move(v Vector) Point {
	return Point{p.X + v.X, p.Y + v.Y, p.Z + v.Z}
}

// MoveBack moves the point backwards along direction of v
// by |v| units
func (p Point) MoveBack(v Vector) Point {
	return Point{p.X - v.X, p.Y - v.Y, p.Z - v.Z}
}

// Subtract returns the direction vector from p2 to p1
func (p1 Point) Subtract(p2 Point) Vector {
	return Vector{p1.X - p2.X, p1.Y - p2.Y, p1.Z - p2.Z}
}

// String returns the string representation of v
func (t Vector) String() string {
	formatString := "<%.3f,%.3f,%.3f>"
	return fmt.Sprintf(formatString, t.X, t.Y, t.Z)
}

// IsEqualTo reports whether v1 and v2 are equal
// by performing element-wise float comparison
func (v1 Vector) IsEqualTo(v2 Vector) bool {
	return utils.FloatEqual(v1.X, v2.X) &&
		utils.FloatEqual(v1.Y, v2.Y) &&
		utils.FloatEqual(v1.Z, v2.Z)
}

// Add returns the resultant vector adding v1 and v2
func (v1 Vector) Add(v2 Vector) Vector {
	return Vector{v1.X + v2.X, v1.Y + v2.Y, v1.Z + v2.Z}
}

// Subtract returns the difference between t1 and t2
func (t1 Vector) Subtract(t2 Vector) Vector {
	return Vector{t1.X - t2.X, t1.Y - t2.Y, t1.Z - t2.Z}
}

// Multiply returns the result of multiplying t with scaler s
func (t Vector) Multiply(s float64) Vector {
	return Vector{s * t.X, s * t.Y, s * t.Z}
}

// Divide returns the result of dividing t with scaler s
//
// Returns an error if division by zero is attempted
func (t Vector) Divide(s float64) (Vector, error) {
	if utils.FloatEqual(0, s) {
		return Vector{}, &errors.DivisionByZeroError{
			Details:fmt.Sprintf("Cannot divide %s by zero", t.String())}
	}
	return Vector{t.X / s, t.Y / s, t.Z / s}, nil
}

// Negated returns the negated version of t
// that has all X,Y, and Z elements negated
func (t Vector) Negated() Vector {
	return Vector{-t.X, -t.Y, -t.Z}
}

// Magnitude returns the magnitude of the vector represented by t
func (t Vector) Magnitude() float64 {
	return math.Sqrt(math.Pow(t.X, 2) + math.Pow(t.Y, 2) + math.Pow(t.Z, 2))
}

// Normalized returns the normalized version of the vector
// represented by t
func (v Vector) Normalized() (Vector, error) {
	return v.Divide(v.Magnitude())
}

// Dot returns the dot product of v1 and v2
func (v1 Vector) Dot(v2 Vector) float64 {
	return v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z
}

// Cross returns the cross product between v1 and v2
func (v1 Vector) Cross(v2 Vector) Vector {
	return Vector{v1.Y*v2.Z - v1.Z*v2.Y, v2.X*v1.Z - v2.Z*v1.X, v1.X*v2.Y - v1.Y*v2.X}
}
