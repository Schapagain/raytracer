// package matrices provides functions to work with matrices
// and methods for various matrix operations
package matrices

import (
	"errors"
	"fmt"
	"math"
	"strings"

	"github.com/schapagain/raytracer/tuples"
	"github.com/schapagain/raytracer/utils"
)

type Matrix interface {
	Get(int, int) (float64, error)
	Set(int, int, float64) error
	Cols() int
	Rows() int
	GetRow(int) ([]float64, error)
	GetCol(int) ([]float64, error)
	String() string
	IsEqualTo(Matrix) bool
	Multiply(Matrix) (Matrix, error)
	Transposed() Matrix
	SubMatrix(int, int) (Matrix, error)
	Det() (float64, error)
	Cofactor(int, int) (float64, error)
	Minor(int, int) (float64, error)
	Inverse() (Matrix, error)
}

type matrix struct {
	data       []float64
	rows, cols int
}

var (
	ErrOutOfBounds          = errors.New("matrices: index is out of bounds")
	ErrInvalidInitialValues = errors.New("matrices: invalid values for initialization")
	ErrDimensionMismatch    = errors.New("matrices: matrix dimensions are not compatible for the operation")
	ErrMatrixNotInvertible  = errors.New("matrics: given matrix is not invertible")
)

// NewMatrix returns a rows X cols matrix initialized with zeros
//
// It returns an error if either rows or cols is less than one
func NewMatrix(rows, cols int) (Matrix, error) {
	if rows < 1 || cols < 1 {
		return nil, ErrInvalidInitialValues
	}
	return &matrix{
		data: make([]float64, rows*cols),
		rows: rows,
		cols: cols,
	}, nil
}

// NewIdentityMatrix returns an identity matrix of the provided dimensions
//
// It returns an error if a new square matrix
// cannot be created with the given dimensions
func NewIdentityMatrix(dimensions int) (Matrix, error) {
	mat, err := NewMatrix(dimensions, dimensions)
	if err != nil {
		return nil, err
	}

	for i := 0; i < dimensions; i++ {
		mat.Set(i, i, 1)
	}
	return mat, nil
}

// NewMatrixFromSlice builds a new matrix from the provided 2D slice
//
// It returns an error if the built matrix
// would have row or column count less than one
func NewMatrixFromSlice(initialValues [][]float64) (Matrix, error) {
	rows := len(initialValues)
	if rows < 1 {
		return nil, ErrInvalidInitialValues
	}
	cols := len(initialValues[0])
	mat, err := NewMatrix(rows, cols)
	if err != nil {
		return nil, err
	}

	for rowNum := 0; rowNum < rows; rowNum++ {
		for colNum := 0; colNum < cols; colNum++ {
			mat.Set(rowNum, colNum, initialValues[rowNum][colNum])
		}
	}

	return mat, nil
}

// NewMatrixFromTuple returns a column matrix representation of Vector v
func NewMatrixFromVector(v tuples.Vector) Matrix {
	mat, _ := NewMatrixFromSlice([][]float64{{v.X}, {v.Y}, {v.Z}, {0}})
	return mat
}

// NewMatrixFromPoint returns a column matrix representation of Point p
func NewMatrixFromPoint(p tuples.Point) Matrix {
	mat, _ := NewMatrixFromSlice([][]float64{{p.X}, {p.Y}, {p.Z}, {1}})
	return mat
}

// String returns the string representation of matrix m
func (m *matrix) String() string {
	s := strings.Builder{}
	for i, val := range m.data {
		if i > 0 && i%m.cols == 0 {
			s.WriteString("\n")
		}
		if i%m.cols != 0 {
			s.WriteString(" ")
		}
		s.WriteString(fmt.Sprintf("%14.3f", val))
	}
	return s.String()
}

// Cols returns the number of columns in matrix m
func (m *matrix) Cols() int {
	return m.cols
}

// Rows returns the number of rows in matrix m
func (m *matrix) Rows() int {
	return m.rows
}

// Get returns the value of row i and column j of matrix m
//
// It returns an error if attempted to get value at location outside of the matrix
// i.e, if i >= m.Rows() OR j >= m.Cols()
func (m *matrix) Get(i, j int) (float64, error) {
	if i >= m.rows || j >= m.cols {
		return 0, ErrOutOfBounds
	}
	return m.data[i*m.cols+j], nil
}

// GetRow returns the ith row in matrix m
//
// It returns an error if ith row doesn't exist in m. i.e, if i >= m.Rows()
func (m *matrix) GetRow(i int) ([]float64, error) {
	if i >= m.Rows() {
		return []float64{}, ErrOutOfBounds
	}
	return m.data[i*m.Cols() : i*m.Cols()+m.Cols()], nil
}

// GetCol returns the jth column in matrix m
//
// It returns an error if jth column doesn't exist in m. i.e, if j >= m.Cols()
func (m *matrix) GetCol(j int) ([]float64, error) {
	if j >= m.Cols() {
		return []float64{}, ErrOutOfBounds
	}
	jCol := make([]float64, m.Rows())

	for rowNum := 0; rowNum < m.Rows(); rowNum++ {
		jCol[rowNum] = m.data[rowNum*m.Cols()+j]
	}

	return jCol, nil

}

// Set sets the given val at row i and col j in matrix m
//
// It returns an error if attempted to set value at location outside of the matrix
// i.e, if i >= m.Rows() OR j >= m.Cols()
func (m *matrix) Set(i, j int, val float64) error {
	if i >= m.rows || j >= m.cols {
		return ErrOutOfBounds
	}
	m.data[i*m.cols+j] = val
	return nil
}

// IsEqualTo compares each value in m1 and m2 and returns if
// values in corresponding locations are equal under float comparison
func (m1 *matrix) IsEqualTo(m2 Matrix) bool {
	areEqual := true
	m2Mat := m2.(*matrix)
	if len(m1.data) != len(m2Mat.data) {
		return false
	}
	for i, val := range m1.data {
		if !utils.FloatEqual(val, m2Mat.data[i]) {
			return false
		}
	}
	return areEqual
}

// Multiply returns the result of multiplying m1 and m2
//
// It returns an error if m1 and m2 are incompatible for matrix multiplication
// i.e, if m1.Cols() != m2.Rows()
func (m1 *matrix) Multiply(m2 Matrix) (Matrix, error) {
	if m1.Cols() != m2.Rows() {
		return &matrix{}, ErrDimensionMismatch
	}
	productMat, _ := NewMatrix(m1.Rows(), m2.Cols())
	for colNum := 0; colNum < productMat.Cols(); colNum++ {
		currCol, _ := m2.GetCol(colNum)
		for rowNum := 0; rowNum < productMat.Rows(); rowNum++ {
			currRow, _ := m1.GetRow(rowNum)
			dot, _ := utils.Dot(currCol, currRow)
			productMat.Set(rowNum, colNum, dot)
		}
	}
	return productMat, nil
}

// Transposed returns a new matrix
// built by swapping rows and columns of matrix m
func (m *matrix) Transposed() Matrix {
	mat, _ := NewMatrix(m.Cols(), m.Rows())
	for i := 0; i < m.Rows(); i++ {
		for j := 0; j < m.Cols(); j++ {
			mVal, _ := m.Get(i, j)
			mat.Set(j, i, mVal)
		}
	}
	return mat
}

// SubMatrix returns a copy of the matrix m with row and col removed
func (m *matrix) SubMatrix(row, col int) (Matrix, error) {
	if row < 0 || row >= m.Rows() || col < 0 || col >= m.Cols() || m.Rows() < 2 || m.Cols() < 2 {
		return nil, ErrOutOfBounds
	}
	mat, _ := NewMatrix(m.Rows()-1, m.Cols()-1)
	ct := 0
	for i := 0; i < m.Rows(); i++ {
		if i != row {
			for j := 0; j < m.Cols(); j++ {
				if j != col {
					mVal, _ := m.Get(i, j)
					mat.Set(int(ct/mat.Cols()), ct%mat.Cols(), mVal)
					ct++
				}
			}
		}
	}
	return mat, nil
}

// Minor returns the determinant of the submatrix of m
// formed after removing row and col
func (m *matrix) Minor(row, col int) (float64, error) {
	subMat, err := m.SubMatrix(row, col)
	if err != nil {
		return 0, err
	}
	return subMat.Det()
}

// Cofactor returns the cofactor of matrix m at row,col
func (m *matrix) Cofactor(row, col int) (float64, error) {
	minor, err := m.Minor(row, col)
	if err != nil {
		return 0, err
	}
	return minor * math.Pow(-1, float64(col+row)), nil
}

// Det returns the determinant of matrix m
func (m *matrix) Det() (float64, error) {
	if m.Rows() != m.Cols() {
		return 0, ErrDimensionMismatch
	}
	if m.Rows() < 2 {
		return m.data[0], nil
	}
	det := 0.0
	for colNum := 0; colNum < m.Cols(); colNum++ {
		cofac, _ := m.Cofactor(0, colNum)
		val, _ := m.Get(0, colNum)
		det += val * cofac
	}
	return det, nil
}

// Inverse returns the inverse of matrix m
func (m *matrix) Inverse() (Matrix, error) {
	det, err := m.Det()
	if err != nil {
		return nil, err
	}
	if utils.FloatEqual(0, det) {
		return nil, ErrMatrixNotInvertible
	}
	invMat, _ := NewMatrix(m.Rows(), m.Cols())
	for i := 0; i < m.Rows(); i++ {
		for j := 0; j < m.Cols(); j++ {
			cof, _ := m.Cofactor(i, j)
			invMat.Set(j, i, cof/det)
		}
	}
	return invMat, nil
}
