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

type ImageSize struct {
	Width  int
	Height int
}

type ImageURL struct {
	Image   ImageSize
	UrlPath string
}

type ImageFileSystem struct {
	Image    ImageSize
	FilePath string
}

func (i *ImageSize) scaleImage(img image.Image, width int) image.Image {
	sz := img.Bounds()
	height := (sz.Max.Y * width * 10) / (sz.Max.X * 16)
	img = resize.Resize(uint(width), uint(height), img, resize.Lanczos3)

	i.Height = height
	i.Width = width

	return img
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

func (i ImageURL) ConvertImage() (string, error) {
	rawImg, err := getImageFromURL(i.UrlPath)
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}

	result := transformToAscii(rawImg, i.Image)

	return string(result), nil
}

func (i ImageFileSystem) ConvertImage() (string, error) {
	rawImg, err := getImageFromFile(i.FilePath)

	if err != nil {
		log.Println("%w", err)
		return "", err
	}

	result := transformToAscii(rawImg, i.Image)

	return string(result), nil
}
