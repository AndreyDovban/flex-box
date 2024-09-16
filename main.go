package main

import (
	"flexbox/button"
	"flexbox/flex"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

/**
ToDo
1. Изменить привязанные данные на мапу
2. Уьрать перерисовку всего окна
*/

func main() {
	app := app.New()
	mainWindow := app.NewWindow("Learn")

	data := binding.NewString()

	direction := "row"
	align := "start"
	justify := "start"

	text1 := widget.NewLabel("First Label")
	text1.Wrapping = fyne.TextWrapBreak
	text2 := widget.NewLabel("Middle Label")
	text3 := widget.NewLabel("bottomright")

	s := widget.NewLabel("Target")

	p := widget.NewRichText(
		&widget.TextSegment{
			Text: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."})

	p.Resize(fyne.NewSize(160, 100))

	p.Wrapping = fyne.TextWrapBreak
	p.Truncation = fyne.TextTruncateEllipsis

	but := button.NewMyListItemWidget("Hello", "Click")
	but.Resize(fyne.NewSize(120, 40))

	flexLayout := flex.NewFlexBox(direction, align, justify, 10, 10)
	flexBlock := container.New(
		flexLayout,
		but,
		text3,
		p,
		text2, s,
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

	mainWindow.SetContent(ttt)

	mainWindow.CenterOnScreen()
	mainWindow.Resize(fyne.NewSize(1000, 600))
	mainWindow.Show()

	app.Run()
}
