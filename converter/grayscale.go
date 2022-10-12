package converter

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

func TransformImageGrayscale(img image.Image, width, height int) (image.Image, error) {
	grayScale := image.NewGray(image.Rectangle{image.Point{0, 0}, image.Point{width, height}})

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			imageColor := img.At(x, y)
			rr, gg, bb, _ := imageColor.RGBA()
			r := math.Pow(float64(rr), 2.2)
			g := math.Pow(float64(gg), 2.2)
			b := math.Pow(float64(bb), 2.2)
			m := math.Pow(0.2125*r+0.7154*g+0.0721*b, 1/2.2)

			Y := uint16(m + 0.5)

			grayColor := color.Gray{uint8(Y >> 8)}
			grayScale.Set(x, y, grayColor)
		}
	}

	// Encode the grayscale image to the new file
	newFileName := "grayscale.png"
	newfile, err := os.Create(newFileName)

	defer newfile.Close()
	png.Encode(newfile, grayScale)

	return grayScale, err
}
