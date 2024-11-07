package widgets

import (
	"fmt"
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
	elem.ExtendBaseWidget(elem)

	return elem
}

func (elem *TappedText) CreateRenderer() fyne.WidgetRenderer {
	c := container.NewCenter(elem.Title)

	return widget.NewSimpleRenderer(c)
}

func (elem *TappedText) Cursor() desktop.Cursor {
	return desktop.PointerCursor
}

func (elem *TappedText) Tapped(e *fyne.PointEvent) {
	fmt.Println(elem.MinSize())
	elem.OnTapped()
}

func (elem *TappedText) TappedSecondary(_ *fyne.PointEvent) {}

func (elem *TappedText) MouseIn(e *desktop.MouseEvent) {
	elem.hovered = true
	th := elem.Theme()
	v := fyne.CurrentApp().Settings().ThemeVariant()
	fmt.Println(th.Color("hyperlink", 0))
	elem.Title.Color = th.Color(theme.ColorNameFocus, v)
	elem.Refresh()
}

func (elem *TappedText) MouseOut() {
	elem.hovered = false
	elem.Title.Color = elem.Color
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
