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

	file, error := os.Open("assets/courtyard.bmp")
	check(error)
	defer file.Close()

	var bitmapFileHeader bmp.BitmapFileHeader
	buffer := make([]byte, 14)
	file.Read(buffer)
	binary.Read(bytes.NewReader(buffer), binary.LittleEndian, &bitmapFileHeader)

	var bitmapInfoHeader bmp.BitmapInfoHeader
	buffer1 := make([]byte, 40)
	file.Read(buffer1)
	binary.Read(bytes.NewReader(buffer1), binary.LittleEndian, &bitmapInfoHeader)

	fmt.Println(bitmapFileHeader)
	fmt.Println(bitmapInfoHeader)

}
