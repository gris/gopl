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
			img.Set(px, py, newton(z))
		}
	}
	png.Encode(os.Stdout, img) // NOTA: ignorando erros
}

func newton(z complex128) color.Color {
	const iterations = 200
	const contrast = 15
	var v0 complex128 = z
	var v1 complex128 = z
	for n := uint8(0); n < iterations; n++ {
		v1 = v0 - (cmplx.Pow(v0, 4)-1)/(4*cmplx.Pow(v0, 3))

		if cmplx.Abs(v1-v0) < 0.01 {
			if cmplx.Abs(v1-1) < 0.01 {
				return color.RGBA{255 - contrast*n, 0, 0, 255}
			}

			if cmplx.Abs(v1+1) < 0.01 {
				return color.RGBA{0, 255 - contrast*n, 0, 255}
			}

			if cmplx.Abs(v1+1i) < 0.01 {
				return color.RGBA{0, 0, 255 - contrast*n, 255}
			}

			if cmplx.Abs(v1-1i) < 0.01 {
				return color.RGBA{221 - contrast*n, 160, 221 - contrast*n, 255}
			}

			return color.Gray{255 - contrast*n}
		}
		v0 = v1
	}
	return color.Black
}
