package main

import (
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"
)

var (
	iCol    = flag.Int("column", -1, "Column to copy")
	iColS   = flag.Int("columnStart", 0, "Column to copy start")
	iColE   = flag.Int("columnEnd", 0, "Column to copy end")
	iLine   = flag.Int("line", -1, "Line to copy")
	iLineS  = flag.Int("lineStart", 0, "Line to copy start")
	iLineE  = flag.Int("lineEnd", 0, "Line to copy end")
	sImgSrc = flag.String("imgSrc", "", "Image file")
	sImgDst = flag.String("imgDst", "", "Image file for output")
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Simple image editor, that duplicates specified line or column of JPEG image to selected image's location and saves it into new file\n")
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	if *sImgSrc == "" {
		fmt.Printf("You must specify image file in JPEG format\n")
		flag.Usage()
		return
	}

	reader, err := os.Open(*sImgSrc)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()
	m, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	bounds := m.Bounds()
	fmt.Printf("Image geometry: %v\n", bounds)

	// Copy image to new RGBA canvas
	m1 := image.NewRGBA(m.Bounds())
	for x := 0; x < bounds.Max.X; x++ {
		for y := 0; y < bounds.Max.Y; y++ {
			m1.Set(x, y, m.At(x, y))
		}
	}

	fmt.Printf("iLine=%v iLineS=%v iLineE=%v iCol=%v iColS=%v iColE=%v\n", *iLine, *iLineS, *iLineE, *iCol, *iColS, *iColE)

	// copying lines
	if *iLine != -1 {
		for x := 0; x < bounds.Max.X; x++ {
			for l := *iLineS; l <= *iLineE; l++ {
				m1.Set(x, l, m.At(x, *iLine))
			}
		}
	}

	// copying columns
	if *iCol != -1 {
		for y := 0; y < bounds.Max.Y; y++ {
			for c := *iColS; c <= *iColE; c++ {
				m1.Set(c, y, m.At(*iCol, y))
			}
		}
	}

	// save image
	sImgDst1 := *sImgDst
	if sImgDst1 == "" {
		sImgDst1 = "copy_" + *sImgSrc
	}
	fmt.Printf("Output file=%v\n", sImgDst1)
	w, err := os.Create(sImgDst1)
	if err != nil {
		log.Fatal(err)
	}
	defer w.Close()
	var opt jpeg.Options
	opt.Quality = 100
	jpeg.Encode(w, m1, &opt)
}
