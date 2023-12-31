// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"
)

const (
    width, height = 600, 320            // canvas size in pixels
    cells         = 100                 // number of grid cells
    xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
    xyscale       = width / 2 / xyrange // pixels per x or y unit
    zscale        = height * 0.4        // pixels per z unit
    angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)
var outFilename = flag.String("o", "surface.svg", "output svg file")

func main() {
	flag.Parse()
	f, err := os.Create(*outFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
    fmt.Fprintf(f,"<svg xmlns='http://www.w3.org/2000/svg' "+
        "style='stroke: grey; fill: white; stroke-width: 0.7' "+
        "width='%d' height='%d'>", width, height)
    for i := 0; i < cells; i++ {
        for j := 0; j < cells; j++ {
            ax, ay := corner(i+1, j)
			if !isValid(ax, ay) {
				continue 
			  }
            bx, by := corner(i, j)
			if !isValid(bx, by) {
				continue 
			  }
            cx, cy := corner(i, j+1)
			if !isValid(cx, cy) {
				continue 
			  }
            dx, dy := corner(i+1, j+1)
			if !isValid(dx,dy) {
				continue 
			  }
            fmt.Fprintf(f, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
                ax, ay, bx, by, cx, cy, dx, dy)
        }
    }
    fmt.Fprintf(f, "</svg>")
}

func corner(i, j int) (float64, float64) {
    // Find point (x,y) at corner of cell (i,j).
    x := xyrange * (float64(i)/cells - 0.5)
    y := xyrange * (float64(j)/cells - 0.5)

    // Compute surface height z.
    z := f(x, y)

    // Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
    sx := width/2 + (x-y)*cos30*xyscale
    sy := height/2 + (x+y)*sin30*xyscale - z*zscale
    return sx, sy
}

func f(x, y float64) float64 {
    // r := math.Hypot(x, y) // distance from (0,0)
    // return math.Sin(r) / r

	//egg box
	// return math.Cos(x) * math.Cos(y)

	// moguls
	// return math.Sin(x) * math.Cos(y)

	return math.Atan2(x, y)
}

func isValid(x, y float64) bool {
	return !math.IsNaN(x) && !math.IsInf(x, 0) &&
		   !math.IsNaN(y) && !math.IsInf(y, 0) 
  }