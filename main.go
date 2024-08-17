package main

import (
	"encoding/json"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("JSON Converter")

	inputEntry := widget.NewMultiLineEntry()
	inputEntry.SetPlaceHolder("输入原始 JSON 字符串...")

	resultLabel := widget.NewLabel("converted JSON will be displayed here")
	copyButton := widget.NewButton("copy result", func() {
		myWindow.Clipboard().SetContent(resultLabel.Text)
	})

	// listen to input changes
	inputEntry.OnChanged = func(s string) {
		var data interface{}
		err := json.Unmarshal([]byte(s), &data)
		if err != nil {
			resultLabel.SetText("invalid JSON string")
			return
		}

		prettyJSON, _ := json.MarshalIndent(data, "", "    ")

		resultLabel.SetText(string(prettyJSON))
	}

	// layout
	leftContainer := container.NewVBox(inputEntry)
	rightContainer := container.NewVBox(resultLabel, copyButton)
	mainContainer := container.NewHSplit(leftContainer, rightContainer)

	myWindow.SetContent(mainContainer)
	myWindow.Resize(fyne.NewSize(800, 600))
	myWindow.ShowAndRun()
}
