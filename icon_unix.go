//go:build !windows

package main

import (
	"bytes"
	"image/png"
)

func generateIcon(text string) []byte {
	img := drawBaseIcon(text)

	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		panic(err)
	}

	return buf.Bytes()
}
