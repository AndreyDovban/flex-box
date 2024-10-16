package widgets

import (
	"flexbox/styles"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var arrowIconRed = &fyne.StaticResource{
	StaticName:    "arrow-forward.svg",
	StaticContent: []byte("<svg xmlns=\"http://www.w3.org/2000/svg\" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\"><path d=\"M0 0h24v24H0z\" fill=\"none\"/><path fill=\"red\" d=\"M12 4l-1.41 1.41L16.17 11H4v2h12.17l-5.58 5.59L12 20l8-8z\"/></svg>\n"),
}

type ArrowLabel struct {
	widget.BaseWidget
	Icon  *widget.Icon
	Label *widget.Label
	F     fyne.CanvasObject

	Selected     bool
	SelectedIcon *widget.Icon
}

func NewArrowLabel(text string) *ArrowLabel {
	elem := &ArrowLabel{
		Label:        widget.NewLabel(text),
		Icon:         widget.NewIcon(theme.NavigateNextIcon()),
		F:            container.NewGridWrap(fyne.NewSize(24, 0)),
		SelectedIcon: widget.NewIcon(arrowIconRed),
	}
	elem.ExtendBaseWidget(elem)

	return elem
}

func (elem *ArrowLabel) CreateRenderer() fyne.WidgetRenderer {
	f := elem.F

	icon := elem.Icon
	icon.Resize(fyne.NewSize(24, 24))

	selIc := elem.SelectedIcon

	c := container.NewStack(
		Fon(color.Transparent, styles.Grey_400, 8),
		container.NewBorder(nil, nil,
			container.NewStack(
				f,
				icon,
				selIc,
			),
			nil,
			elem.Label,
		),
	)
	elem.Icon.Hide()
	selIc.Hide()

	return widget.NewSimpleRenderer(c)
}

func (elem *ArrowLabel) Cursor() desktop.Cursor {
	return desktop.PointerCursor
}

// MouseIn is called when a desktop pointer enters the widget
func (elem *ArrowLabel) MouseIn(*desktop.MouseEvent) {
	if !elem.Selected {
		elem.Icon.Show()
		elem.Refresh()
	}
}

// MouseMoved is called when a desktop pointer hovers over the widget
func (elem *ArrowLabel) MouseMoved(*desktop.MouseEvent) {
}

// MouseOut is called when a desktop pointer exits the widget
func (elem *ArrowLabel) MouseOut() {
	elem.Icon.Hide()
	elem.Refresh()
}

func (elem *ArrowLabel) SetSelect() {
	elem.Selected = true
	elem.SelectedIcon.Show()
	elem.Refresh()
}
