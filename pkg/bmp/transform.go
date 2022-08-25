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

				avgRed := averageRed(*pixel, *n1, *n2, *n3, *n4, *n5, *n6, *n7, *n8)
				avgGreen := averageGreen(*pixel, *n1, *n2, *n3, *n4, *n5, *n6, *n7, *n8)
				avgBlue := averageBlue(*pixel, *n1, *n2, *n3, *n4, *n5, *n6, *n7, *n8)

				pixel.Red = byte(avgRed)
				n1.Red = byte(avgRed)
				n2.Red = byte(avgRed)
				n3.Red = byte(avgRed)
				n4.Red = byte(avgRed)
				n5.Red = byte(avgRed)
				n6.Red = byte(avgRed)
				n7.Red = byte(avgRed)
				n8.Red = byte(avgRed)

				pixel.Green = byte(avgGreen)
				n1.Green = byte(avgGreen)
				n2.Green = byte(avgGreen)
				n3.Green = byte(avgGreen)
				n4.Green = byte(avgGreen)
				n5.Green = byte(avgGreen)
				n6.Green = byte(avgGreen)
				n7.Green = byte(avgGreen)
				n8.Green = byte(avgGreen)

				pixel.Blue = byte(avgBlue)
				n1.Blue = byte(avgBlue)
				n2.Blue = byte(avgBlue)
				n3.Blue = byte(avgBlue)
				n4.Blue = byte(avgBlue)
				n5.Blue = byte(avgBlue)
				n6.Blue = byte(avgBlue)
				n7.Blue = byte(avgBlue)
				n8.Blue = byte(avgBlue)
			}

		}
	}

}

func averageRed(pixels ...Pixel) int {
	i := 0
	for _, color := range pixels {
		i += int(color.Red)
	}

	return i / (len(pixels))
}
func averageBlue(pixels ...Pixel) int {
	i := 0
	for _, color := range pixels {
		i += int(color.Blue)
	}

	return i / (len(pixels))
}
func averageGreen(pixels ...Pixel) int {
	i := 0
	for _, color := range pixels {
		i += int(color.Green)
	}

	return i / (len(pixels))
}
