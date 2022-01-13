package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func reminder() {

	w := fyne.CurrentApp().NewWindow("Reminder")
	input := widget.NewMultiLineEntry()
	input.SetPlaceHolder("Set your description for remainder")

	content := container.NewVBox(widget.NewLabel("Description"), input)

	input1 := widget.NewEntry()
	input1.SetPlaceHolder("Seconds")

	content1 := container.NewVBox(widget.NewLabel("Time "), input1)
	input2 := widget.NewEntry()
	input2.SetPlaceHolder("Minutes")
	input3 := widget.NewEntry()
	input3.SetPlaceHolder("Hours")
	contentSave := container.NewVBox(widget.NewButton("Save", func() {
		log.Println("Content was:", input.Text, input1.Text)
		var timeout = input1.Text
		i, err := strconv.Atoi(timeout)
		if err != nil {
			fmt.Println(err)
		}
		var timeout1 = input2.Text
		i1, err := strconv.Atoi(timeout1)
		if err != nil {
			fmt.Println(err)
		}
		var timeout2 = input3.Text
		i2, err := strconv.Atoi(timeout2)
		if err != nil {
			fmt.Println(err)
		}
		var i3 = (i + i1*60 + i2*600)
		go showNotification(myApp, i3, input.Text)
	}))

	w.SetContent(container.NewVBox(content, content1, input2, input3, contentSave))
	w.Resize(fyne.NewSize(450, 400))
	w.Show()

}

func showNotification(a fyne.App, timeout int, s string) {
	time.Sleep(time.Second * time.Duration(timeout))
	a.SendNotification(fyne.NewNotification("Reminder", s))
}
