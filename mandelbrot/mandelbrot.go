package mandelbrot

import "math/cmplx"

func mandelbrotIteration(z, c complex128) complex128 {
	return z*z + c
}

func hasEscaped(z complex128) bool {
	return cmplx.Abs(z) > 2
}

func iterateUntilEscape(z, c complex128, maxIterations int) (complex128, int) {
	if hasEscaped(z) {
		return z, 0
	}
	i := 0
	for i < maxIterations {
		z = mandelbrotIteration(z, c)
		if hasEscaped(z) {
			return z, i
		}
		i++
	}
	return z, maxIterations // Need to return something if we never escape
}
