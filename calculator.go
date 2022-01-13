package main

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Knetic/govaluate"
)

func calculator() {

	output := " "
	input := widget.NewLabel(output)
	historyStr := " "
	historyin := widget.NewLabel(historyStr)
	var historyArr []string
	ishistory := false
	historyBtn := widget.NewButton("History", func() {
		if ishistory {
			historyStr = ""
		} else {
			for i := len(historyArr) - 1; i >= 0; i-- {
				historyStr = historyStr + historyArr[i]
				historyStr += "\n"
			}
		}
		ishistory = !ishistory
		historyin.SetText(historyStr)
	})
	backBtn := widget.NewButton("Back", func() {
		output = output[:len(output)-1]
		input.SetText(output)
	})
	clearBtn := widget.NewButton("Clear", func() {
		if len(output) > 0 {
			output = ""
			input.SetText(output)
		}
	})
	openBracketBtn := widget.NewButton("(", func() {
		output = output + "("
		input.SetText(output)
	})
	closeBracketBtn := widget.NewButton(")", func() {
		output = output + ")"
		input.SetText(output)
	})
	devideBtn := widget.NewButton("/", func() {
		output = output + "/"
		input.SetText(output)
	})
	sevenBtn := widget.NewButton("7", func() {
		output = output + "7"
		input.SetText(output)
	})
	eightBtn := widget.NewButton("8", func() {
		output = output + "8"
		input.SetText(output)
	})
	ninedeBtn := widget.NewButton("9", func() {
		output = output + "9"
		input.SetText(output)
	})
	multideBtn := widget.NewButton("*", func() {
		output = output + "*"
		input.SetText(output)
	})
	fourBtn := widget.NewButton("4", func() {
		output = output + "4"
		input.SetText(output)
	})
	fiveBtn := widget.NewButton("5", func() {
		output = output + "5"
		input.SetText(output)
	})
	sixeBtn := widget.NewButton("6", func() {
		output = output + "6"
		input.SetText(output)
	})
	minusdeBtn := widget.NewButton("-", func() {
		output = output + "-"
		input.SetText(output)
	})
	oneBtn := widget.NewButton("1", func() {
		output = output + "1"
		input.SetText(output)
	})
	twoBtn := widget.NewButton("2", func() {
		output = output + "2"
		input.SetText(output)
	})
	threeBtn := widget.NewButton("3", func() {
		output = output + "3"
		input.SetText(output)
	})
	plusBtn := widget.NewButton("+", func() {
		output = output + "+"
		input.SetText(output)
	})
	zeroBtn := widget.NewButton("0", func() {
		output = output + "0"
		input.SetText(output)
	})
	dotBtn := widget.NewButton(".", func() {
		output = output + "."
		input.SetText(output)
	})
	equalBtn := widget.NewButton("=", func() {
		expression, err := govaluate.NewEvaluableExpression(output)
		if err == nil {
			result, err := expression.Evaluate(nil)
			if err == nil {
				ans := strconv.FormatFloat(result.(float64), 'f', -1, 64)
				strToAppend := output + "= " + ans
				historyArr = append(historyArr, strToAppend)
				output = ans
			} else {
				output = "Error"
			}
		} else {
			output = "Error"
		}
		input.SetText(output)

	})

	w := myApp.NewWindow("Calculator")
	w.Resize(fyne.NewSize(500, 280))

	w.SetContent(container.NewVBox(
		input,
		historyin,
		container.NewGridWithColumns(1,
			container.NewGridWithColumns(2,
				historyBtn,
				backBtn,
			),
			container.NewGridWithColumns(4,
				clearBtn,
				openBracketBtn,
				closeBracketBtn,
				devideBtn),
			container.NewGridWithColumns(4,
				ninedeBtn,
				eightBtn,
				sevenBtn,
				multideBtn,
			),
			container.NewGridWithColumns(4,
				fourBtn,
				fiveBtn,
				sixeBtn,
				minusdeBtn,
			),
			container.NewGridWithColumns(4,
				oneBtn,
				twoBtn,
				threeBtn,
				plusBtn,
			),
			container.NewGridWithColumns(2,
				container.NewGridWithColumns(2,
					zeroBtn,
					dotBtn,
				),
				equalBtn,
			),
		),
	))
	w.Show()
}
