package bmp

/*
Pixel

This structure describes a color consisting of relative intensities of
red, green, and blue.

Adapted from http://msdn.microsoft.com/en-us/library/aa922590.aspx.
*/
type Pixel struct {
	Blue  byte
	Green byte
	Red   byte
}

/*
FileHeader

The FileHeader structure contains information about the type, size,
and layout of a file that contains a DIB [device-independent bitmap].
Adapted from http://msdn.microsoft.com/en-us/library/dd183374(VS.85).aspx.
*/
type FileHeader struct {
	Type      uint16
	Size      uint32
	Reserved1 uint16
	Reserved2 uint16
	OffBits   uint16
}

/*
InfoHeader

The InfoHeader structure contains information about the
dimensions and color format of a DIB [device-independent bitmap].

Adapted from http://msdn.microsoft.com/en-us/library/dd183376(VS.85).aspx.
*/
type InfoHeader struct {
	Size          uint32
	Width         int32
	Height        int32
	Planes        uint16
	BitCount      uint16
	Compression   uint32
	SizeImage     uint32
	XPelsPerMeter int32
	YPelsPerMeter int32
	ClrUsed       uint32
	ClrImportant  uint32
}
