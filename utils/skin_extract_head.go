package utils

import (
	"image"
)

// Return the 8x8 head of a 64x32 skin
func SkinExtractHead(img image.Image) image.Image {
	img2 := image.NewRGBA(image.Rect(0, 0, 8, 8))

	for x := 8; x < 16; x += 1 {
		for y := 8; y < 16; y += 1 {
			img2.Set(x-8, y-8, img.At(x, y))
		}
	}

	return img2
}
