//go:build !windows

package main

import (
	"bytes"
	"image"
	"image/png"
)

func generateIcon(text string) []byte {

	var img image.Image //:= drawBaseIcon(text)

	if ShowCpu {
		img = drawBaseIcon(text)
	} else {
		img = drawBaseIconWithPercent(text)
	}

	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		panic(err)
	}

	return buf.Bytes()
}
