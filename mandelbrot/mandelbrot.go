package mandelbrot

import "math/cmplx"

type Config struct {
	MaxIterations int
	EscapeRadius  float64
}

func NewConfig() *Config {
	return &Config{
		MaxIterations: 100,
		EscapeRadius:  2.0,
	}
}

func (cfg *Config) Calculate(c complex128) (bool, int) {
	return iterateUntilEscape(0, c, cfg.MaxIterations)
}

func mandelbrotIteration(z, c complex128) complex128 {
	return z*z + c
}

func hasEscaped(z complex128) bool {
	return cmplx.Abs(z) > 2
}

func iterateUntilEscape(z, c complex128, maxIterations int) (bool, int) {
	if hasEscaped(z) {
		return true, 0
	}

	for i := 1; i <= maxIterations; i++ { // Start from 1
		z = mandelbrotIteration(z, c)
		if hasEscaped(z) {
			return true, i // Now i correctly represents iterations performed
		}
	}
	return false, maxIterations
}
