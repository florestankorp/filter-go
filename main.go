package main

import (
	"os"
)

func main() {
	var result []byte

	data, _ := os.ReadFile("courtyard.bmp")
	BITMAP_FILE_HEADER := data[0:14]
	BITMAP_INFO_HEADER := data[14:54]
	BITMAP := data[54:len(data)]

	result = append(result, BITMAP_FILE_HEADER...)
	result = append(result, BITMAP_INFO_HEADER...)
	result = append(result, BITMAP...)

	os.WriteFile("result.bmp", result, 0644)
}
