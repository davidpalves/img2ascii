package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/davidpalves/img2ascii/converter"
)

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

	if strings.TrimSpace(*urlPath) != "" {
		ascii, err := converter.ConvertImageFromURL(*urlPath, *width)
		if err != nil {
			log.Panic("Could not convert image from URL: " + *urlPath)
		}
		fmt.Print(ascii)
	}

	if strings.TrimSpace(*filePath) != "" {
		ascii, err := converter.ConvertImageFromFilePath(*filePath, *width)
		if err != nil {
			log.Panic("Could not convert image from URL: " + *filePath)
		}
		fmt.Print(ascii)
	}
}
