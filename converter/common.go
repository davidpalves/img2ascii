package converter

import (
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

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
	var img image.Image
	var err error

	f, err := os.Open(filePath)

	if err != nil {
		log.Printf("failed opening %s: %s", filePath, err)
		panic(err.Error())
	}

	defer f.Close()
	fileExt := filepath.Ext(f.Name())

	if fileExt == ".jpg" {
		img, err = jpeg.Decode(f)
	} else if fileExt == ".png" {
		img, _, err = image.Decode(f)
	}

	return img, err
}

func getImageFromURL(urlPath string) (image.Image, error) {
	temp_directory, err := os.MkdirTemp("", "temp_")
	if err != nil {
		log.Println("Failed to create temporary directory: %w", err)
	}

	fileUrl, err := url.Parse(urlPath)
	path := fileUrl.Path
	segments := strings.Split(path, "/")
	fileName := segments[len(segments)-1]

	resp, err := http.Get(urlPath)
	if err != nil {
		log.Println("Failed to download image: %w", err)
	}
	defer resp.Body.Close()

	filePath := fmt.Sprintf("%s/%s", temp_directory, fileName)

	tempImg, _ := os.Create(filePath)
	defer tempImg.Close()

	io.Copy(tempImg, resp.Body)

	return getImageFromFile(filePath)
}
