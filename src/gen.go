package main

import (
	"bitmap"
	"flag"
)

func main(){
	var inFileName, outFileName string
	var width, height int

	flag.StringVar(&inFileName, "input", "", "Input file to use")
	flag.StringVar(&inFileName, "i", "", "Input file to use")

	flag.StringVar(&outFileName, "output", "", "Output file to use")
	flag.StringVar(&outFileName, "o", "", "Output file to use")

	flag.IntVar(&width, "width", 0, "Width of bitmap file generated")
	flag.IntVar(&width, "w", 0, "Width of bitmap file generated")

	flag.IntVar(&height, "height", 0, "Height of bitmap file generated")
	flag.IntVar(&height, "h", 0, "Height of bitmap file generated")

	flag.Parse()

	var bmp bitmap.Bitmap

	bmp.Generate(inFileName, width, height)
	bmp.Write(outFileName)
}

