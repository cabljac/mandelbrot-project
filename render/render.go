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
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	viewport := coords.NewViewport(width, height)
	cfg := mandelbrot.NewConfig()

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			point := viewport.PixelToComplex(x, y)
			c := point.Complex()
			escaped, iterations := cfg.Calculate(c)

			var pixelColor color.RGBA
			if !escaped {
				pixelColor = color.RGBA{0, 0, 0, 255}
			} else {
				intensity := uint8((iterations * 255) / cfg.MaxIterations)
				pixelColor = color.RGBA{intensity, intensity, intensity, 255}
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
