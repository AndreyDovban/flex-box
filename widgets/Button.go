package widgets

import (
	"flexbox/styles"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

type Button struct {
	widget.DisableableWidget
	Text       *canvas.Text
	Background *canvas.Rectangle

	OnTapped func() `json:"-"`

	focused bool
}

func NewButton(text string, tapped func()) *Button {
	elem := &Button{
		Text:     canvas.NewText(text, styles.White),
		OnTapped: tapped,
	}
	elem.ExtendBaseWidget(elem)
	return elem
}

func (elem *Button) CreateRenderer() fyne.WidgetRenderer {

	text := elem.Text
	text.TextSize = 16
	text.TextStyle = fyne.TextStyle{
		Bold: true,
	}

	but := container.NewCenter(text)

	if elem.Disabled() {
		elem.Background = canvas.NewRectangle(styles.Grey_300)
	} else {
		elem.Background = canvas.NewRectangle(styles.Red)
	}
	elem.Background.CornerRadius = 8

	contain := container.NewStack(
		elem.Background,
		but,
	)

	return widget.NewSimpleRenderer(contain)
}

func (elem *Button) Tapped(_ *fyne.PointEvent) {
	if elem.Disabled() {
		return
	}
	elem.OnTapped()
}

func (elem *Button) TappedSecondary(_ *fyne.PointEvent) {
}

func (elem *Button) FocusGained() {
	elem.focused = true
	elem.Background.StrokeWidth = 1
	elem.Background.StrokeColor = styles.Grey_700
	elem.Refresh()
}

func (elem *Button) FocusLost() {
	elem.focused = false
	elem.Background.StrokeWidth = 0
	elem.Background.StrokeColor = color.Transparent
	elem.Refresh()
}

func (b *Button) TypedRune(rune) {
}

func (elem *Button) TypedKey(ev *fyne.KeyEvent) {
	if ev.Name == fyne.KeySpace || ev.Name == fyne.KeyEnter || ev.Name == fyne.KeyReturn {
		elem.OnTapped()
	}
}

func (elem *Button) MouseDown(*desktop.MouseEvent) {
	if elem.Disabled() {
		return
	}
	elem.Background.FillColor = color.Black
	elem.Refresh()
}

func (elem *Button) MouseUp(*desktop.MouseEvent) {
	if elem.Disabled() {
		return
	}
	elem.Background.FillColor = styles.Red_Dark
	elem.Refresh()
}

func (elem *Button) Cursor() desktop.Cursor {
	if elem.Disabled() {
		return desktop.DefaultCursor
	}
	return desktop.PointerCursor
}

func (elem *Button) MouseIn(*desktop.MouseEvent) {
	if elem.Disabled() {
		return
	}
	elem.Background.FillColor = styles.Red_Dark
	elem.Refresh()
}

func (elem *Button) MouseMoved(*desktop.MouseEvent) {
}

func (elem *Button) MouseOut() {
	if elem.Disabled() {
		return
	}
	elem.Background.FillColor = styles.Red
	elem.Refresh()
}
