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
	Alignment fyne.TextAlign
	Title     *canvas.Text
	OnTapped  func() `json:"-"`

	TextSize float32
	hovered  bool
}

func NewTappedText(text string, function func()) *TappedText {

	elem := &TappedText{
		Title:    canvas.NewText(text, nil),
		OnTapped: function,
	}

	elem.Refresh()

	return elem
}

func (elem *TappedText) CreateRenderer() fyne.WidgetRenderer {
	elem.ExtendBaseWidget(elem)

	th := elem.Theme()
	h := th.Color(theme.ColorNameHyperlink, 0)

	elem.Title.Color = h

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
	th := fyne.CurrentApp().Settings().Theme()
	h := th.Color(theme.ColorNameHeaderBackground, 0)
	elem.Title.Color = h
	elem.Refresh()
}

func (elem *TappedText) MouseOut() {
	elem.hovered = false
	th := fyne.CurrentApp().Settings().Theme()
	h := th.Color(theme.ColorNameHyperlink, 0)
	elem.Title.Color = h

	elem.Refresh()
}

func (elem *TappedText) Update() {
	th := fyne.CurrentApp().Settings().Theme()
	h := th.Color(theme.ColorNameHyperlink, 0)
	elem.Title.Color = h
	elem.Title.Refresh()
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

func (elem *TappedText) SetColor(c color.Color) {
	elem.Title.Color = c
	elem.Refresh()
}

func (elem *TappedText) SetSize(size float32) {
	elem.Title.TextSize = size
	elem.Refresh()
}

func (elem *TappedText) MinSize() fyne.Size {
	return elem.Title.MinSize()
}
