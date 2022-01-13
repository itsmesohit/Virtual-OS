package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func showIim(clock *widget.Label) {
	formatted := time.Now().Format("03:04:05")
	clock.SetText(formatted)
}
func clock(a fyne.App) {
	// a := app.New()
	w := a.NewWindow("Clock")
	clock := widget.NewLabel("")
	showIime(clock)
	w.SetContent(clock)
	go func() {
		t := time.NewTicker(time.Second)
		for range t.C {
			showIime(clock)
		}
	}()
	w.Show()
}
