package main

import (
	"bytes"
	"encoding/binary"
	"filter-go/bmp"
	"fmt"
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

	height := bitmapInfoHeader.Height
	width := bitmapInfoHeader.Width

	fmt.Println(width, height) // 600, -400

}
