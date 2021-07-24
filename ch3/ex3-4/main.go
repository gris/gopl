// Surface calcula uma superfície de renderização SVG de uma função de superfície 3D.
package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
)

const (
	defaultWidth, defaultHeight = 600, 320                   // tamanho do canvas em pixels
	cells                       = 100                        // número de células da grade
	xyrange                     = 30.0                       // intervalos dos eixos (-xyrange..+xyrange)
	xyscale                     = defaultWidth / 2 / xyrange // pixels por unidade x ou y
	zscale                      = defaultHeight * 0.4        // pixels por unidade z
	angle                       = math.Pi / 6                // ângulo dos eixos x, y (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // seno(30°), cosseno(30°)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")
		height, err := strconv.Atoi(r.URL.Query().Get("height"))
		if err != nil {
			height = defaultHeight
		}
		width, err := strconv.Atoi(r.URL.Query().Get("width"))
		if err != nil {
			width = defaultWidth
		}

		color := r.URL.Query().Get("color")

		if color == "" {
			color = "grey"
		}

		surface(w, uint(height), uint(width), color)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func surface(w io.Writer, height uint, width uint, color string) {

	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; strokewidth: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j, height, width)
			bx, by := corner(i, j, height, width)
			cx, cy := corner(i, j+1, height, width)
			dx, dy := corner(i+1, j+1, height, width)
			fmt.Fprintf(w, "<polygon points='%g,%g,%g,%g,%g,%g,%g,%g' style='stroke: %s; fill: white; strokewidth: 0.7'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, color)
		}
	}
	fmt.Fprintln(w, "</svg>")

}

func corner(i, j int, height uint, width uint) (float64, float64) {
	// Encontra o ponto (x,y) no canto da célula (i,j)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Calcula a altura z da superfície
	z := f(x, y)

	// Faz uma projeção isométrica de (x,y,z) sobre (sx,sy) do canvas SVG 2D
	sx := float64(width)/2 + (x-y)*cos30*xyscale
	sy := float64(height)/2 + (x+y)*sin30*xyscale - z*zscale

	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distância de (0,0)
	return math.Sin(r) / r
}
