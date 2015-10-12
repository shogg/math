package contour

import (
	. "github.com/go-gl/mathgl/mgl64"
	"github.com/gonum/matrix/mat64"
	"github.com/shogg/math/linalg"
	"github.com/shogg/math/root"
)

type F func(Vec3) float64
type DF func(Vec3) Vec3

var (
	base = []Vec3{
		{1.0, 0.0, 0.0},
		{0.0, 1.0, 0.0},
		{0.0, 0.0, 1.0}}
	cube = []Vec3{
		{0.0, 0.0, 0.0}, // 0
		{1.0, 0.0, 0.0}, // 1
		{0.0, 1.0, 0.0}, // 2
		{1.0, 1.0, 0.0}, // 3
		{0.0, 0.0, 1.0}, // 4
		{1.0, 0.0, 1.0}, // 5
		{0.0, 1.0, 1.0}, // 6
		{1.0, 1.0, 1.0}} // 7
	edges = [][]int{
		{0, 1}, {0, 2}, {1, 3}, {2, 3},
		{0, 4}, {1, 5}, {2, 6}, {3, 7},
		{4, 5}, {4, 6}, {5, 7}, {6, 7}}
)

func Dual(f F, df DF, size int) (vertices []Vec3, faces [][3]int) {

	hdata := make([][2]Vec3, 0, 8)
	a := make([]float64, 0, 8*3)
	b := make([]float64, 0, 8)
	vindex := make(map[Vec3]int)

	it := cartesian(size, size, size)
	for x, y, z, next := it(); next; x, y, z, next = it() {
		o := Vec3{float64(x), float64(y), float64(z)}

		var signs [8]bool
		for i, v := range cube {
			signs[i] = f(o.Add(v)) > 0
		}

		hdata = hdata[0:0]
		for _, e := range edges {
			if signs[e[0]] == signs[e[1]] {
				continue
			}

			h := hermite(f, df, o.Add(cube[e[0]]), o.Add(cube[e[1]]))
			hdata = append(hdata, h)
		}

		a = a[0:0]
		for _, h := range hdata {
			a = append(a, h[1][:]...)
		}

		b = b[0:0]
		for _, h := range hdata {
			b = append(b, h[0].Dot(h[1]))
		}

		mata := mat64.NewDense(len(b), 3, a)
		vecb := mat64.NewVector(len(b), b)
		v := vec3(linalg.LeastSquares(mata, vecb))

		if v.Sub(o).Len() > 1.5 {
			continue
		}

		vindex[o] = len(vertices)
		vertices = append(vertices, v)
	}

	it = cartesian(size, size, size)
	for x, y, z, next := it(); next; x, y, z, next = it() {
		o := Vec3{float64(x), float64(y), float64(z)}

		var ok bool
		var o0 int

		if o0, ok = vindex[o]; !ok {
			continue
		}

		for i := 0; i < 3; i++ {
			for j := 0; j < i; j++ {

				var oi, oj, oij int

				if oi, ok = vindex[o.Add(base[i])]; !ok {
					continue
				}
				if oj, ok = vindex[o.Add(base[j])]; !ok {
					continue
				}
				if oij, ok = vindex[o.Add(base[i].Add(base[j]))]; !ok {
					continue
				}

				faces = append(faces, [3]int{o0, oi, oj}, [3]int{oij, oj, oi})
			}
		}
	}

	return vertices, faces
}

func vec3(vec *mat64.Vector) Vec3 {
	return Vec3{vec.At(0, 0), vec.At(1, 0), vec.At(2, 0)}
}

func hermite(f F, df DF, v0, v1 Vec3) [2]Vec3 {

	lambda := func(t float64) float64 {
		return f(v0.Mul(1.0 - t).Add(v1.Mul(t)))
	}

	t0 := root.Brent(lambda, 0.0, 1.0)
	x0 := v0.Mul(1.0 - t0).Add(v1.Mul(t0))

	return [2]Vec3{x0, df(x0)}
}

func cartesian(s0, s1, s2 int) func() (int, int, int, bool) {

	i0, i1, i2, next := 0, 0, 0, true

	return func() (int, int, int, bool) {

		if i0 < s0 {
			i0++
		} else if i1 < s1 {
			i0 = 0
			i1++
		} else if i2 < s2 {
			i1 = 0
			i2++
		} else {
			next = false
		}

		return i0, i1, i2, next
	}
}
