//go:build windows

package main

import (
	"bytes"
	"image"

	ico "github.com/Kodeworks/golang-image-ico"
)

func generateIcon(text string) []byte {
	var img image.Image //:= drawBaseIcon(text)

	if ShowCpu {
		img = drawBaseIcon(text)
	} else {
		img = drawBaseIconWithPercent(text)
	}

	var buf bytes.Buffer
	if err := ico.Encode(&buf, img); err != nil {
		panic(err)
	}

	return buf.Bytes()
}
