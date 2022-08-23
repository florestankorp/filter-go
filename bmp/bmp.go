package bmp

type BYTE uint8
type DWORD uint32
type LONG int32
type WORD uint16

/*
RGBTRIPLE

This structure describes a color consisting of relative intensities of
red, green, and blue.

Adapted from http://msdn.microsoft.com/en-us/library/aa922590.aspx.
*/
type RGBTRIPLE struct {
	Blue  byte
	Green byte
	Red   byte
}

/*
BITMAPFILEHEADER

The BITMAPFILEHEADER structure contains information about the type, size,
and layout of a file that contains a DIB [device-independent bitmap].
Adapted from http://msdn.microsoft.com/en-us/library/dd183374(VS.85).aspx.
*/
type BITMAPFILEHEADER struct {
	BfType      WORD
	BfSize      DWORD
	BfReserved1 WORD
	BfReserved2 WORD
	BfOffBits   DWORD
}

/*
BITMAPINFOHEADER

The BITMAPINFOHEADER structure contains information about the
dimensions and color format of a DIB [device-independent bitmap].

Adapted from http://msdn.microsoft.com/en-us/library/dd183376(VS.85).aspx.
*/
type BITMAPINFOHEADER struct {
	BiSize          DWORD
	BiWidth         LONG
	BiHeight        LONG
	BiPlanes        WORD
	BiBitCount      WORD
	BiCompression   DWORD
	BiSizeImage     DWORD
	BiXPelsPerMeter LONG
	BiYPelsPerMeter LONG
	BiClrUsed       DWORD
	BiClrImportant  DWORD
}
