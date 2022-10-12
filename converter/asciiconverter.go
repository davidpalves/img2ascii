package converter

import (
	"bytes"
	"image"
	"image/color"
	"reflect"
)

const ASCII_CHARS string = " $@B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/\\|()1{}[]?-_+~i!lI;:,\"^`.   "

func transformToAscii(img image.Image, width, height int) []byte {
	table := []byte(ASCII_CHARS)
	buffer := new(bytes.Buffer)

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			gray := color.GrayModel.Convert(img.At(j, i))
			y := reflect.ValueOf(gray).FieldByName("Y").Uint()
			position := int(y * 69 / 255)
			_ = buffer.WriteByte(table[position])
		}
		_ = buffer.WriteByte('\n')
	}

	return buffer.Bytes()
}
