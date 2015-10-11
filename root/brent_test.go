package root_test

import (
	"github.com/shogg/math/root"
	"math"
	"testing"
)

func TestBrent(t *testing.T) {

	// f(x) = e^-x * log(x)
	f := func(x float64) float64 {
		return math.Pow(math.E, -x) * math.Log(x)
	}

	x0 := root.Brent(f, 0.05, 1.7)

	if x0 != 1.0 {
		t.Errorf("Root at 1.0 expected, was: %f", x0)
	}
}
