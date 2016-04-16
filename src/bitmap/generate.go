package bitmap

import (
	"io"
	"os"
)

const headerSize = 54

func (bmp *Bitmap)Generate(file string, width, height int){
	bmp.Resize(width, height)
	bmp.CoreHead = COREHEADER {
		Ident: 0x4D42,
		Reserved1: 0,
		Reserved2: 0,
		Size: DWORD(headerSize + bmp.ByteSize),
		Offset: headerSize,
	}
	bmp.InfoHead = INFOHEADER{
		Size: headerSize - 14,
		Planes: 1,
		BitCount: 24,
		Compression: 0,
		XPelsPerMeter: 2835,
		YPelsPerMeter: 2835,
		ClrUsed: 0,
		ClrImportant: 0,
	}

	var buf io.Reader
	if file != ""{
		inFile, err := os.Open(file)
		defer inFile.Close()
		checkErr(err)
		buf = io.Reader(inFile)
	} else {
		buf = io.Reader(os.Stdin)
	}
	
	bmp.Data = make([]byte, bmp.ByteSize)
	buf.Read(bmp.Data)
}
