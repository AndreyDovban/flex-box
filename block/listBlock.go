package block

import (
	"flexbox/widgets"
	"image/color"
	"log"

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

	names := binding.NewStringList()
	names.Set([]string{"Jon", "Bob", "Kate", "Frank", "Andrey"})

	addElem := widget.NewButton("Add elem", func() {
		names.Append("New text")
	})

	list := widget.NewListWithData(
		names,
		func() fyne.CanvasObject { return widget.NewLabel("") },
		func(di binding.DataItem, co fyne.CanvasObject) {
			txt, _ := di.(binding.String).Get()
			co.(*widget.Label).SetText(txt)
		},
	)
	list.Resize(fyne.NewSize(200, 400))
	list.OnSelected = func(id widget.ListItemID) {
		s, _ := names.GetItem(id)
		text, _ := s.(binding.String).Get()
		log.Println(text)

	}

	listContent := container.NewCenter(
		container.NewStack(
			widgets.Fon(color.Transparent, color.White, 8),
			container.New(
				layout.NewGridWrapLayout(fyne.NewSize(400, 400)),
				container.New(layout.NewCustomPaddedLayout(16, 16, 16, 16),
					container.New(layout.NewCustomPaddedVBoxLayout(24),
						addElem,
						hr,
						container.NewGridWrap(
							fyne.NewSize(400-32, 260),
							list,
						),
					),
				),
			),
		),
	)

	return listContent
}
