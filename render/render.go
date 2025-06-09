package render

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"runtime"
	"sync"

	"github.com/cabljac/coords"
	"github.com/cabljac/mandelbrot"
)

func worker(id int, jobs <-chan int, img *image.RGBA, viewport *coords.Viewport, cfg *mandelbrot.Config, width int) {
	for rowNum := range jobs {

		for x := range width {
			point := viewport.PixelToComplex(x, rowNum)
			c := point.Complex()

			escaped, iterations := cfg.Calculate(c)

			var pixelColor color.RGBA
			if !escaped {
				pixelColor = color.RGBA{0, 0, 0, 255}
			} else {
				intensity := uint8((iterations * 255) / cfg.MaxIterations)
				pixelColor = color.RGBA{intensity, intensity, intensity, 255}
			}

			img.Set(x, rowNum, pixelColor)
		}
	}
}

func GenerateImageWithWorkerPoolCustom(width, height, numWorkers int) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	viewport := coords.NewViewport(width, height)
	cfg := mandelbrot.NewConfig()

	jobs := make(chan int, height)
	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerId int) {
			defer wg.Done()
			worker(workerId, jobs, img, viewport, cfg, width)
		}(i)
	}

	for y := 0; y < height; y++ {
		jobs <- y
	}
	close(jobs)

	wg.Wait()
	return img
}

func GenerateImageWithWorkerPool(width, height int) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	viewport := coords.NewViewport(width, height)
	cfg := mandelbrot.NewConfig()

	// channel to hold row numbers
	jobs := make(chan int, height)

	numWorkers := runtime.NumCPU()

	var wg sync.WaitGroup

	for i := range numWorkers {
		wg.Add(1)
		// goroutine to send work to channel
		go func(workerId int) {
			defer wg.Done()

			worker(workerId, jobs, img, viewport, cfg, width)

		}(i)
	}
	for y := range height {
		jobs <- y
	}
	// NO MORE JOBS!
	close(jobs)

	wg.Wait()
	return img
}

func GenerateImageInParallelByRow(width, height int) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	viewport := coords.NewViewport(width, height)
	cfg := mandelbrot.NewConfig()

	var wg sync.WaitGroup

	for y := range height {
		wg.Add(1)

		go func(row int) {
			defer wg.Done()

			for x := range width {
				point := viewport.PixelToComplex(x, row)

				c := point.Complex()

				escaped, iterations := cfg.Calculate(c)

				var pixelColor color.RGBA
				if !escaped {
					pixelColor = color.RGBA{0, 0, 0, 255}
				} else {
					intensity := uint8((iterations * 255) / cfg.MaxIterations)
					pixelColor = color.RGBA{intensity, intensity, intensity, 255}
				}
				img.Set(x, row, pixelColor)

			}
		}(y)
	}

	wg.Wait()

	return img
}

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
