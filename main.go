package main

import (
	"flexbox/flex"
	"flexbox/mywidgets"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

/**
ToDo
1. Цифровой ввод
2.
*/

func main() {
	app := app.New()
	mainWindow := app.NewWindow("Learn")

	data := binding.NewString()

	direction := "column"
	align := "center"
	justify := "center"

	vid := mywidgets.NewMyWidget("click")
	vid.Resize(fyne.NewSize(200, 40))

	// vid := widget.NewLabel("click")
	// vid.Resize(fyne.NewSize(200, 40))

	num := mywidgets.NewNumericalEntry()
	num.Resize(fyne.NewSize(200, 40))

	text1 := widget.NewLabel("First Label")
	text1.Wrapping = fyne.TextWrapBreak
	text2 := widget.NewLabel("Middle Label")
	text3 := widget.NewLabel("bottomright")

	paragraf := widget.NewRichText(
		&widget.TextSegment{
			Text: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."})
	paragraf.Resize(fyne.NewSize(160, 100))
	paragraf.Wrapping = fyne.TextWrapBreak
	paragraf.Truncation = fyne.TextTruncateEllipsis

	flexLayout := flex.NewFlexBox(direction, align, justify, 10, 10)
	flexBlock := container.New(
		flexLayout,
		num,
		vid,
		text3,
		paragraf,
		text2,
	)

	buttons := container.NewVBox(
		widget.NewButton("Direction row", func() { data.Set("row"); direction = "row" }),
		widget.NewButton("Direction column", func() { data.Set("column"); direction = "column" }),
		widget.NewButton("Align start", func() { data.Set("start"); align = "start" }),
		widget.NewButton("Align center", func() { data.Set("center"); align = "center" }),
		widget.NewButton("Align end", func() { data.Set("end"); align = "end" }),
		widget.NewButton("Justify start", func() { data.Set("start"); justify = "start" }),
		widget.NewButton("Justify center", func() { data.Set("center"); justify = "center" }),
		widget.NewButton("Justify emd", func() { data.Set("end"); justify = "end" }),
		widget.NewButton("Justify around", func() { data.Set("around"); justify = "around" }),
		widget.NewButton("Justify between", func() { data.Set("between"); justify = "between" }),
	)

	ttt := container.NewBorder(
		nil, nil, buttons, nil, flexBlock)

	data.AddListener(binding.NewDataListener(func() {

		log.Println(flexBlock.Layout)

		ttt.Remove(flexBlock)
		flexLayout.Dir = direction
		flexLayout.Align = align
		flexLayout.Justify = justify

		ttt.Add(flexBlock)
		ttt.Refresh()
		log.Println(direction, align, justify)
	}))

	ctrlTab := &desktop.CustomShortcut{KeyName: fyne.KeyTab, Modifier: fyne.KeyModifierControl}

	mainWindow.Canvas().AddShortcut(ctrlTab, func(shortcut fyne.Shortcut) {
		log.Println("We tapped Ctrl+Tab")
	})

	mainWindow.SetContent(ttt)

	mainWindow.CenterOnScreen()
	mainWindow.Resize(fyne.NewSize(1000, 600))
	mainWindow.Show()

	app.Run()
}
