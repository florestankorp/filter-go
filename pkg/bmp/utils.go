package bmp

import (
	"bytes"
	"encoding/binary"
	"filter-go/pkg/utils"
	"fmt"
	"os"
)

func EncodeHeader(bufferSize int, file *os.File, header interface{}) error {
	buffer := make([]byte, bufferSize)
	_, err := file.Read(buffer)
	utils.Check(err)

	if error := binary.Read(bytes.NewReader(buffer), binary.LittleEndian, header); error != nil {
		return fmt.Errorf("failed to parse DIB header: %w", error)
	}

	return nil
}
func reverse(pixels []Pixel) []Pixel {
	for i, j := 0, len(pixels)-1; i < j; i, j = i+1, j-1 {
		pixels[i], pixels[j] = pixels[j], pixels[i]
	}

	return pixels
}

func Make2DArray(width int, height int, pixels []Pixel) [][]Pixel {
	buffer := make([][]Pixel, 0, width*height)

	for i := 0; i < height; i++ {
		start := i * width
		end := width + (i * width)
		slice := pixels[start:end]
		buffer = append(buffer, slice)
	}

	return buffer
}
