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

	input := widget.NewMultiLineEntry()
	input.SetPlaceHolder("paste JSON string here")
	input.TextStyle.Bold = true

	result := widget.NewMultiLineEntry()
	result.SetPlaceHolder("converted result will be shown here")
	result.TextStyle.Bold = true

	copyButton := widget.NewButton("copy result", func() {
		myWindow.Clipboard().SetContent(result.Text)
	})

	// listen to input changes
	input.OnChanged = func(s string) {
		var data interface{}
		err := json.Unmarshal([]byte(s), &data)
		if err != nil {
			result.SetText("invalid JSON string")
			return
		}

		prettyJSON, _ := json.MarshalIndent(data, "", "    ")

		result.SetText(string(prettyJSON))
	}

	// layout
	textContainer := container.NewHSplit(input, result)
	mainContainer := container.NewBorder(nil, copyButton, nil, nil, textContainer)

	myWindow.SetContent(mainContainer)
	myWindow.Resize(fyne.NewSize(1000, 500))
	myWindow.CenterOnScreen()
	myWindow.ShowAndRun()
}
