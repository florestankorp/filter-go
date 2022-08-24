package bmp

import (
	"bytes"
	"encoding/binary"
	"filter-go/pkg/utils"
	"fmt"
	"os"
)

/*
Pixel

This structure describes a color consisting of relative intensities of
red, green, and blue.

Adapted from http://msdn.microsoft.com/en-us/library/aa922590.aspx.
*/
type Pixel struct {
	Blue  byte
	Green byte
	Red   byte
}

/*
FileHeader

The FileHeader structure contains information about the type, size,
and layout of a file that contains a DIB [device-independent bitmap].
Adapted from http://msdn.microsoft.com/en-us/library/dd183374(VS.85).aspx.
*/
type FileHeader struct {
	Type      uint16
	Size      uint32
	Reserved1 uint16
	Reserved2 uint16
	OffBits   uint16
}

/*
InfoHeader

The InfoHeader structure contains information about the
dimensions and color format of a DIB [device-independent bitmap].

Adapted from http://msdn.microsoft.com/en-us/library/dd183376(VS.85).aspx.
*/
type InfoHeader struct {
	Size          uint32
	Width         int32
	Height        int32
	Planes        uint16
	BitCount      uint16
	Compression   uint32
	SizeImage     uint32
	XPelsPerMeter int32
	YPelsPerMeter int32
	ClrUsed       uint32
	ClrImportant  uint32
}

func EncodeHeader(bufferSize int, file *os.File, header interface{}) error {
	buffer := make([]byte, bufferSize)
	_, err := file.Read(buffer)
	utils.Check(err)

	if error := binary.Read(bytes.NewReader(buffer), binary.LittleEndian, header); error != nil {
		return fmt.Errorf("failed to parse DIB header: %w", error)
	}

	return nil
}

func Grayscale(pixels *[]Pixel) {
	for i := range *pixels {
		avg := (int((*pixels)[i].Red) + int((*pixels)[i].Green) + int((*pixels)[i].Blue)) / 3

		(*pixels)[i].Red = byte(avg)
		(*pixels)[i].Green = byte(avg)
		(*pixels)[i].Blue = byte(avg)
	}

}

// Reflect image horizontally
func Reflect(width int, height int, pixels *[]Pixel) {
	buffer := make([]Pixel, 0, width*height)

	for i := 0; i < height; i++ {
		start := i * width
		end := width + (i * width)
		slice := (*pixels)[start:end]
		buffer = append(buffer, reverse(slice)...)
	}

	*pixels = buffer
}

// // Blur image
// func blur(height int, width  int, pixel image[height][width]){
//     return;
// }

// // Detect edges
// func edges(height int, width  int, pixel image[height][width]){
//     return;
// }

func reverse(pixels []Pixel) []Pixel {
	for i, j := 0, len(pixels)-1; i < j; i, j = i+1, j-1 {
		pixels[i], pixels[j] = pixels[j], pixels[i]
	}

	return pixels
}
