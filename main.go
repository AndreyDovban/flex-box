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

func main() {
	app := app.New()
	mainWindow := app.NewWindow("Learn")

	direction := binding.NewString()
	direction.Set("row")
	dv, _ := direction.Get()

	align := binding.NewString()
	align.Set("start")
	av, _ := align.Get()

	justify := binding.NewString()
	justify.Set("start")
	jv, _ := justify.Get()

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

	but := button.NewMyListItemWidget("Hello", "Click")
	but.Resize(fyne.NewSize(120, 40))

	flexLayout := flex.NewFlexBox(dv, av, jv, 10, 10)
	flexBlock := container.New(
		flexLayout,
		but,
		text3,
		p,
		text2,
	)

	direction.AddListener(binding.NewDataListener(func() {
		dv, _ = direction.Get()
		av, _ = align.Get()
		jv, _ = justify.Get()
		log.Println(dv, av, jv)
		flexLayout := flex.NewFlexBox(dv, av, jv, 10, 10)
		flexBlock = container.New(
			flexLayout,
			but,
			text3,
			p,
			text2,
		)
		flexBlock.Refresh()

	}))

	buttons := container.NewVBox(
		widget.NewButton("Direction row", func() { direction.Set("row") }),
		widget.NewButton("Direction column", func() { direction.Set("column") }),
		widget.NewButton("Align start", func() { align.Set("start") }),
		widget.NewButton("Align center", func() { align.Set("center") }),
		widget.NewButton("Align end", func() { align.Set("end") }),
		widget.NewButton("Justify start", func() { justify.Set("start") }),
		widget.NewButton("Justify center", func() { justify.Set("center") }),
		widget.NewButton("Justify emd", func() { justify.Set("end") }),
		widget.NewButton("Justify around", func() { justify.Set("around") }),
		widget.NewButton("Justify between", func() { justify.Set("between") }),
	)

	mainWindow.SetContent(
		container.NewBorder(
			nil, nil, buttons, nil, flexBlock))

	mainWindow.CenterOnScreen()
	mainWindow.Resize(fyne.NewSize(800, 400))
	mainWindow.Show()

	app.Run()
}
