package widgets

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Label struct {
	widget.BaseWidget
	Title *widget.Label
}

func NewLabel(title string) *Label {
	item := &Label{
		Title: widget.NewLabel(title),
	}
	item.ExtendBaseWidget(item)

	return item
}

func (item *Label) CreateRenderer() fyne.WidgetRenderer {
	label := item.Title
	label.Move(fyne.NewPos(-10, 0))
	label.Wrapping = fyne.TextWrapWord

	label.Resize(item.Size())

	c := container.NewWithoutLayout(label)

	return widget.NewSimpleRenderer(c)
}
