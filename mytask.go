package main

import (
	"encoding/json"
	"io/ioutil"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func mytask() {
	type Tasks struct {
		Task_des string
		Date     string
	}

	var myTaskData []Tasks
	data_from_file, _ := ioutil.ReadFile("Todo.txt")
	json.Unmarshal(data_from_file, &myTaskData)

	w := fyne.CurrentApp().NewWindow("My Task")
	w.Resize(fyne.NewSize(800, 400))

	l_task := widget.NewLabel("...")
	l_task.TextStyle = fyne.TextStyle{Bold: true}

	l_date := widget.NewLabel("...")

	e_task := widget.NewEntry()
	e_task.SetPlaceHolder("Enter Your Task...")
	e_date := widget.NewEntry()
	e_date.SetPlaceHolder("Due Date...")
	submit_btn := widget.NewButton("Submit", func() {
		myData1 := &Tasks{
			Task_des: e_task.Text,
			Date:     e_date.Text,
		}
		myTaskData = append(myTaskData, *myData1)
		final_data, _ := json.MarshalIndent(myTaskData, "", " ")
		ioutil.WriteFile("Todo.txt", final_data, 0644)

		e_task.Text = ""
		e_date.Text = ""

		e_task.Refresh()
		e_date.Refresh()
	})

	del_button := widget.NewButton("Del", func() {
		var TempData []Tasks
		for _, e := range myTaskData {

			if l_task.Text != e.Task_des {
				TempData = append(TempData, e)
			}
		}
		myTaskData = TempData
		result, _ := json.MarshalIndent(myTaskData, "", " ")
		ioutil.WriteFile("Todo.txt", result, 0644)
	})
	list := widget.NewList(
		func() int { return len(myTaskData) },
		func() fyne.CanvasObject { return widget.NewLabel("") },
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			co.(*widget.Label).SetText(myTaskData[lii].Task_des)
		},
	)
	list.OnSelected = func(id widget.ListItemID) {
		l_task.Text = myTaskData[id].Task_des
		l_task.Refresh()

		l_date.Text = myTaskData[id].Date
		l_date.Refresh()
	}
	update_button := widget.NewButton("Update", func() {
		var TempData []Tasks
		update := &Tasks{
			Task_des: e_task.Text,
			Date:     e_date.Text,
		}
		for _, e := range myTaskData {
			if l_task.Text == e.Task_des {
				TempData = append(TempData, *update)
			} else {
				TempData = append(TempData, e)
			}
		}
		myTaskData = TempData
		result, _ := json.MarshalIndent(myTaskData, "", " ")
		ioutil.WriteFile("Todo.txt", result, 0644)

		e_task.Text = ""
		e_date.Text = ""
		e_task.Refresh()
		e_date.Refresh()
		list.Refresh()
	})
	task := widget.NewLabel("Your Task is...")
	date := widget.NewLabel("Due Date is...")
	w.SetContent(
		container.NewHSplit(
			list,
			container.NewVBox(
				task, l_task, date, l_date, e_task,
				e_date, submit_btn, del_button,
				update_button,
			),
		),
	)
	w.Show()
}
