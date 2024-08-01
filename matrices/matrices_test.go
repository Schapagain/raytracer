package matrices

import (
	"testing"

	"github.com/schapagain/raytracer/tuples"
	"github.com/schapagain/raytracer/utils"
)

// TestNewMatrix initializes a new matrix and checks if
// the number of rows and cols have been properly set
// and that the matrix is initialized with zeros
//
// It also implicitly checks if Get(i,j) returns the value at location (i,j)
func TestNewMatrix(t *testing.T) {
	m, _ := NewMatrix(4, 3)

	if m.Cols() != 3 {
		t.Fatalf("Expected matrix to have %d cols, but got %d", 3, m.Cols())
	}
	if m.Rows() != 4 {
		t.Fatalf("Expected matrix to have %d rows, but got %d", 3, m.Rows())
	}

	i := 0
	j := 0
	for {
		if i >= m.Rows() {
			break
		}
		for {
			if j >= m.Cols() {
				break
			}
			mVal, err := m.Get(i, j)
			if err == nil && mVal != 0 {
				t.Fatalf("Expected location (%d,%d) to be initialized to %f, but got %f", i, j, 0.0, mVal)
			}
			if err != nil {
				t.Fatalf("No error expected getting location (%d,%d), but received one: %q", i, j, err)
			}
			j++
		}
		i++
	}

	_, err := m.Get(4, 3)
	if err == nil {
		t.Fatalf("Expected error while accessing location (7,7), but got none")
	}
}

// TestMarixStringRepr checks if the string representation of the matrix is as expected
//
// It also implicitly checks if Set(i,j,val) sets location (i,j) properly with val
func TestMarixStringRepr(t *testing.T) {
	m, _ := NewMatrix(2, 2)
	err := m.Set(1, 1, 3.456)
	if err != nil {
		t.Fatalf("No error expected while setting location (1,1), but received one: %q", err)
	}
	expString := `         0.000          0.000
         0.000          3.456`

	mString := m.String()
	if mString != expString {
		t.Fatalf("Expected matrix string to be:\n%s\nGot:\n%s\n", expString, mString)
	}

	err = m.Set(2, 2, 3.4)
	if err == nil {
		t.Fatalf("Expected error while setting location (2,2), but got none")
	}
}

// TestMatrixIsEqual creates pairs of matrices and
// checks if matrices with equivalent values at corresponding locations
// are deemed equal
func TestMatrixIsEqual(t *testing.T) {

	m1, _ := NewMatrix(2, 2)

	m2, _ := NewMatrix(2, 2)

	m3, _ := NewMatrix(2, 3)

	t.Run("zero matrices", func(t *testing.T) {
		if !m1.IsEqualTo(m2) {
			t.Fatalf("Expected zero matrices of equal dimensions to be equal")
		}
	})
	t.Run("unequal dimensions", func(t *testing.T) {
		if m2.IsEqualTo(m3) {
			t.Fatalf("Not Expected matrices of unequal dimensions to be equal")
		}
	})

	m1.Set(1, 1, 3.43)
	t.Run("differing values", func(t *testing.T) {
		if m1.IsEqualTo(m2) {
			t.Fatalf("Not expected matrices\n%s\nAND\n%s\n with different values to be equal", m1, m2)
		}
	})

	m2.Set(1, 1, 3.430000001)
	t.Run("approx. equal values", func(t *testing.T) {
		if !m1.IsEqualTo(m2) {
			t.Fatalf("Expected matrices\n%s\nAND\n%s\n with approx. equal values to be equal", m1, m2)
		}
	})
}

// TestGetRow checks whether the correct row is returned from the
// matrix
func TestGetRow(t *testing.T) {
	m, _ := NewMatrix(3, 4)
	m.Set(2, 1, 3.22)
	m.Set(1, 3, 3.000)
	expRows := [][]float64{
		{0, 0, 0, 0},
		{0, 0, 0, 2.999999999},
		{0, 3.22, 0, 0},
	}
	rowNum := 0
	for {
		if rowNum > m.Rows() {
			break
		}
		row, err := m.GetRow(rowNum)
		if rowNum == m.Rows() {
			if err == nil {
				t.Fatalf("Error expected when accessing row %d, but received none", rowNum)
			}
		} else {
			if err != nil {
				t.Fatalf("No error expected when accessing row %d, but received one: %q", rowNum, err)
			} else {
				if !utils.FloatSlicesEqual(row, expRows[rowNum]) {
					t.Fatalf("Expected row %d to be %v, but got %v", rowNum, expRows[rowNum], row)
				}
			}
		}
		rowNum++
	}
}

// TestGetCol checks whether the correct row is returned from the
// matrix
func TestGetCol(t *testing.T) {
	m, _ := NewMatrix(4, 3)
	m.Set(3, 1, 3.000)
	m.Set(1, 2, 3.22)
	expCols := [][]float64{
		{0, 0, 0, 0},
		{0, 0, 0, 2.999999999},
		{0, 3.22, 0, 0},
	}
	colNum := 0
	for {
		if colNum > m.Cols() {
			break
		}
		col, err := m.GetCol(colNum)
		if colNum == m.Cols() {
			if err == nil {
				t.Fatalf("Error expected when accessing col %d, but received none", colNum)
			}
		} else {
			if err != nil {
				t.Fatalf("No error expected when accessing row %d, but received one: %q", colNum, err)
			} else {
				if !utils.FloatSlicesEqual(col, expCols[colNum]) {
					t.Fatalf("Expected col %d to be %v, but got %v", colNum, expCols[colNum], col)
				}
			}
		}
		colNum++
	}
}

// TestNewMatrixFromSlice creates new matrices from 2D slices, and
// checks whether a valid matrix representations are created
func TestNewMatrixFromSlice(t *testing.T) {
	testCases := []struct {
		name   string
		matA   [][]float64
		expErr bool
	}{
		{
			"only one element",
			[][]float64{{3.4}},
			false,
		},
		{
			"zero matrix",
			[][]float64{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
			false,
		},
		{
			"identity matrix",
			[][]float64{{4, 3}, {1, -2}, {6.012, 7}, {9, 3.45}},
			false,
		},
		{
			"empty matrix",
			[][]float64{},
			true,
		},
		{
			"empty matrix",
			[][]float64{{}},
			true,
		},
	}

	for _, testCase := range testCases {
		m, err := NewMatrixFromSlice(testCase.matA)
		t.Run(testCase.name, func(t *testing.T) {
			if testCase.expErr && err == nil {
				t.Fatalf("Error expected when building matrix from %v but got none", testCase.matA)
			}
			if !testCase.expErr {
				if err != nil {
					t.Fatalf("No error expected when building a matrix from %v but got one: ", err)
				} else {
					rowNum := 0
					for {
						if rowNum >= len(testCase.matA) {
							break
						}
						colNum := 0
						for {
							if colNum >= len(testCase.matA[rowNum]) {
								break
							}
							matVal, err := m.Get(rowNum, colNum)
							expVal := testCase.matA[rowNum][colNum]
							if err != nil {
								t.Fatalf("No error expected when accessing location (%d,%d) but got one: %q", rowNum, colNum, err)
							} else {
								if expVal != matVal {
									t.Fatalf("Expected location (%d,%d) to have a value of %f, but got %f", rowNum, colNum, expVal, matVal)
								}
							}

							colNum++
						}

						rowNum++
					}
				}
			}
		})
	}

}

// TestMultiply checks whether matrix multiplication is computed correctly
// between compatible matrices, and whether an error is returned for
// incompatible ones
func TestMultiply(t *testing.T) {
	testCases := []struct {
		name    string
		matA    [][]float64
		matB    [][]float64
		expProd [][]float64
		expErr  bool
	}{
		{
			"zero matrices",
			[][]float64{{0, 0, 0}},
			[][]float64{{0}, {0}, {0}},
			[][]float64{{0}},
			false,
		},
		{
			"zero matrices",
			[][]float64{{0}, {0}, {0}},
			[][]float64{{0, 0, 0}},
			[][]float64{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
			false,
		},
		{
			"identity matrix",
			[][]float64{{4, 3}, {1, -2}, {6.012, 7}, {9, 3.45}},
			[][]float64{{1, 0}, {0, 1}},
			[][]float64{{4, 3}, {1, -2}, {6.012, 7}, {9, 3.45}},
			false,
		}, {
			"invalid dimensions",
			[][]float64{{4, 3}, {1, -2}, {6.012, 7}, {9, 3.45}},
			[][]float64{{1, 0}, {0, 0}, {1, 1}},
			[][]float64{{}},
			true,
		},
	}

	for _, testCase := range testCases {
		matA, _ := NewMatrixFromSlice(testCase.matA)
		matB, _ := NewMatrixFromSlice(testCase.matB)
		expProd, _ := NewMatrixFromSlice(testCase.expProd)
		r1 := matA.Rows()
		c1 := matA.Cols()
		r2 := matB.Rows()
		c2 := matB.Cols()
		t.Run(testCase.name, func(t *testing.T) {
			prod, err := matA.Multiply(matB)
			if testCase.expErr && err == nil {
				t.Fatalf("Error expected when multiplying matrices with dimensions %dX%d and %dX%d, but received none", r1, c1, r2, c2)
			}
			if !testCase.expErr {
				if err != nil {
					t.Fatalf("No error expected when multiplying matrices with dimensions %dX%d and %dX%d, but received one: %q", r1, c1, r2, c2, err)
				} else {
					if !prod.IsEqualTo(expProd) {
						t.Fatalf("Expected product to be:\n%s\nbut, got:\n%s\n", expProd, prod)
					}
				}
			}
		})
	}
}

// TestNewIdentityMatrix creates new identity matrices and
// checks if compatible matrices and vectors multiplied with them
// remain unchanged
func Test3DIdentityMatrix(t *testing.T) {
	iden3d, err := NewIdentityMatrix(3)
	if err != nil {
		t.Fatalf("No error expected when creating a 3x3 identity matrix")
	} else {
		colV := NewMatrixFromVector(tuples.Vector{X: -2, Y: 1.03, Z: 300})
		mat, _ := NewMatrixFromSlice([][]float64{{-4, 1.03, 90}, {0, 0, -3}, {0.94, 0, 234}})
		prod, _ := iden3d.Multiply(colV)
		if !colV.IsEqualTo(prod) {
			t.Fatalf("Expected\n%v\nto remain unchanged after multiplication with identity vector, but got:\n%v", colV, prod)
		}
		prod, _ = iden3d.Multiply(mat)
		if !mat.IsEqualTo(prod) {
			t.Fatalf("Expected\n%v\nto remain unchanged after multiplication with identity vector, but got:\n%v", mat, prod)
		}
	}
}

// TestTranspose checks if transposes of matrices are created correctly
func TestTranspose(t *testing.T) {
	testCases := []struct {
		name     string
		matA     [][]float64
		expTrans [][]float64
	}{
		{
			"row zero to column",
			[][]float64{{0, 0, 0}},
			[][]float64{{0}, {0}, {0}},
		},
		{
			"column zero to row",
			[][]float64{{0}, {0}, {0}},
			[][]float64{{0, 0, 0}},
		},
		{
			"identity matrix",
			[][]float64{{1, 0}, {0, 1}},
			[][]float64{{1, 0}, {0, 1}},
		}, {
			"column block to row",
			[][]float64{{4, 3}, {1, -2}, {6.012, 7}, {9, 3.45}},
			[][]float64{{4, 1, 6.012, 9}, {3, -2, 7, 3.45}},
		},
	}
	for _, testCase := range testCases {
		matA, _ := NewMatrixFromSlice(testCase.matA)
		expTrans, _ := NewMatrixFromSlice(testCase.expTrans)
		t.Run(testCase.name, func(t *testing.T) {
			trans := matA.Transposed()
			if !trans.IsEqualTo(expTrans) {
				t.Fatalf("Expected transpose of\n%s\nto be:\n%s\nbut, got:\n%s\n", matA, expTrans, trans)
			}
		})
	}
}

// TestSubMatrix checks if submatrices are computed correctly
func TestSubMatrix(t *testing.T) {
	testCases := []struct {
		name   string
		matA   [][]float64
		args   []int
		expSub [][]float64
		expErr bool
	}{
		{
			"zero matrix",
			[][]float64{{0, 0, 0}, {0, 0, 0}},
			[]int{0, 2},
			[][]float64{{0, 0}},
			false,
		},
		{
			"remove first row and col",
			[][]float64{{4, 3}, {1, -2}, {6.012, 7}, {9, 3.45}},
			[]int{0, 0},
			[][]float64{{-2}, {7}, {3.45}},
			false,
		},
		{
			"remove last row and col",
			[][]float64{{4, 3, 0}, {1, -2, 12}, {0.99, 6.012, 7}},
			[]int{2, 2},
			[][]float64{{4, 3}, {1, -2}},
			false,
		},
		{
			"invalid submatrix",
			[][]float64{{4, 3}, {1, -2}, {6.012, 7}, {9, 3.45}},
			[]int{5, 1},
			[][]float64{{-2}, {7}, {3.45}},
			true,
		},
	}
	for _, testCase := range testCases {
		matA, _ := NewMatrixFromSlice(testCase.matA)
		expSub, _ := NewMatrixFromSlice(testCase.expSub)
		subMat, err := matA.SubMatrix(testCase.args[0], testCase.args[1])
		if testCase.expErr {
			if err == nil {
				t.Fatalf("Expected error while extracting submatrix(%d,%d) of\n%s\n, but received none", testCase.args[0], testCase.args[1], matA)
			}
		} else {
			if err != nil {
				t.Fatalf("Expected no error while extracting submatrix(%d,%d) of\n%s\n, but received one: %q", testCase.args[0], testCase.args[1], matA, err)
			} else {
				t.Run(testCase.name, func(t *testing.T) {
					if !subMat.IsEqualTo(expSub) {
						t.Fatalf("Expected submatrix(%d,%d) of\n%s\nto be:\n%s\nbut, got:\n%s\n", testCase.args[0], testCase.args[1], matA, expSub, subMat)
					}
				})
			}
		}
	}
}
