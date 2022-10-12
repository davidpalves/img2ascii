package converter

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
)

func (i ImageURL) DesaturateImage(output string) (string, error) {
	rawImg, err := getImageFromURL(i.UrlPath)
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}

	result, err := transformImageGrayscale(rawImg, output)
	if err != nil {
		log.Println("%w", err)
		return "", err
	}

	return string(result), nil
}

func (i ImageFileSystem) DesaturateImage(output string) (string, error) {
	rawImg, err := getImageFromFile(i.FilePath)

	if err != nil {
		log.Println("%w", err)
		return "", err
	}

	result, err := transformImageGrayscale(rawImg, output)
	if err != nil {
		log.Println("%w", err)
		return "", err
	}

	return string(result), nil
}

func transformImageGrayscale(img image.Image, output string) (string, error) {
	height := img.Bounds().Max.Y
	width := img.Bounds().Max.X
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
	newFileName := output
	newfile, err := os.Create(newFileName)

	defer newfile.Close()
	png.Encode(newfile, grayScale)

	return newFileName, err
}
