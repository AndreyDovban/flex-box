package block

import (
	"flexbox/styles"
	"flexbox/widgets"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func ExampleBlock(colorTheme binding.Bool) *fyne.Container {
	changeThemeBut := widget.NewButton("toggle color", func() {
		v, _ := colorTheme.Get()
		colorTheme.Set(!v)

	})
	changeThemeBut.Importance = widget.HighImportance

	left := canvas.NewRectangle(nil)
	left.CornerRadius = 8
	left.SetMinSize(fyne.NewSize(200, 40))

	/*++++++++++++++++++++++++++++++++++++++++++++++++++*/

	names := []string{"Jon", "Bob", "Kate", "Frank", "Andrey"}

	target := binding.NewInt()
	target.Set(-1)

	showText := widget.NewLabel("")

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
				elem.SetSelect()
			} else {
				elem.UnSelect()
			}

		},
	)

	list.OnSelected = func(id widget.ListItemID) {
		target.Set(id)
		list.RefreshItem(id)
		showText.SetText(names[id])

	}
	list.HideSeparators = true
	list.Resize(fyne.NewSize(210, 400))

	center := canvas.NewRectangle(styles.Grey_800)
	center.StrokeColor = styles.Grey_600
	center.StrokeWidth = 1
	center.CornerRadius = 8

	content := container.New(
		layout.NewCustomPaddedLayout(32, 32, 32, 32),
		container.NewBorder(nil, nil,
			container.NewStack(
				left,
				container.NewVBox(
					changeThemeBut,
					container.NewWithoutLayout(
						list,
					),
				),
			),
			nil,
			center,
		),
	)

	colorTheme.AddListener(binding.NewDataListener(func() {
		v, _ := colorTheme.Get()
		if v {
			center.FillColor = styles.Grey_100
		} else {
			center.FillColor = styles.Grey_800
		}

	}))

	return content
}
