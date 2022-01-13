package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var button1 fyne.Widget
var button2 fyne.Widget
var button3 fyne.Widget
var button4 fyne.Widget
var button5 fyne.Widget
var date fyne.Widget
var homeBtn *fyne.Container
var myApp fyne.App = app.New()
var w fyne.Window = myApp.NewWindow("Virtual OS")
var img = canvas.NewImageFromFile("./img/bg.jpg")

// var image = canvas.NewImageFromFile("./img/Virtual.png")

func showIime(clock *widget.Label) {
	formatted := time.Now().Format("03:04:05")
	clock.SetText(formatted)
}

func main() {

	w := myApp.NewWindow("Pep OS")

	s, _ := fyne.LoadResourceFromPath("./icon/1.png")
	w.SetIcon(s)
	w.Resize(fyne.NewSize(1480, 760))
	image := canvas.NewImageFromFile("./img/os1.jpg")

	image.FillMode = canvas.ImageFillStretch

	r, _ := fyne.LoadResourceFromPath("./img/home.png")
	button1 := widget.NewButtonWithIcon("", r, func() {

		image = canvas.NewImageFromFile("./img/wall.png")
		image.Refresh()
	})

	button1.Resize(fyne.NewSize(50, 50))
	button1.Move(fyne.NewPos(550, 700))

	r1, _ := fyne.LoadResourceFromPath("./img/btn3.png")
	button2 := widget.NewButtonWithIcon("", r1, func() {
		go productivity()
	})
	button2.Resize(fyne.NewSize(50, 50))
	button2.Move(fyne.NewPos(610, 700))

	r2, _ := fyne.LoadResourceFromPath("./img/thispc.png")
	button3 := widget.NewButtonWithIcon("", r2, func() {
		go intertainement()
	})
	button3.Resize(fyne.NewSize(50, 50))
	button3.Move(fyne.NewPos(670, 700))

	r3, _ := fyne.LoadResourceFromPath("./img/game.png")
	button4 := widget.NewButtonWithIcon("", r3, func() {
		go games()
	})
	button4.Resize(fyne.NewSize(50, 50))
	button4.Move(fyne.NewPos(730, 700))

	r4, _ := fyne.LoadResourceFromPath("./img/btn4.png")
	button5 := widget.NewButtonWithIcon("", r4, func() {
		go work()

	})
	button5.Resize(fyne.NewSize(50, 50))
	button5.Move(fyne.NewPos(790, 700))

	formatted := time.Now().Format("01-02-2006")
	date = widget.NewButton(formatted, func() {
		fmt.Println("Date")
		go reminder()
	})

	date.Resize(fyne.NewSize(120, 30))
	date.Move(fyne.NewPos(1280, 20))

	clock := widget.NewLabel("")
	showIime(clock)
	go func() {
		t := time.NewTicker(time.Second)
		for range t.C {
			showIime(clock)
		}
	}()
	homeBtn = container.NewMax(
		image,
		container.NewWithoutLayout(
			button1, button2, button3, button4, button5, date, clock,
		),
	)
	w.SetContent(homeBtn)

	w.CenterOnScreen()
	w.ShowAndRun()
}
