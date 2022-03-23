package geo

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

type outputImage struct {
	output      *image.RGBA
	width       int
	height      int
	aspectRatio float64
}

func CreateImage(width int, aspectRatio float64) outputImage {
	upLeft := image.Point{0, 0}

	height := float64(width) / aspectRatio

	lowRight := image.Point{int(width), int(height)}

	output := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	return outputImage{
		output:      output,
		width:       width,
		aspectRatio: aspectRatio,
		height:      int(height),
	}
}

func (i *outputImage) Save(path string) {
	f, _ := os.Create(path)

	png.Encode(f, i.output)

	f.Close()
}

func (i *outputImage) SetPixelColor(x, y int, colorValue color.RGBA) {
	i.output.Set(x, y, colorValue)
}
