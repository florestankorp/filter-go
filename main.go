package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"filter-go/bmp"
	"fmt"
	"io"
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

	reader := bufio.NewReader(file)

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

	height := int(-bitmapInfoHeader.Height) // height is negative
	width := int(bitmapInfoHeader.Width)

	var padding int64 = int64((4 - ((width)*int(unsafe.Sizeof(bmp.RGBTriple{})))%4) % 4)

	// allocate memory for pixels
	pixels := make([]bmp.RGBTriple, width*height)

	for i := range pixels {
		buffer := make([]byte, 3)
		file.Read(buffer)
		binary.Read(bytes.NewReader(buffer), binary.LittleEndian, &pixels[i])
		file.Seek(padding, io.SeekCurrent)
	}

	fmt.Println(reader.Peek(2))
	fmt.Println(pixels[len(pixels)-1])
	fmt.Println(bitmapFileHeader.Size)

}
