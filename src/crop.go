package main

import (
	"flag"
	"bitmap"
)

func main(){
	var inFileName, outFileName string
	var up, down, left, right int

	flag.StringVar(&inFileName, "input", "", "Input file to use")
	flag.StringVar(&inFileName, "i", "", "Input file to use")

	flag.StringVar(&outFileName, "output", "", "Output file to use")
	flag.StringVar(&outFileName, "o", "", "Output file to use")

	flag.IntVar(&up, "up", 0, "Amount to crop from top of picture")
	flag.IntVar(&up, "u", 0, "Amount to crop from top of picture")

	flag.IntVar(&down, "down", 0, "Amount to crop from bottom of picture")
	flag.IntVar(&down, "d", 0, "Amount to crop from bottom of picture")

	flag.IntVar(&left, "left", 0, "Amount to crop from left of picture")
	flag.IntVar(&left, "l", 0, "Amount to crop from left of picture")

	flag.IntVar(&right, "right", 0, "Amount to crop from right of picture")
	flag.IntVar(&right, "r", 0, "Amount to crop from right of picture")

	flag.Parse()

	if up < 0 || down < 0 || left < 0 || right < 0{
		panic("cannot crop negative pixles")
	}

	var bmp bitmap.Bitmap

	bmp.Read(inFileName)
	bmp.Crop(up, down, left, right)
	bmp.Write(outFileName)
}