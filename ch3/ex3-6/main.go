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
	// Provavelmente essa função tá errada, visto o resultado final da figura
	const iterations = 200
	const contrast = 15
	var v1 complex128
	var v2 complex128
	var v3 complex128
	var v4 complex128
	z1 := complex(real(z), imag(z))
	z2 := complex(real(z), imag(z)+imag(1))
	z3 := complex(real(z)+1, imag(z))
	z4 := complex(real(z)+1, imag(z)+imag(1))
	for n := uint8(0); n < iterations; n++ {
		v1 = v1*v1 + z1
		v2 = v2*v2 + z2
		v3 = v3*v3 + z3
		v4 = v4*v4 + z4
		if (cmplx.Abs(v1)+cmplx.Abs(v2)+cmplx.Abs(v3)+cmplx.Abs(v4))/4 > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
