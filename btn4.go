package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func work() {

	w := myApp.NewWindow("Other")

	r1, _ := fyne.LoadResourceFromPath("./icon/10.png")
	button1 := widget.NewButtonWithIcon("Dark Theme", r1, func() {
		img = canvas.NewImageFromFile("./img/dark.png")
		img.Refresh()
		myApp.Settings().SetTheme(theme.DarkTheme())
	})
	button1.Resize(fyne.NewSize(150, 40))
	button1.Move(fyne.NewPos(30, 70))

	r2, _ := fyne.LoadResourceFromPath("./icon/11.png")
	button2 := widget.NewButtonWithIcon("Light Theme", r2, func() {
		img = canvas.NewImageFromFile("./img/bg.png")
		img.Refresh()
		myApp.Settings().SetTheme(theme.LightTheme())
	})
	button2.Resize(fyne.NewSize(150, 40))
	button2.Move(fyne.NewPos(190, 70))

	r3, _ := fyne.LoadResourceFromPath("./icon/12.png")
	button3 := widget.NewButtonWithIcon("About", r3, func() {
		go about()
	})
	button3.Resize(fyne.NewSize(100, 40))
	button3.Move(fyne.NewPos(350, 70))

	w.Resize(fyne.NewSize(550, 420))
	w.CenterOnScreen()
	w.SetContent(container.NewMax(img,
		container.NewWithoutLayout(
			button1, button2,
		),
	))
	w.Show()
}
