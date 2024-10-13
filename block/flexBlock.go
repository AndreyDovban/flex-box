package block

import (
	"flexbox/flex"
	"flexbox/widgets"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

/** Flex block */
func FlexBlock() *fyne.Container {

	data := binding.NewString()

	direction := "column"
	align := "center"
	justify := "center"

	num := widgets.NewNumericalEntry()
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

	flexTabContent := container.NewBorder(
		nil, nil, buttons, nil, flexBlock)

	data.AddListener(binding.NewDataListener(func() {
		flexTabContent.Remove(flexBlock)
		flexLayout.Dir = direction
		flexLayout.Align = align
		flexLayout.Justify = justify

		flexTabContent.Add(flexBlock)
		flexTabContent.Refresh()
	}))

	return flexTabContent
}
