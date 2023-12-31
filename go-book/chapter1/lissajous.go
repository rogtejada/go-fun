//generate a lissajous gif into stdout

package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var palette = []color.Color{color.Black, color.RGBA{178, 222, 39, 1}, color.White, color.RGBA{20, 110, 150, 1}, color.RGBA{200, 10, 100, 1}, color.RGBA{30, 180, 180, 1}, color.RGBA{38, 180, 120, 1}}

const (
	//indexes from pallete
	backgroundIndex = 0
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     //number of X oscillator revolution
		res     = 0.001 //angular resolution
		size    = 100   //image canvas cover [-size..+size]
		nframes = 64    //number of anim frames
		delay   = 8     //delay between frames in 10ms units

	)

	freq := rand.Float64() * 3.0 //relative frequency of y osc
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 //phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(rand.Intn(5-1)+1))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
