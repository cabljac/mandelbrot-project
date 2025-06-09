package main

import (
	"fmt"
	"log"

	"github.com/cabljac/render"
)

func main() {
	fmt.Println("Generating Mandelbrot image...")

	// TODO: at the moment going too big breaks something and results in a corrupted image

	size := 16364

	img := render.GenerateImageWithWorkerPool(size, size)

	err := render.SavePNG(img, "mandelbrot.png")
	if err != nil {
		log.Fatal("Failed to save image:", err)
	}

	fmt.Println("Saved mandelbrot.png!")
}
