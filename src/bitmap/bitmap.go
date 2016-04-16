package bitmap

import (
	"os"
	"encoding/binary"
	"io"
)

type (
	DWORD int32
	WORD int16
	LONG int32
)

type COREHEADER struct {
	Ident WORD
	Size DWORD
	Reserved1,
	Reserved2 WORD
	Offset DWORD
}

type INFOHEADER struct {
	Size DWORD
	Width,
	Height LONG
	Planes,
	BitCount WORD
	Compression,
	SizeImage DWORD
	XPelsPerMeter,
	YPelsPerMeter LONG
	ClrUsed,
	ClrImportant DWORD
}

type Bitmap struct {
	CoreHead COREHEADER
	InfoHead INFOHEADER
	Data []byte
	//Extra information stored to reduce recalculation
	ByteWidth,
	ByteSize int
}

func (bmp *Bitmap) Resize(width, height int) {
	bmp.ByteWidth = int((float32(width) * float32(bmp.InfoHead.BitCount) + 31.0) / 32.0) * 4
	bmp.ByteSize = bmp.ByteWidth * height

	var newdata []byte = make([]byte, bmp.ByteSize)
	copy(newdata, bmp.Data)
	bmp.Data = newdata

	bmp.CoreHead.Size = DWORD(bmp.ByteSize) + 54
	bmp.InfoHead.Width = LONG(width)
	bmp.InfoHead.Height = LONG(height)
	bmp.InfoHead.SizeImage = DWORD(bmp.ByteSize)
}

func checkErr(err error){
	if err != nil {
		panic(err)
	}
}

func (bmp *Bitmap) Read(file string){
	var buf io.Reader
	if file != ""{
		inFile, err := os.Open(file)
		defer inFile.Close()
		checkErr(err)
		buf = io.Reader(inFile)
	} else {
		buf = io.Reader(os.Stdin)
	}

	//Core Header
	checkErr(binary.Read(buf, binary.LittleEndian, &bmp.CoreHead.Ident))
	checkErr(binary.Read(buf, binary.LittleEndian, &bmp.CoreHead.Size))
	checkErr(binary.Read(buf, binary.LittleEndian, &bmp.CoreHead.Reserved1))
	checkErr(binary.Read(buf, binary.LittleEndian, &bmp.CoreHead.Reserved2))
	checkErr(binary.Read(buf, binary.LittleEndian, &bmp.CoreHead.Offset))

	//Info Header
	checkErr(binary.Read(buf, binary.LittleEndian, &bmp.InfoHead.Size))
	checkErr(binary.Read(buf, binary.LittleEndian, &bmp.InfoHead.Width))
	checkErr(binary.Read(buf, binary.LittleEndian, &bmp.InfoHead.Height))
	checkErr(binary.Read(buf, binary.LittleEndian, &bmp.InfoHead.Planes))
	checkErr(binary.Read(buf, binary.LittleEndian, &bmp.InfoHead.BitCount))
	checkErr(binary.Read(buf, binary.LittleEndian, &bmp.InfoHead.Compression))
	checkErr(binary.Read(buf, binary.LittleEndian, &bmp.InfoHead.SizeImage))
	checkErr(binary.Read(buf, binary.LittleEndian, &bmp.InfoHead.XPelsPerMeter))
	checkErr(binary.Read(buf, binary.LittleEndian, &bmp.InfoHead.YPelsPerMeter))
	checkErr(binary.Read(buf, binary.LittleEndian, &bmp.InfoHead.ClrUsed))
	checkErr(binary.Read(buf, binary.LittleEndian, &bmp.InfoHead.ClrImportant))

	bmp.ByteWidth = int((float32(bmp.InfoHead.Width) * float32(bmp.InfoHead.BitCount) + 31.0) / 32.0) * 4
	bmp.ByteSize = bmp.ByteWidth * int(bmp.InfoHead.Height)

	//Read Data
	bmp.Data = make([]byte, bmp.ByteSize)
	buf.Read(bmp.Data)

}

func (bmp *Bitmap) Write(file string){
	var buf io.Writer

	if file != "" {
		outFile, err := os.Create(file)
		checkErr(err)
		defer outFile.Close()
		buf = io.Writer(outFile)
	} else {
		buf = io.Writer(os.Stdout)
	}

	checkErr(binary.Write(buf, binary.LittleEndian, bmp.CoreHead.Ident))
	checkErr(binary.Write(buf, binary.LittleEndian, bmp.CoreHead.Size))
	checkErr(binary.Write(buf, binary.LittleEndian, bmp.CoreHead.Reserved1))
	checkErr(binary.Write(buf, binary.LittleEndian, bmp.CoreHead.Reserved2))
	checkErr(binary.Write(buf, binary.LittleEndian, bmp.CoreHead.Offset))

	//Info Header
	checkErr(binary.Write(buf, binary.LittleEndian, bmp.InfoHead.Size))
	checkErr(binary.Write(buf, binary.LittleEndian, bmp.InfoHead.Width))
	checkErr(binary.Write(buf, binary.LittleEndian, bmp.InfoHead.Height))
	checkErr(binary.Write(buf, binary.LittleEndian, bmp.InfoHead.Planes))
	checkErr(binary.Write(buf, binary.LittleEndian, bmp.InfoHead.BitCount))
	checkErr(binary.Write(buf, binary.LittleEndian, bmp.InfoHead.Compression))
	checkErr(binary.Write(buf, binary.LittleEndian, bmp.InfoHead.SizeImage))
	checkErr(binary.Write(buf, binary.LittleEndian, bmp.InfoHead.XPelsPerMeter))
	checkErr(binary.Write(buf, binary.LittleEndian, bmp.InfoHead.YPelsPerMeter))
	checkErr(binary.Write(buf, binary.LittleEndian, bmp.InfoHead.ClrUsed))
	checkErr(binary.Write(buf, binary.LittleEndian, bmp.InfoHead.ClrImportant))

	buf.Write(bmp.Data)

}
