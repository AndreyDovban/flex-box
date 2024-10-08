package widgets

import (
	"flexbox/styles"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type CustomSelect struct {
	widget.BaseWidget
	Select *widget.Select
}

func NewCustomSelect(placeholder string, options []string, changed func(string)) *CustomSelect {
	item := &CustomSelect{
		Select: widget.NewSelect(options, changed),
	}
	item.ExtendBaseWidget(item)
	item.Select.PlaceHolder = placeholder

	return item
}

func (item *CustomSelect) CreateRenderer() fyne.WidgetRenderer {
	b := canvas.NewRectangle(color.RGBA{0, 0, 0, 0})
	b.StrokeWidth = 2
	b.StrokeColor = styles.Grey_300
	b.CornerRadius = 8
	c := container.NewStack(b,
		container.New(
			layout.NewCustomPaddedLayout(2, 2, 2, 2),
			item.Select,
		),
	)
	return widget.NewSimpleRenderer(c)
}
