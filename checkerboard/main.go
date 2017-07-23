package main

import (
	"flag"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"

	"golang.org/x/image/colornames"
)

func main() {
	border := flag.Bool("b", false, "generate border")
	flag.Parse()

	const width, height = 128, 32
	const filename = "checkerboard.png"

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	generateCheckerboard(img)
	if *border {
		generateBorder(img, colornames.White)
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

func generateCheckerboard(img *image.RGBA) {
	log.Println("Generating Checkerboard...")
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			if (x+y)%2 == 0 {
				img.Set(x, y, colornames.White)
			} else {
				img.Set(x, y, colornames.Black)
			}

		}
	}
}

func generateBorder(img *image.RGBA, col color.Color) {
	log.Println("Generating Border...")
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		img.Set(img.Bounds().Min.X, y, col)
		img.Set(img.Bounds().Max.X-1, y, col)
	}
	for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
		img.Set(x, img.Bounds().Min.Y, col)
		img.Set(x, img.Bounds().Max.Y-1, col)
	}
}
