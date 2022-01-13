package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func intertainement() {
	w := myApp.NewWindow("This PC")

	r1, _ := fyne.LoadResourceFromPath("./icon/6.png")
	button1 := widget.NewButtonWithIcon("Camera", r1, func() {
		go camera()
	})
	button1.Resize(fyne.NewSize(110, 40))
	button1.Move(fyne.NewPos(30, 70))

	r2, _ := fyne.LoadResourceFromPath("./img/btn2.png")
	button2 := widget.NewButtonWithIcon("Gallery", r2, func() {
		go gallery()
	})
	button2.Resize(fyne.NewSize(105, 40))
	button2.Move(fyne.NewPos(150, 70))

	r3, _ := fyne.LoadResourceFromPath("./icon/7.png")
	button3 := widget.NewButtonWithIcon("Audio Player", r3, func() {
		go music()
	})
	button3.Resize(fyne.NewSize(150, 40))
	button3.Move(fyne.NewPos(265, 70))

	w.Resize(fyne.NewSize(500, 450))
	w.CenterOnScreen()
	w.SetContent(container.NewMax(img,
		container.NewWithoutLayout(
			button1, button2, button3,
		),
	))
	w.Show()
}
