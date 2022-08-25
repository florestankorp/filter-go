package bmp

type Color int64

const (
	Red Color = iota
	Green
	Blue
)

func Grayscale(image *[][]Pixel) {
	for i, scanLine := range *image {
		for j, pixel := range scanLine {
			avg := (int(pixel.Red) + int(pixel.Green) + int(pixel.Blue)) / 3

			(*image)[i][j].Red = byte(avg)
			(*image)[i][j].Green = byte(avg)
			(*image)[i][j].Blue = byte(avg)
		}
	}

}

// Reflect image horizontally
func Reflect(height int, image *[][]Pixel) {
	buffer := make([][]Pixel, 0, height)

	for _, scanLine := range *image {
		buffer = append(buffer, reverse(scanLine))
	}

	*image = buffer
}

func Blur(width int, height int, image *[][]Pixel) {
	for i, scanLine := range *image {
		for j := range scanLine {
			isFirstLine := i == 0
			isLastLine := i == height-1
			isFirstCol := j == 0
			isLastCol := j == width-1

			pixel := &(*image)[i][j]

			if !isFirstLine &&
				!isLastLine &&
				!isFirstCol &&
				!isLastCol {

				// Neighbors
				n1 := &(*image)[i-1][j-1]
				n2 := &(*image)[i-1][j]
				n3 := &(*image)[i-1][j+1]
				n4 := &(*image)[i][j-1]
				n5 := &(*image)[i][j+1]
				n6 := &(*image)[i+1][j-1]
				n7 := &(*image)[i+1][j]
				n8 := &(*image)[i+1][j+1]

				avg := average(*pixel, *n1, *n2, *n3, *n4, *n5, *n6, *n7, *n8)

				(*pixel) = avg
				(*n1) = avg
				(*n2) = avg
				(*n3) = avg
				(*n4) = avg
				(*n5) = avg
				(*n6) = avg
				(*n7) = avg
				(*n8) = avg

			}
		}
	}
}

func average(pixels ...Pixel) Pixel {
	var red, blue, green int

	for _, p := range pixels {
		red += int(p.Red)
		blue += int(p.Blue)
		green += int(p.Green)
	}

	return Pixel{
		Red:   byte(red / len(pixels)),
		Blue:  byte(blue / len(pixels)),
		Green: byte(green / len(pixels)),
	}
}
