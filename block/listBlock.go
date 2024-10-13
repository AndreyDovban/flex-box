package block

import (
	"flexbox/widgets"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

/** List With Data block */
func ListBlock() *fyne.Container {

	hr := canvas.NewLine(color.White)
	hr.Resize(fyne.NewSize(200, 0))

	names := []string{"Jon", "Bob", "Kate", "Frank", "Andrey"}

	target := binding.NewInt()
	target.Set(-1)

	showText := widget.NewLabel("")

	addElem := widget.NewButton("Add elem", func() {
		names = append(names, "New text")
	})

	list := widget.NewList(
		func() int { return len(names) },
		func() fyne.CanvasObject {
			return widgets.NewArrowLabel("")
		},
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			targetId, _ := target.Get()

			elem := co.(*widgets.ArrowLabel)
			elem.Label.SetText(names[lii])

			if targetId == lii {
				elem.SelectedIcon.Show()
			} else {
				elem.SelectedIcon.Hide()
			}

		},
	)

	list.OnSelected = func(id widget.ListItemID) {
		target.Set(id)
		list.RefreshItem(id)
		showText.SetText(names[id])

	}

	list.HideSeparators = true

	unselectBut := widget.NewButton("Unselect", func() {
		showText.SetText("")
		target.Set(-1)
		list.UnselectAll()
		list.FocusLost()

	})

	listContent := container.NewCenter(
		container.NewStack(
			widgets.Fon(color.Transparent, color.White, 8),
			container.New(
				layout.NewGridWrapLayout(fyne.NewSize(400, 500)),
				container.New(layout.NewCustomPaddedLayout(16, 16, 16, 16),
					container.New(layout.NewCustomPaddedVBoxLayout(24),
						showText,
						addElem,
						unselectBut,
						hr,
						container.NewGridWrap(
							fyne.NewSize(400-32, 220),
							list,
						),
					),
				),
			),
		),
	)

	return listContent
}
