package main

import (
	"bytes"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"log"
)

const (
	srcImage = "assets/chloe.png"
	fontFile = "assets/LiberationSans-Regular.ttf"
)

var srcPng image.Image
var srcFont *truetype.Font

// Load source image and font on initialisation
func init() {
	srcPng, _ = loadImage()
	srcFont, _ = loadFont()
}

// Load source image from assets
func loadImage() (image.Image, error) {
	imgFile, _ := Asset(srcImage)
	reader := bytes.NewReader(imgFile)

	src, err := png.Decode(reader)

	return src, err
}

// Load font from assets
func loadFont() (*truetype.Font, error) {
	fontBytes, err := Asset(fontFile)
	if err != nil {
		log.Fatal(err)
	}

	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
		log.Fatal(err)
	}

	return f, err
}

// Stamp puts a given name onto the source image and writes a JPEG into the provided buffer
func Stamp(buffer *bytes.Buffer, name string) {
	// Get source image and font
	src := srcPng
	f := srcFont

	// Get source image bounds
	bounds := src.Bounds()

	// Create a new image based on the source dimensions
	img := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))

	// Draw the source image onto the target image
	draw.Draw(img, img.Bounds(), src, image.ZP, draw.Src)

	// Font settings
	fg := image.NewUniform(color.RGBA{R: 83, G: 49, B: 95, A: 255})
	var fontSize float64 = 20

	// Font context
	c := freetype.NewContext()
	c.SetDPI(72)
	c.SetFont(f)
	c.SetFontSize(fontSize)
	c.SetClip(img.Bounds())
	c.SetDst(img)
	c.SetSrc(fg)

	// Write outside the visible area to get rendered text width
	pt := freetype.Pt(2000, 2000)
	width, err := c.DrawString(name, pt)
	if err != nil {
		log.Fatal(err)
	}

	// Put name right-aligned to 346,410
	pt2 := freetype.Pt(346+(2000-width.X.Round()), 410)
	_, err = c.DrawString(name, pt2)
	if err != nil {
		log.Fatal()
	}

	// Encode as JPEG and write to buffer
	err = jpeg.Encode(buffer, img, &jpeg.Options{80})
	if err != nil {
		log.Fatal(err)
	}
}
