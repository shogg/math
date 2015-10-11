package root

import (
	. "math"
)

const MAX_ITER = 20

var (
	eps float64
	t   float64
)

type F func(x float64) float64

func init() {
	eps = Nextafter(1.0, 2.0) - 1.0
	t = 10e-16
}

func Brent(f F, a, b float64) float64 {

	fa := f(a)
	fb := f(b)
	c := a
	fc := fa
	d := b - a
	e := d

	for i := 0; i < MAX_ITER; i++ {

		if fb*fc > 0.0 {
			c = a
			fc = fa
			d = b - a
			e = d
		}

		if Abs(fc) < Abs(fb) {
			a = b
			fa = fb
			b = c
			fb = fc
			c = a
			fc = fa
		}

		tol := 2.0*eps*Abs(b) + t
		m := (c - b) / 2.0

		if Abs(m) <= tol || Abs(fb) == 0.0 {
			break
		}

		if Abs(e) < tol || Abs(fa) <= Abs(fb) {
			d = m
			e = m

		} else {
			var p, q float64
			s := fb / fa
			if a == c {
				p = 2 * m * s
				q = 1 - s
			} else {
				q = fa / fc
				r := fb / fc
				p = s * (2.0*m*q*(q-r) - (b-a)*(r-1))
				q = (q - 1) * (r - 1) * (s - 1)
			}

			if p > 0 {
				q = -q
			} else {
				p = -p
			}

			s = e
			e = d

			if 2*p < 3*m*q-Abs(tol*q) && p < Abs(s*q/2) {
				d = p / q
			} else {
				d = m
				e = m
			}
		}

		a = b
		fa = fb

		if Abs(d) > tol {
			b += d
		} else if m > 0 {
			b += tol
		} else {
			b -= tol
		}

		fb = f(b)
	}

	return b
}
