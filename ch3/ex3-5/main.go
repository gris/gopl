// Mandelbrot gera uma imagem PNG do fractal de Mandelbrot.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Ponto (px, py) da imagem representa o valor complexo z
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img) // NOTA: ignorando erros
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 50
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			if 200-contrast*n > 150 {
				return color.RGBA{255, 0, 0, 255}
			}
			if 200-contrast*n > 100 {
				return color.RGBA{0, 255, 0, 255}
			}
			return color.RGBA{0, 0, 255, 255}
		}
	}
	return color.Black
}
