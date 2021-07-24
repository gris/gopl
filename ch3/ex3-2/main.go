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

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // seno(30°), cosseno(30°)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; strokewidth: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, ok1 := corner(i+1, j)
			bx, by, ok2 := corner(i, j)
			cx, cy, ok3 := corner(i, j+1)
			dx, dy, ok4 := corner(i+1, j+1)
			if !ok1 || !ok2 || !ok3 || !ok4 {
				continue
			}
			fmt.Printf("<polygon points='%g,%g,%g,%g,%g,%g,%g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, bool) {
	// Encontra o ponto (x,y) no canto da célula (i,j)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Calcula a altura z da superfície
	z, ok := saddle(x, y)
	if !ok {
		return 0, 0, false
	}
	// Faz uma projeção isométrica de (x,y,z) sobre (sx,sy) do canvas SVG 2D
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale

	return sx, sy, true
}

func eggbox(x, y float64) (value float64, ok bool) {

	return 0.2 * (math.Cos(x) + math.Cos(y)), true
}

func saddle(x, y float64) (value float64, ok bool) {

	return 0.002 * (x*x - y*y), true
}
