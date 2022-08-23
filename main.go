package main

import (
	"bytes"
	"encoding/binary"
	"filter-go/bmp"
	"os"
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

	var bitmapFileHeader bmp.BITMAPFILEHEADER
	binary.Read(buffer, binary.LittleEndian, &bitmapFileHeader)

	var bitmapInfoHeader bmp.BITMAPINFOHEADER
	binary.Read(buffer, binary.LittleEndian, &bitmapInfoHeader)

	if bitmapFileHeader.BfType != 0x4d42 ||
		bitmapFileHeader.BfOffBits != 54 ||
		bitmapInfoHeader.BiSize != 40 ||
		bitmapInfoHeader.BiBitCount != 24 ||
		bitmapInfoHeader.BiCompression != 0 {

		panic("Unsupported file format.\n")

	}

}
