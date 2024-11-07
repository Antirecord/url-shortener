package main

import (
	"flag"
	"fmt"
)

func main() {
	imgFile := flag.String("file", "", "input image file")
	destDir := flag.String("dest", "./output", "destination folder")
	width := flag.Int("w", 1024, "width of the image")
	isThumb := flag.Bool("thumb", false, "create thumb")

	// разбор командной строки
	flag.Parse()
	fmt.Println("Image file:", *imgFile)
	fmt.Println("Destination folder:", *destDir)
	fmt.Println("Width:", *width)
	fmt.Println("Thumbs:", *isThumb)
}
