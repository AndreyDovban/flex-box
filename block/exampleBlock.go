package block

import (
	"flexbox/styles"
	"flexbox/widgets"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func ExampleBlock() *fyne.Container {

	left := canvas.NewRectangle(nil)
	left.StrokeColor = styles.Grey_600
	left.StrokeWidth = 1
	left.CornerRadius = 8
	left.SetMinSize(fyne.NewSize(200, 40))

	elem := canvas.NewRectangle(styles.Grey_800)
	elem.StrokeColor = styles.Grey_600
	elem.StrokeWidth = 4
	elem.CornerRadius = 8
	elem.Move(fyne.NewPos(0, 500))
	elem.Resize(fyne.NewSize(212, 30))

	bas := canvas.NewRectangle(styles.Grey_800)
	bas.Move(fyne.NewPos(203, 500))
	bas.Resize(fyne.NewSize(12, 30))

	example := container.NewWithoutLayout(elem, bas)

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

	center := canvas.NewRectangle(styles.Grey_800)
	center.StrokeColor = styles.Grey_600
	center.StrokeWidth = 4
	center.CornerRadius = 8

	content := container.New(
		layout.NewCustomPaddedLayout(32, 32, 32, 32),
		container.NewBorder(nil, nil,
			container.NewStack(
				left,
				list,
				example,
			),
			nil,
			center,
		),
	)

	fmt.Println(elem.MinSize())
	fmt.Println(elem.Size())
	fmt.Println(elem.Position().Y)
	fmt.Println(elem.Position().X)

	return content
}
