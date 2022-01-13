package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func productivity() {

	w := myApp.NewWindow("Productivity")

	r1, _ := fyne.LoadResourceFromPath("./icon/2.png")
	button1 := widget.NewButtonWithIcon("Notepad", r1, func() {
		go Notepad()
	})
	button1.Resize(fyne.NewSize(130, 40))
	button1.Move(fyne.NewPos(30, 70))

	r2, _ := fyne.LoadResourceFromPath("./icon/3.png")
	button2 := widget.NewButtonWithIcon("Calculator", r2, func() {
		go calculator()
	})
	button2.Resize(fyne.NewSize(130, 40))
	button2.Move(fyne.NewPos(170, 70))

	r3, _ := fyne.LoadResourceFromPath("./icon/4.png")
	button3 := widget.NewButtonWithIcon("My Task", r3, func() {
		go mytask()

	})
	button3.Resize(fyne.NewSize(110, 40))
	button3.Move(fyne.NewPos(310, 70))

	r4, _ := fyne.LoadResourceFromPath("./icon/5.png")
	button4 := widget.NewButtonWithIcon("Reminder", r4, func() {
		go reminder()
	})
	button4.Resize(fyne.NewSize(130, 40))
	button4.Move(fyne.NewPos(430, 70))

	w.Resize(fyne.NewSize(650, 420))
	w.CenterOnScreen()

	w.SetContent(container.NewMax(img,
		container.NewWithoutLayout(
			button1, button2, button3, button4,
		),
	))
	w.Show()
}
