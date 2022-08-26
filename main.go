package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"filter-go/pkg/bmp"
	"filter-go/pkg/utils"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"unsafe"
)

const (
	sizeOfPixel = 3
	fHeaderSize = 14
	iHeaderSize = 40
)

func main() {

	log.SetFlags(0) // logger without additional info

	argsWithoutProg := os.Args[1:]

	grayscaleFlagPtr := flag.Bool("g", false, "Make image grayscale")
	blurFlagPtr := flag.Bool("b", false, "Blur image")
	reflectFlagPtr := flag.Bool("r", false, "Reflect image horizontally")

	flag.Parse()

	if flag.NFlag() > 1 {
		e := errors.New("error: only one flag allowed")
		log.Fatal(e)
	}

	if len(argsWithoutProg) != 3 {
		fmt.Println("oops! something went wrong...")
		e := errors.New("usage: ./filter-go [flag] <infile>.bmp <outfile>.bmp")
		log.Fatal(e)
	}

	if !*grayscaleFlagPtr &&
		!*blurFlagPtr &&
		!*reflectFlagPtr &&
		len(argsWithoutProg) == 0 {
		log.SetFlags(0)
		e := errors.New("error: no flag provided")
		log.Fatal(e)
	}

	_, err := os.Stat(argsWithoutProg[1])
	if errors.Is(err, os.ErrNotExist) {
		log.Fatal(err.Error())
	}

	inFile, err := os.Open(argsWithoutProg[1])
	utils.Check(err)

	defer inFile.Close()

	// BYTES TO STRUCT CONVERSION
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

		e := errors.New("error: file format of input is not supported")
		log.Fatal(e)
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

	if *grayscaleFlagPtr {
		bmp.Grayscale(&image)
	}

	if *blurFlagPtr {
		bmp.Blur(width, height, &image)
	}

	if *reflectFlagPtr {
		bmp.Reflect(height, &image)
	}

	// STRUCT TO BYTE CONVERSION
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

	// WRITE TO NEW FILE

	outFolder := "out"

	_, err = os.Stat(outFolder)
	if errors.Is(err, os.ErrNotExist) {
		os.Mkdir(outFolder, os.ModePerm) // make "out" folder if it doesn't exist
	}

	if !strings.HasSuffix(argsWithoutProg[2], ".bmp") {
		e := errors.New("error: output file does not have ending '.bmp'")
		log.Fatal(e)
	}

	outFile, err := os.Create(outFolder + "/" + argsWithoutProg[2])
	utils.Check(err)

	defer outFile.Close()

	writer := bufio.NewWriter(outFile)
	_, err = writer.Write(BMPBytes)
	utils.Check(err)

	writer.Flush()
}
