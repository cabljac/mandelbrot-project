package coords

type Point struct {
	Real, Imag float64
}

// Convert Point to complex128 for calculations
func (p Point) Complex() complex128 {
	return complex(p.Real, p.Imag)
}

// Viewport defines what part of complex plane we're viewing
type Viewport struct {
	MinReal, MaxReal float64
	MinImag, MaxImag float64
	Width, Height    float64
}

// Constructor with reasonable defaults for Mandelbrot set
func NewViewport(width, height int) *Viewport {

	aspectRatio := float64(height) / float64(width)
	imagRange := 3.5 * aspectRatio // Scale imaginary range by aspect ratio

	return &Viewport{
		MinReal: -2.5,
		MaxReal: 1.0,
		MinImag: -imagRange / 2,
		MaxImag: imagRange / 2,
		Width:   float64(width),
		Height:  float64(height),
	}
}

// Convert pixel coordinates to complex plane
func (v *Viewport) PixelToComplex(x, y int) Point {
	real := v.MinReal + (float64(x)/v.Width)*((v.MaxReal)-(v.MinReal))
	imag := v.MinImag + (float64(y)/v.Height)*((v.MaxImag)-(v.MinImag))

	return Point{
		Real: real,
		Imag: imag,
	}
}
