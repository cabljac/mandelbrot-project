package render

import (
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/cabljac/coords"
	"github.com/cabljac/mandelbrot"
)

// Generate a Mandelbrot image
func GenerateImage(width, height int) *image.RGBA {
	// 1. Create blank image
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// 2. Create viewport and mandelbrot config
	viewport := coords.NewViewport(width, height) // Use constructor
	cfg := mandelbrot.NewConfig()                 // Use constructor

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			// z := coords.Point{Real: float64(x), Imag: float64(y)}.Complex()
			point := viewport.PixelToComplex(x, y) // This returns coords.Point
			c := point.Complex()                   // Convert to complex128
			escaped, _ := cfg.Calculate(c)         // Calculate with complex128

			var pixelColor color.RGBA
			if !escaped {
				// Point is in the set - black
				pixelColor = color.RGBA{0, 0, 0, 255}
			} else {
				// Point escaped - white (for now)
				pixelColor = color.RGBA{255, 255, 255, 255}
			}

			img.Set(x, y, pixelColor)
		}
	}
	return img

}

func SavePNG(img image.Image, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	return png.Encode(file, img)
}
