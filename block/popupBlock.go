package block

import (
	"flexbox/widgets"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func PopupBlock(app fyne.App) *fyne.Container {

	myWindowPop := app.NewWindow("CanvasPop")

	p := widget.NewPopUp(
		widget.NewLabel("Hello"),
		myWindowPop.Canvas(),
	)
	p.Hide()

	elem := widgets.NewButton("click", func() {
		p.Show()
	})

	content := container.NewCenter(
		container.New(
			layout.NewCustomPaddedHBoxLayout(32),
			elem,
			p,
		),
	)

	return content
}
