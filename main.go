package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"filter-go/pkg/bmp"
	"filter-go/pkg/utils"
	"io"
	"os"
	"unsafe"
)

const (
	sizeOfPixel = 3
	fHeaderSize = 14
	iHeaderSize = 40
)

func main() {

	inFile, err := os.Open("assets/yard.bmp")
	utils.Check(err)

	defer inFile.Close()

	var fHeader bmp.FileHeader
	err = bmp.EncodeHeader(fHeaderSize, inFile, &fHeader)
	utils.Check(err)

	var iHeader bmp.InfoHeader
	err = bmp.EncodeHeader(iHeaderSize, inFile, &iHeader)
	utils.Check(err)

	if fHeader.Type != 0x4d42 ||
		fHeader.OffBits != 54 ||
		iHeader.Size != iHeaderSize ||
		iHeader.BitCount != 24 ||
		iHeader.Compression != 0 {

		panic("Unsupported file format.\n")
	}

	// height is negative
	height := int(-iHeader.Height)
	width := int(iHeader.Width)

	var padding int64 = int64((4 - (width*sizeOfPixel)%4) % 4)

	// allocate memory for pixels
	pixels := make([]bmp.Pixel, width*height)

	// fill pixels with RGB values
	for i := range pixels {
		buffer := make([]byte, sizeOfPixel)
		_, err = inFile.Read(buffer)
		utils.Check(err)

		err = binary.Read(bytes.NewReader(buffer), binary.LittleEndian, &pixels[i])
		utils.Check(err)

		// advance read-header to account for padding
		_, err = inFile.Seek(padding, io.SeekCurrent)
		utils.Check(err)
	}

	image := bmp.Make2DArray(width, height, pixels)

	// bmp.Grayscale(&image)
	bmp.Blur(width, height, &image)

	// STRUCTS TO BYTE CONVERSION
	BMPBytes := make([]byte, 0, fHeader.Size)

	fhSlice := *(*[fHeaderSize]byte)(unsafe.Pointer(&fHeader))
	fiArray := *(*[iHeaderSize]byte)(unsafe.Pointer(&iHeader))

	// remove padding from type in fileHeader
	BMPBytes = append(BMPBytes, fhSlice[:2]...)
	BMPBytes = append(BMPBytes, fhSlice[4:]...)

	// add zeroes back to the end of the fileHeader
	BMPBytes = append(BMPBytes, fhSlice[2:4]...)
	BMPBytes = append(BMPBytes, fiArray[:]...)

	// convert pixels to byte array
	for i := range pixels {
		byteArray := *(*[sizeOfPixel]byte)(unsafe.Pointer(&pixels[i]))
		BMPBytes = append(BMPBytes, byteArray[:]...)
	}

	// add zeroes to the end of the file
	BMPBytes = append(BMPBytes, []byte{0, 0}...)

	// WRITE BYTES TO FILE
	outFile, err := os.Create("./out/result.bmp")
	utils.Check(err)
	defer outFile.Close()

	writer := bufio.NewWriter(outFile)
	_, err = writer.Write(BMPBytes)
	utils.Check(err)

	writer.Flush()
}
