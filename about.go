package main

import (
	"os/exec"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func about() {

	w := fyne.CurrentApp().NewWindow("About")
	image := canvas.NewImageFromFile("./img/about.png")

	a1, _ := fyne.LoadResourceFromPath("./img/btn2.png")
	btn1 := widget.NewButtonWithIcon("LinkedIn", a1, func() {
		exec.Command("rundll32", "url.dll,FileProtocolHandler", "https://www.linkedin.com/in/sohitkushwaha/").Start()
	})
	btn1.Resize(fyne.NewSize(110, 40))
	btn1.Move(fyne.NewPos(315, 550))

	a2, _ := fyne.LoadResourceFromPath("./img/btn2.png")
	btn2 := widget.NewButtonWithIcon("GitHub", a2, func() {
		exec.Command("rundll32", "url.dll,FileProtocolHandler", "https://github.com/sohitdon").Start()
	})
	btn2.Resize(fyne.NewSize(105, 40))
	btn2.Move(fyne.NewPos(435, 550))

	a3, _ := fyne.LoadResourceFromPath("./img/btn2.png")
	btn3 := widget.NewButtonWithIcon("Nados", a3, func() {
		exec.Command("rundll32", "url.dll,FileProtocolHandler", "https://nados.pepcoding.com/profile/a61e1154-5f2f-4740-9b13-4a7d6a4f334e").Start()
	})
	btn3.Resize(fyne.NewSize(150, 40))
	btn3.Move(fyne.NewPos(315, 620))

	w.Resize(fyne.NewSize(600, 700))
	w.CenterOnScreen()
	w.SetContent(container.NewMax(image,
		container.NewWithoutLayout(
			btn1, btn2, btn3,
		),
	))
	w.Show()
}
