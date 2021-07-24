// Surface calcula uma superfície de renderização SVG de uma função de superfície 3D.
package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // tamanho do canvas em pixels
	cells         = 100                 // número de células da grade
	xyrange       = 30.0                // intervalos dos eixos (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels por unidade x ou y
	zscale        = height * 0.4        // pixels por unidade z
	angle         = math.Pi / 6         // ângulo dos eixos x, y (=30°)
)

type CornerType int

const (
	middle CornerType = 0 // not the peak or valley of surface
	peak   CornerType = 1 // the peak of surface corner
	valley CornerType = 2 // the valley of surface corner
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // seno(30°), cosseno(30°)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; strokewidth: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, ct1 := corner(i+1, j)
			bx, by, ct2 := corner(i, j)
			cx, cy, ct3 := corner(i, j+1)
			dx, dy, ct4 := corner(i+1, j+1)

			var color string
			if ct1 == peak || ct2 == peak || ct3 == peak || ct4 == peak {
				color = "#f00"
			} else if ct1 == valley || ct2 == valley || ct3 == valley || ct4 == valley {
				color = "#00f"
			} else {
				// same as default
				color = "grey"
			}
			fmt.Printf("<polygon points='%g,%g,%g,%g,%g,%g,%g,%g' style='stroke: %s'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, color)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, CornerType) {
	// Encontra o ponto (x,y) no canto da célula (i,j)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Calcula a altura z da superfície
	z, colorType := eggbox(x, y)

	// Faz uma projeção isométrica de (x,y,z) sobre (sx,sy) do canvas SVG 2D
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale

	return sx, sy, colorType
}

func eggbox(x, y float64) (float64, CornerType) {
	colorType := middle
	if (math.Sin(x) + math.Sin(y)) == 0 {
		if math.Cos(x)+math.Cos(y) < 0 {
			colorType = valley
		} else if math.Cos(x)+math.Cos(y) > 0 {
			colorType = peak
		}
	}
	return 0.2 * (math.Cos(x) + math.Cos(y)), colorType
}
