package converter

import (
	"fmt"
	"image"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/nfnt/resize"
)

func scaleImage(img image.Image, width int) (image.Image, int, int) {
	sz := img.Bounds()
	height := (sz.Max.Y * width * 10) / (sz.Max.X * 16)
	img = resize.Resize(uint(width), uint(height), img, resize.Lanczos3)

	return img, width, height
}

func getImageFromFile(filePath string) (image.Image, error) {
	f, err := os.Open(filePath)

	if err != nil {
		log.Printf("failed opening %s: %s", filePath, err)
		panic(err.Error())
	}

	defer f.Close()
	image, _, err := image.Decode(f)

	return image, err
}

func getImageFromURL(urlPath string) (image.Image, error) {
	temp_directory, err := os.MkdirTemp("", "temp_")
	if err != nil {
		log.Println("Failed to create temporary directory: %w", err)
	}

	filePath := fmt.Sprintf("%s/tempImg.png", temp_directory)
	tempImg, _ := os.Create(filePath)
	defer tempImg.Close()

	resp, err := http.Get(urlPath)
	if err != nil {
		log.Println("Failed to download image: %w", err)
	}
	defer resp.Body.Close()

	io.Copy(tempImg, resp.Body)

	return getImageFromFile(filePath)
}

func ConvertImageFromURL(urlPath string, width int) (string, error) {
	rawImg, err := getImageFromURL(urlPath)
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}

	img, WIDTH, HEIGHT := scaleImage(rawImg, width)

	result := transformToAscii(img, WIDTH, HEIGHT)

	return string(result), nil
}

func ConvertImageFromFilePath(filePath string, width int) (string, error) {
	rawImg, err := getImageFromFile(filePath)

	if err != nil {
		log.Println("%w", err)
		return "", err
	}

	img, width, height := scaleImage(rawImg, width)

	result := transformToAscii(img, width, height)

	return string(result), nil
}
