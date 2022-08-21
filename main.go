package main

import (
	"bytes"
	"os"
)

func main() {
	var result []byte

	data, _ := os.ReadFile("assets/courtyard.bmp")
	dataReader := bytes.NewReader(data)

	BITMAP_FILE_HEADER := make([]byte, 14)
	BITMAP_INFO_HEADER := make([]byte, 40)
	BITMAP := make([]byte, len(data))

	dataReader.Read(BITMAP_FILE_HEADER)
	dataReader.Read(BITMAP_INFO_HEADER)
	dataReader.Read(BITMAP)

	result = append(result, BITMAP_FILE_HEADER...)
	result = append(result, BITMAP_INFO_HEADER...)
	result = append(result, BITMAP...)

	os.WriteFile("out/result.bmp", result, 0644)
}
