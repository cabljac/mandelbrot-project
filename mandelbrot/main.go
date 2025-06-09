package mandelbrot

func MandelbrotIteration(z, c complex128) complex128 {
	return z*z + c
}
