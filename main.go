package main

import (
	"image"
	"image/color"
	"strconv"
	"time"

	"github.com/getlantern/systray"

	"github.com/shirou/gopsutil/v3/cpu" // only cpu so far
)

var TextColor color.Color = image.Black

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	//systray.SetIcon()
	//systray.SetTitle(")
	//systray.SetTooltip("")
	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")
	mColor := systray.AddMenuItem("Toggle color", "Change text's color to red or back")

	mQuit.SetTitle("Exit")
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	go func() {
		for range ticker.C {
			generateReport()
		}
	}()
	//mQuit.SetIcon(icon.Data)
	for {
		select {
		case <-mQuit.ClickedCh:
			systray.Quit()
		case <-mColor.ClickedCh:
			if TextColor == (color.RGBA{R: 255, G: 0, B: 0, A: 255}) {
				TextColor = color.RGBA{R: 0, G: 0, B: 0, A: 255}
			} else {
				TextColor = color.RGBA{R: 255, G: 0, B: 0, A: 255}
			}

		} // jank but future proof
	}
}

func generateReport() {
	c, err := cpu.Percent(0, false)
	if err != nil {
		return
	}

	iconBytes := generateIcon(strconv.FormatFloat(c[0], 'f', 0, 64))
	if err != nil {
		return
	}

	systray.SetIcon(iconBytes)
	// compiler is mad for some reason but it compiles...
}

func onExit() {
	// clean up here
}
