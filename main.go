package main

import (
	"image"
	"image/color"
	"strconv"
	"time"

	"github.com/getlantern/systray"

	"github.com/shirou/gopsutil/v3/cpu" // only cpu so far
	"github.com/shirou/gopsutil/v3/mem"
)

var TextColor color.Color = image.Black
var ShowCpu bool = true // if false show memory usage

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	//systray.SetIcon()
	//systray.SetTitle(")
	//systray.SetTooltip("")
	mColor := systray.AddMenuItem("Toggle color (Red/black)", "Change text's color to red or back")
	mMon := systray.AddMenuItem("Toggle stat (RAM/CPU)", "Change what is displayed")
	systray.AddSeparator()
	mT := systray.AddMenuItem("CPU", "")
	systray.AddSeparator()
	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	mT.Disable()
	updateDisplay(*mT)

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
		case <-mMon.ClickedCh:
			ShowCpu = !ShowCpu
			updateDisplay(*mT)
		}
	}
}

func updateDisplay(mT systray.MenuItem) {
	if ShowCpu {
		mT.SetTitle("showing CPU")
	} else {
		mT.SetTitle("showing MEM")
	}
}

func generateReport() {
	c, err := cpu.Percent(0, false)
	if err != nil {
		return
	}
	var iconBytes []byte
	if ShowCpu {
		iconBytes = generateIcon(strconv.FormatFloat(c[0], 'f', 0, 64))
		if err != nil {
			return
		}
	} else {
		v, _ := mem.VirtualMemory()
		iconBytes = generateIcon(strconv.FormatFloat(v.UsedPercent, 'f', 0, 64))
		if err != nil {
			return
		}
	}
	systray.SetIcon(iconBytes)
	// compiler is mad for some reason but it compiles...
}

func onExit() {
	// clean up here, idk
}
