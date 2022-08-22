package main

import (
	"os"
)

func main() {
	var result []byte

	data, _ := os.ReadFile("assets/courtyard.bmp")
	bitmapFileHeader := data[:14]
	bitmapInfoHeader := data[14:54]
	bitmap := data[54:]

	result = append(result, bitmapFileHeader...)
	result = append(result, bitmapInfoHeader...)
	result = append(result, bitmap...)

	os.WriteFile("out/result.bmp", result, 0644)
}
