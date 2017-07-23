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
	img = generateCheckerboard(img)

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

func generateCheckerboard(img *image.RGBA) *image.RGBA {
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			if (x+y)%2 == 0 {
				img.Set(x, y, colornames.White)
			} else {
				img.Set(x, y, colornames.Black)
			}

		}
	}
	return img
}
