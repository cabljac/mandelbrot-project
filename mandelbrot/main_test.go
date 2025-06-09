package mandelbrot

import (
	"math/cmplx"
	"math/rand"
	"testing"
)

func randomComplexNumber(t *testing.T) complex128 {
	t.Helper()
	real := rand.Float64()*4 - 2
	imag := rand.Float64()*4 - 2

	return complex(real, imag)
}

func TestMandelbrotIterationSimple(t *testing.T) {
	result := MandelbrotIteration(1i, 1)
	expected := complex128(0)
	tolerance := 1e-10

	if cmplx.Abs(result-expected) > tolerance {
		t.Errorf("Abs(3+4i) = %f; want %f", result, expected)
	}
}

func TestMandelbrotIterationRandom(t *testing.T) {

	for _ = range 1000 {

		z, c := randomComplexNumber(t), randomComplexNumber(t)

		result := MandelbrotIteration(z, c)
		expected := z*z + c
		tolerance := 1e-10

		if cmplx.Abs(result-expected) > tolerance {
			t.Errorf("Abs(3+4i) = %f; want %f", result, expected)
		}
	}

}
