package main

import (
	"bytes"
	"encoding/binary"
	"filter-go/bmp"
	"os"
	"unsafe"
)

func check(e error) {
	if e != nil {
		panic(e.Error())
	}
}

func main() {

	data, error := os.ReadFile("assets/courtyard.bmp")
	check(error)

	buffer := bytes.NewReader(data)

	var bitmapFileHeader bmp.BitmapFileHeader
	binary.Read(buffer, binary.LittleEndian, &bitmapFileHeader)

	var bitmapInfoHeader bmp.BitmapInfoHeader
	binary.Read(buffer, binary.LittleEndian, &bitmapInfoHeader)

	if bitmapFileHeader.Type != 0x4d42 ||
		bitmapFileHeader.OffBits != 54 ||
		bitmapInfoHeader.Size != 40 ||
		bitmapInfoHeader.BitCount != 24 ||
		bitmapInfoHeader.Compression != 0 {

		panic("Unsupported file format.\n")

	}

	// height is a negative number
	height := -bitmapInfoHeader.Height
	width := bitmapInfoHeader.Width

	var rgbTriple bmp.RGBTriple
	padding := (4 - (int(width)*int(unsafe.Sizeof(rgbTriple)))%4) % 4

}
