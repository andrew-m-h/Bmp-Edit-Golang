package main

import (
	"bitmap"
	"flag"
)

func main(){
	var inFileName, outFileName string
	var threshold float64

	flag.StringVar(&inFileName, "input", "", "Input file to use")
	flag.StringVar(&inFileName, "i", "", "Input file to use")

	flag.StringVar(&outFileName, "output", "", "Output file to use")
	flag.StringVar(&outFileName, "o", "", "Output file to use")

	flag.Float64Var(&threshold, "threshold", 0.5, "Threshold to use")
	flag.Float64Var(&threshold, "t", 0.5, "Threshold to use")

	flag.Parse()

	var bmp bitmap.Bitmap

	bmp.Read(inFileName)
	bmp.Threshold(float32(threshold))
	bmp.Write(outFileName)
}
