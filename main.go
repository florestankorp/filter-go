package main

import (
	"filter-go/bmp"
	"fmt"
	"os"
	"unsafe"
)

func check(e error) {
	if e != nil {
		panic(e.Error())
	}
}

func main() {

	file, error := os.Open("assets/courtyard.bmp")
	check(error)
	defer file.Close()

	var bitmapFileHeader bmp.BitmapFileHeader
	var bitmapInfoHeader bmp.BitmapInfoHeader

	bmp.DecodeHeader(14, file, &bitmapFileHeader)
	bmp.DecodeHeader(40, file, &bitmapInfoHeader)

	if bitmapFileHeader.Type != 0x4d42 ||
		bitmapFileHeader.OffBits != 54 ||
		bitmapInfoHeader.Size != 40 ||
		bitmapInfoHeader.BitCount != 24 ||
		bitmapInfoHeader.Compression != 0 {

		panic("Unsupported file format.\n")
	}

	height := -bitmapInfoHeader.Height // height is negative
	width := bitmapInfoHeader.Width

	var rgbTriple bmp.RGBTriple
	padding := (4 - (int(width)*int(unsafe.Sizeof(rgbTriple)))%4) % 4

	fmt.Println(height)
	fmt.Println(padding)

}
