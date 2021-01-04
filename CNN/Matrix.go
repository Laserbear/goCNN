package CNN

import (
	"fmt"
	"math/rand"
)

type Matrix struct {
	values     []float64
	dimensions [2]uint
	length     uint
}

func (m1 Matrix) multiplyMatrix(m2 Matrix) Matrix {
	if m1.dimensions[1] != m2.dimensions[0] {
		panic(fmt.Sprintf("It's all a goddamn farce, matrix with dimension %v cannot be multiplied with matrix of dimension %v\n", m1.dimensions, m2.dimensions))
	}
	productMatrixLength := uint(m1.dimensions[0] * m2.dimensions[1])
	productMatrixValues := make([]float64, productMatrixLength)
	m2Transpose := m2.transpose()
	fmt.Printf("Tranpose values: %v\n", m2Transpose.values)

	for i := uint(0); i < m1.dimensions[0]; i++ {
		for j := uint(0); j < m2.dimensions[1]; j++ {
			vector_length := m1.dimensions[1] //length of vector to dotproduct
			productMatrixValues[i*m2.dimensions[1]+j] = dotproduct(
				m1.values[i*vector_length:i*vector_length+vector_length], m2Transpose.values[j*vector_length:j*vector_length+vector_length])
		}
	}
	return Matrix{values: []float64(productMatrixValues), dimensions: [2]uint{m1.dimensions[0], m2.dimensions[1]}, length: productMatrixLength}
}

func dotproduct(A []float64, B []float64) (dotproduct float64) {
	sum := float64(0)
	for i := range A {
		sum += A[i] * B[i]
	}
	return sum
}

func (m Matrix) transpose() Matrix {
	tranposedMatrixValues := make([]float64, m.length)
	for i := uint(0); i < m.dimensions[0]; i++ {
		for j := uint(0); j < m.dimensions[1]; j++ {
			tranposedMatrixValues[j*m.dimensions[0]+i] = m.values[i*m.dimensions[1]+j]
		}
	}
	return Matrix{values: tranposedMatrixValues, dimensions: [2]uint{m.dimensions[1], m.dimensions[0]}, length: m.length}
}

func (m1 Matrix) add(m2 Matrix) Matrix {
	if m1.dimensions[0] != m2.dimensions[0] || m1.dimensions[1] != m2.dimensions[1] {
		panic(fmt.Sprintf("Why god why? you can't add a matrix of dimension %v with a matrix of dimension %v!\n", m1.dimensions, m2.dimensions))
	}
	summedMatrixValues := make([]float64, m1.length)
	for i := uint(0); i < m1.length; i++ {
		summedMatrixValues[i] = m1.values[i] + m2.values[i]
	}
	return Matrix{values: summedMatrixValues, dimensions: m1.dimensions, length: m1.length}
}

func (m Matrix) multiplyFloat(s float64) Matrix {
	productMatrixValues := make([]float64, m.length)
	for i := uint(0); i < m.length; i++ {
		productMatrixValues[i] = m.values[i] * s
	}
	return Matrix{values: productMatrixValues, dimensions: m.dimensions, length: m.length}
}

func randomFloatMatrix(m uint, n uint) Matrix {
	values := make([]float64, m*n)
	for i := uint(0); i < m*n; i++ {
		values[i] = rand.Float64()
	}
	return Matrix{values: values, dimensions: [2]uint{m, n}, length: m * n}
}

/**
func main () {
	testMatrix := Matrix{values: []float64{1, 0, 0, 0, 1, 0, 0, 0, 1}, dimensions: [2]uint{3, 3}, length: 9}
	testMatrix2 := Matrix{values: []float64{1, 0, 0, 0, 1, 0, 0, 0, 1}, dimensions: [2]uint{3, 3}, length: 9}
	fmt.Printf("Matrix product %v \n", testMatrix.multiplyMatrix(testMatrix2))
	fmt.Printf("dotproduct %v \n", dotproduct([]float64{1, 1, 1}, []float64{1, 1 ,1}))
} **/
