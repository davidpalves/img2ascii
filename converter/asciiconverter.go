package converter

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"log"
	"os"
	"reflect"
)

const ASCII_CHARS string = " $@B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/\\|()1{}[]?-_+~i!lI;:,\"^`.   "

func transformToAscii(rawImg image.Image, imgSize ImageSize) []byte {
	table := []byte(ASCII_CHARS)
	buffer := new(bytes.Buffer)

	img := imgSize.scaleImage(rawImg, imgSize.Width)

	for i := 0; i < imgSize.Height; i++ {
		for j := 0; j < imgSize.Width; j++ {
			gray := color.GrayModel.Convert(img.At(j, i))
			y := reflect.ValueOf(gray).FieldByName("Y").Uint()
			position := int(y * 69 / 255)
			_ = buffer.WriteByte(table[position])
		}
		_ = buffer.WriteByte('\n')
	}

	return buffer.Bytes()
}

func (i ImageURL) ConvertImageToASCII() ([]byte, error) {
	rawImg, err := getImageFromURL(i.UrlPath)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	result := transformToAscii(rawImg, i.Image)

	return result, nil
}

func (i ImageFileSystem) ConvertImageToASCII() ([]byte, error) {
	rawImg, err := getImageFromFile(i.FilePath)

	if err != nil {
		log.Println("%w", err)
		return nil, err
	}

	result := transformToAscii(rawImg, i.Image)

	return result, nil
}

func OutputASCIIToFile(img []byte, output string) {
	f, err := os.Create(output)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	f.Write(img)
}
