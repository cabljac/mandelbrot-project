package main

import (
	"fmt"
	"log"

	"github.com/cabljac/render"
)

func main() {
	fmt.Println("Generating Mandelbrot image...")

	img := render.GenerateImage(800, 600)

	err := render.SavePNG(img, "mandelbrot.png")
	if err != nil {
		log.Fatal("Failed to save image:", err)
	}

	fmt.Println("Saved mandelbrot.png!")
}
