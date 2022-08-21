package bmp

/*
RGBTRIPLE

This structure describes a color consisting of relative intensities of
red, green, and blue.

Adapted from http://msdn.microsoft.com/en-us/library/aa922590.aspx.
*/
type RGBTRIPLE struct {
	blue  byte
	green byte
	red   byte
}

/*
BITMAPFILEHEADER

The BITMAPFILEHEADER structure contains information about the type, size,
and layout of a file that contains a DIB [device-independent bitmap].
Adapted from http://msdn.microsoft.com/en-us/library/dd183374(VS.85).aspx.
*/
type BITMAPFILEHEADER struct {
	bfType      uint16
	bfSize      uint32
	bfReserved1 uint16
	bfReserved2 uint16
	bfOffBits   uint16
}

/*
BITMAPINFOHEADER

The BITMAPINFOHEADER structure contains information about the
dimensions and color format of a DIB [device-independent bitmap].

Adapted from http://msdn.microsoft.com/en-us/library/dd183376(VS.85).aspx.
*/
type BITMAPINFOHEADER struct {
	biSize          uint32
	biWidth         int32
	biHeight        int32
	biPlanes        uint16
	biBitCount      uint16
	biCompression   uint32
	biSizeImage     uint32
	biXPelsPerMeter int32
	biYPelsPerMeter int32
	biClrUsed       uint32
	biClrImportant  uint32
}
