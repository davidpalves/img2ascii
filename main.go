package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/davidpalves/img2ascii/converter"
)

var result string
var err error

func main() {

	filePath := flag.String("filePath", "", "Path to image file to be converted")
	urlPath := flag.String("urlPath", "", "image URL to be converted")
	width := flag.Int("width", 140, "Width of the output")

	flag.Parse()

	if strings.TrimSpace(*filePath) != "" && strings.TrimSpace(*urlPath) != "" {
		log.Panic("Please insert only one argument")
	}

	if strings.TrimSpace(*filePath) == "" && strings.TrimSpace(*urlPath) == "" {
		log.Panic("Please insert at least one argument")
	}

	imageSize := converter.ImageSize{
		Width:  *width,
		Height: 0,
	}

	if strings.TrimSpace(*urlPath) != "" {
		img := converter.ImageURL{
			UrlPath: *urlPath,
			Image:   imageSize,
		}

		result, err = img.ConvertImage()
	}

	if strings.TrimSpace(*filePath) != "" {
		img := converter.ImageFileSystem{
			FilePath: *filePath,
			Image:    imageSize,
		}

		result, err = img.ConvertImage()
	}

	if err != nil {
		log.Fatal("Could not convert image to ASCII: ", err)
	}

	fmt.Print(result)
}
