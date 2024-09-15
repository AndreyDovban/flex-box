package main

import (
	"flexbox/flex"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	app := app.New()
	mainWindow := app.NewWindow("Learn")

	text1 := widget.NewLabel("First Label")
	text1.Wrapping = fyne.TextWrapBreak
	text2 := widget.NewLabel("Middle Label")
	text3 := widget.NewLabel("bottomright")

	p := widget.NewRichText(
		&widget.TextSegment{
			Text: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."})

	p.Resize(fyne.NewSize(160, 100))

	p.Wrapping = fyne.TextWrapBreak
	p.Truncation = fyne.TextTruncateEllipsis

	flex := &flex.FlexBox{
		Dir:     "column",
		Align:   "center",
		Justify: "center",
		Gap:     10,
		Padding: 10,
	}

	block := container.New(
		flex,
		text3,
		p,
		text2,
	)

	mainWindow.SetContent(block)

	mainWindow.CenterOnScreen()
	mainWindow.Resize(fyne.NewSize(800, 400))
	mainWindow.Show()

	app.Run()
}
