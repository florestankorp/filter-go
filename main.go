package main

import (
	"bufio"
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

	// height := -bitmapInfoHeader.Height // height is negative
	// width := bitmapInfoHeader.Width

	// var rgbTriple bmp.RGBTriple
	// var padding int64 = (4 - (int64(width)*int64(unsafe.Sizeof(rgbTriple)))%4) % 4

	// allocate memory for pixels
	// pixels := make([]bmp.RGBTriple, (int(height) * (int(width) * int(unsafe.Sizeof(rgbTriple)))))
	r4 := bufio.NewReader(file)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Println(b4)
	// for i, _ := range pixels {
	// 	buffer := make([]byte, int(unsafe.Sizeof(rgbTriple)))
	// 	file.Read(buffer)
	// 	binary.Read(bytes.NewReader(buffer), binary.LittleEndian, &pixels[i])

	// }

	// iterate over scanlines
	// for i := 0; i < int(height); i++ {
	// 	buffer := make([]byte, int(width)*int(unsafe.Sizeof(rgbTriple)))
	// 	file.Read(buffer)
	// 	binary.Read(bytes.NewReader(buffer), binary.LittleEndian, &pixels[i])

	// 	file.Seek(padding, 0)
	// }

	// fmt.Println(pixels[0])

}
