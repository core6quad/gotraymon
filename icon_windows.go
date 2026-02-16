//go:build windows

package main

import (
	"bytes"

	ico "github.com/Kodeworks/golang-image-ico"
)

func generateIcon(text string) []byte {
	var img = drawBaseIcon(text)

	var buf bytes.Buffer
	if err := ico.Encode(&buf, img); err != nil {
		panic(err)
	}

	return buf.Bytes()
}
