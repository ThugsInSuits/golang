// Lissajous generates GIF animations of random Lissajous figures.
package main
import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)
// set color is green by color.RGBA 
// set color is white by color.White
// set color is black by color.Black
// set color is blue by color.RGBA{0x00, 0x00, 0xFF, 0xFF}
// set color is red by color.RGBA{0xFF, 0x00, 0x00, 0xFF}
// set color is yellow by color.RGBA{0xFF, 0xFF, 0x00, 0xFF}
// set color is cyan by color.RGBA{0x00, 0xFF, 0xFF, 0xFF}
// set color is purple by color.RGBA{0x80, 0x00, 0x80, 0xFF}
// set color is orange by color.RGBA{0xFF, 0xA5, 0x00, 0xFF}
// set color is brown by color.RGBA{0x80, 0x00, 0x00, 0xFF}


var palette = []color.Color{color.White, color.RGBA{0xFF, 0xFF, 0x00, 0xFF}} 
const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)
func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	lissajous(os.Stdout) 
}
func lissajous(out io.Writer) {
	const (
		cycles = 5 // number of complete x oscillator revolutions
		res = 0.001 // angular resolution
		size = 100 // image canvas covers [-size..+size]
		nframes = 64 // number of animation frames
		delay = 8 // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+1), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}