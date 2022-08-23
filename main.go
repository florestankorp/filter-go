package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"filter-go/bmp"
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

	inFile, error := os.Open("assets/courtyard.bmp")
	check(error)
	defer inFile.Close()

	var bitmapFileHeader bmp.BitmapFileHeader
	var bitmapInfoHeader bmp.BitmapInfoHeader

	bmp.DecodeHeader(14, inFile, &bitmapFileHeader)
	bmp.DecodeHeader(40, inFile, &bitmapInfoHeader)

	if bitmapFileHeader.Type != 0x4d42 ||
		bitmapFileHeader.OffBits != 54 ||
		bitmapInfoHeader.Size != 40 ||
		bitmapInfoHeader.BitCount != 24 ||
		bitmapInfoHeader.Compression != 0 {

		panic("Unsupported file format.\n")
	}

	height := int(-bitmapInfoHeader.Height) // height is negative
	width := int(bitmapInfoHeader.Width)

	var padding int64 = int64((4 - ((width)*3)%4) % 4)

	pixels := make([]bmp.RGBTriple, width*height) // allocate memory for pixels

	// fill pixels with RGB values
	for i := range pixels {
		buffer := make([]byte, 3)
		inFile.Read(buffer)
		binary.Read(bytes.NewReader(buffer), binary.LittleEndian, &pixels[i])
		inFile.Seek(padding, io.SeekCurrent)
	}

	// STRUCTS TO BYTE CONVERSION
	bytes := make([]byte, 0, width*height*3)
	bitmapFileHeaderByteArray := *(*[14]byte)(unsafe.Pointer(&bitmapFileHeader))
	bitmapInfoHeaderByteArray := *(*[40]byte)(unsafe.Pointer(&bitmapInfoHeader))

	// parse out extra zeroes...
	// why are there extra zeroes in the fileHeader!?
	bytes = append(bytes, bitmapFileHeaderByteArray[:2]...)
	bytes = append(bytes, bitmapFileHeaderByteArray[4:]...)
	// need to add zeroes back to the end of the fileHeader :(
	bytes = append(bytes, []byte{0, 0}[:]...)
	bytes = append(bytes, bitmapInfoHeaderByteArray[:]...)

	for i := range pixels {
		byteArray := *(*[3]byte)(unsafe.Pointer(&pixels[i]))
		bytes = append(bytes, byteArray[:]...)
	}

	outFile, err := os.Create("./out/result.bmp")
	check(err)
	defer outFile.Close()

	writer := bufio.NewWriter(outFile)
	_, err = writer.Write(bytes)
	check(err)

	writer.Flush()
}
