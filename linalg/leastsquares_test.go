package linalg_test

import (
	. "github.com/go-gl/mathgl/mgl64"
	"github.com/gonum/matrix/mat64"
	"github.com/shogg/math/linalg"
	"testing"
)

func TestLeastSquares(t *testing.T) {

	matA := mat64.NewDense(5, 3, []float64{
		1, -2, 4,
		1, -1, 1,
		1, 0, 0,
		1, 1, 1,
		1, 2, 4,
	})

	vecb := mat64.NewVector(5, []float64{
		0,
		0,
		1,
		0,
		0,
	})

	x := vec3(linalg.LeastSquares(matA, vecb))

	expected := Vec3{34.0 / 70.0, 0.0, -10.0 / 70.0}
	if x != expected {
		t.Errorf("expected %v, got %v", expected, x)
	}
}

func vec3(vec *mat64.Vector) Vec3 {
	return Vec3{vec.At(0, 0), vec.At(1, 0), vec.At(2, 0)}
}
