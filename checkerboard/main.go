package main

import (
	"image"
	"image/png"
	"log"
	"os"

	"golang.org/x/image/colornames"
)

func main() {
	log.Println("Generating Checkerboard...")
	const width, height = 128, 32
	const filename = "checkerboard.png"

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if (x+y)%2 == 0 {
				img.Set(x, y, colornames.White)
			} else {
				img.Set(x, y, colornames.Black)
			}

		}
	}

	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	if err := png.Encode(f, img); err != nil {
		f.Close()
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
	log.Printf("Saved to image [%s]", filename)
}
