package widgets

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type TappedText struct {
	widget.BaseWidget
	Alignment  fyne.TextAlign
	Title      *canvas.Text
	OnTapped   func() `json:"-"`
	Color      color.Color
	HoverColor color.Color

	TextSize float32
	hovered  bool
}

func NewTappedText(text string, function func()) *TappedText {

	elem := &TappedText{
		Title:    canvas.NewText(text, nil),
		OnTapped: function,
	}

	elem.Title.Alignment = fyne.TextAlignLeading
	elem.Alignment = fyne.TextAlignLeading
	elem.ExtendBaseWidget(elem)

	return elem
}

func (elem *TappedText) CreateRenderer() fyne.WidgetRenderer {
	c := container.NewWithoutLayout(elem.Title)

	return widget.NewSimpleRenderer(c)
}

func (elem *TappedText) Cursor() desktop.Cursor {
	return desktop.PointerCursor
}

func (elem *TappedText) Tapped(e *fyne.PointEvent) {
	elem.OnTapped()
}

func (elem *TappedText) TappedSecondary(_ *fyne.PointEvent) {}

func (elem *TappedText) MouseIn(e *desktop.MouseEvent) {
	elem.hovered = true
	th := elem.Theme()
	v := fyne.CurrentApp().Settings().ThemeVariant()
	elem.Title.Color = th.Color(theme.ColorNameHeaderBackground, v)
	elem.Refresh()
}

func (elem *TappedText) MouseOut() {
	elem.hovered = false
	th := elem.Theme()
	v := fyne.CurrentApp().Settings().ThemeVariant()
	elem.Title.Color = th.Color(theme.ColorNameHyperlink, v)
	elem.Refresh()

}

func (elem *TappedText) MouseMoved(e *desktop.MouseEvent) {
	oldHovered := elem.hovered
	if elem.hovered != oldHovered {
		elem.BaseWidget.Refresh()
	}
}

func (elem *TappedText) SetText(text string) {
	elem.Title.Text = text
	elem.Refresh()
}

func (elem *TappedText) SetSize(size float32) {
	elem.Title.TextSize = size
	elem.Refresh()
}

func (elem *TappedText) MinSize() fyne.Size {
	return elem.Title.MinSize()
}
