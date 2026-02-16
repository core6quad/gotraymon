package main

import (
	"image"
	"image/color"
	"image/draw"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

func drawBaseIcon(text string) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, 16, 16))

	draw.Draw(img, img.Bounds(), &image.Uniform{color.Transparent}, image.Point{}, draw.Src)

	d := &font.Drawer{
		Dst:  img,
		Src:  image.Black,
		Face: basicfont.Face7x13,
	}

	textWidth := d.MeasureString(text).Round()
	x := (16 - textWidth) / 2
	d.Dot = fixed.P(x, 12)

	d.DrawString(text)

	return img
}
