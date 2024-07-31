package matrices

import (
	"errors"
	"fmt"
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
}

type matrix struct {
	data       []float64
	rows, cols int
}

var (
	ErrOutOfBounds          = errors.New("matrices: index is out of bounds")
	ErrInvalidInitialValues = errors.New("matrices: invalid values for initialization")
	ErrDimensionMismatch    = errors.New("matrices: matrix dimensions are not compatible for the operation")
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

// NewMatrixFromVector returns a column matrix representation of v
func NewMatrixFromVector(v tuples.Vector) Matrix {
	mat, _ := NewMatrixFromSlice([][]float64{{v.X}, {v.Y}, {v.Z}})
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
