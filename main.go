package main

import (
	"bytes"
	"encoding/binary"
	"filter-go/bmp"
	"fmt"
	"os"
)

func main() {
	data, _ := os.ReadFile("assets/courtyard.bmp")
	buffer := bytes.NewReader(data)

	var bitmapFileHeader bmp.BITMAPFILEHEADER
	binary.Read(buffer, binary.LittleEndian, &bitmapFileHeader)

	var bitmapInfoHeader bmp.BITMAPINFOHEADER
	binary.Read(buffer, binary.LittleEndian, &bitmapInfoHeader)

	fmt.Println(bitmapInfoHeader.BiWidth)  // should be 600, is 39321600
	fmt.Println(bitmapInfoHeader.BiHeight) // should be -400, is -26214400
}
