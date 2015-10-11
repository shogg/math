package linalg

import (
	. "github.com/gonum/matrix/mat64"
)

// Solve for x: mat^T * mat * x = mat^T * vec
func LeastSquares(mat Matrix, vec *Vector) *Vector {

	mat3 := NewDense(3, 3, nil)
	vec3 := NewVector(3, nil)
	x := NewVector(3, nil)

	mat3.Mul(mat.T(), mat)
	vec3.MulVec(mat.T(), vec)
	x.SolveVec(mat3, vec3)

	return x
}
