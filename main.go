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

const (
	sizeOfPixel      = 3
	sizeOfFileHeader = 14
	sizeOfInfoHeader = 40
)

func check(e error) {
	if e != nil {
		panic(e.Error())
	}
}

func main() {

	inFile, error := os.Open("assets/yard.bmp")
	check(error)
	defer inFile.Close()

	var bitmapFileHeader bmp.BitmapFileHeader
	error = bmp.EncodeHeader(14, inFile, &bitmapFileHeader)
	check(error)

	var bitmapInfoHeader bmp.BitmapInfoHeader
	error = bmp.EncodeHeader(sizeOfInfoHeader, inFile, &bitmapInfoHeader)
	check(error)

	if bitmapFileHeader.Type != 0x4d42 ||
		bitmapFileHeader.OffBits != 54 ||
		bitmapInfoHeader.Size != sizeOfInfoHeader ||
		bitmapInfoHeader.BitCount != 24 ||
		bitmapInfoHeader.Compression != 0 {

		panic("Unsupported file format.\n")
	}

	height := int(-bitmapInfoHeader.Height) // height is negative
	width := int(bitmapInfoHeader.Width)

	var padding int64 = int64((4 - (width*sizeOfPixel)%4) % 4)

	pixels := make([]bmp.Pixel, width*height) // allocate memory for pixels

	// fill pixels with RGB values
	for i := range pixels {
		buffer := make([]byte, sizeOfPixel)
		inFile.Read(buffer)
		binary.Read(bytes.NewReader(buffer), binary.LittleEndian, &pixels[i])
		inFile.Seek(padding, io.SeekCurrent)
	}

	// STRUCTS TO BYTE CONVERSION
	BMPBytes := make([]byte, 0, bitmapFileHeader.Size) // is it better to predefine the capacity or the length?

	// isn't there a better way to convert structs to byte arrays?
	bitmapFileHeaderByteArray := *(*[sizeOfFileHeader]byte)(unsafe.Pointer(&bitmapFileHeader))
	bitmapInfoHeaderByteArray := *(*[sizeOfInfoHeader]byte)(unsafe.Pointer(&bitmapInfoHeader))

	// remove padding from type in fileHeader
	BMPBytes = append(BMPBytes, bitmapFileHeaderByteArray[:2]...)
	BMPBytes = append(BMPBytes, bitmapFileHeaderByteArray[4:]...)

	// add zeroes back to the end of the fileHeader
	BMPBytes = append(BMPBytes, bitmapFileHeaderByteArray[2:4]...)
	BMPBytes = append(BMPBytes, bitmapInfoHeaderByteArray[:]...)

	// convert pixels to bytes
	for i := range pixels {
		byteArray := *(*[sizeOfPixel]byte)(unsafe.Pointer(&pixels[i]))
		BMPBytes = append(BMPBytes, byteArray[:]...)
	}

	// add zeroes to the end of the file
	BMPBytes = append(BMPBytes, []byte{0, 0}...)

	// WRITE BYTES TO FILE
	outFile, err := os.Create("./out/result.bmp")
	check(err)
	defer outFile.Close()

	writer := bufio.NewWriter(outFile)
	_, err = writer.Write(BMPBytes)
	check(err)

	writer.Flush()
}
