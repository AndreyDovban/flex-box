package widgets

import (
	"flexbox/styles"
	"image/color"
	"log"
	"time"

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
	hovered bool

	tapAnim *fyne.Animation
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

	elem.Background = canvas.NewRectangle(styles.Red)
	elem.Background.CornerRadius = 8

	contain := container.NewStack(
		elem.Background,
		but,
	)

	elem.tapAnim = newButtonTapAnimation(elem)
	elem.tapAnim.Curve = fyne.AnimationEaseOut

	return widget.NewSimpleRenderer(contain)
}

func (elem *Button) Tapped(_ *fyne.PointEvent) {
	if !elem.Disabled() {
		elem.tapAnimation()
		elem.CRefresh()
		elem.OnTapped()
	}
}

func (elem *Button) TappedSecondary(_ *fyne.PointEvent) {
}

func (elem *Button) FocusGained() {
	elem.focused = true
	elem.CRefresh()
}

func (elem *Button) FocusLost() {
	elem.focused = false
	elem.CRefresh()
}

func (elem *Button) TypedRune(rune) {
}

func (elem *Button) TypedKey(ev *fyne.KeyEvent) {
	if ev.Name == fyne.KeySpace || ev.Name == fyne.KeyEnter || ev.Name == fyne.KeyReturn {
		if !elem.Disabled() {
			elem.tapAnimation()
			elem.Refresh()
			elem.OnTapped()
		}
	}
}

func (elem *Button) MouseDown(*desktop.MouseEvent) {
	elem.CRefresh()
}

func (elem *Button) MouseUp(*desktop.MouseEvent) {
	elem.CRefresh()
}

func (elem *Button) Cursor() desktop.Cursor {
	if elem.Disabled() {
		return desktop.DefaultCursor
	}
	return desktop.PointerCursor
}

func (elem *Button) MouseIn(*desktop.MouseEvent) {
	elem.hovered = true
	elem.CRefresh()
}

func (elem *Button) MouseMoved(*desktop.MouseEvent) {
}

func (elem *Button) MouseOut() {
	elem.hovered = false
	elem.CRefresh()
}

func (elem *Button) CRefresh() {
	log.Println("Refresh")

	if elem.Disabled() {
		elem.Background.FillColor = styles.Grey_300
		elem.Refresh()
		return
	}

	elem.Background.StrokeWidth = 0
	elem.Background.StrokeColor = color.Transparent
	elem.Background.FillColor = styles.Red

	if elem.focused {
		elem.Background.StrokeWidth = 2
		elem.Background.StrokeColor = styles.Grey_700
	}

	if elem.hovered {
		elem.Background.FillColor = styles.Red_Dark
	}

	elem.Refresh()
}

func (elem *Button) tapAnimation() {
	if elem.tapAnim == nil {
		return
	}
	elem.tapAnim.Stop()

	if fyne.CurrentApp().Settings().ShowAnimations() {
		elem.tapAnim.Start()
	}
}

func newButtonTapAnimation(elem *Button) *fyne.Animation {
	red := color.NRGBA{R: 0xff, A: 0xff}
	blue := color.RGBA{0, 0, 0, 255}
	anima := canvas.NewColorRGBAAnimation(red, blue, time.Millisecond*100, func(c color.Color) {
		elem.Background.FillColor = c
		canvas.Refresh(elem)
	})

	anima.AutoReverse = true
	anima.RepeatCount = 0
	anima.Curve = fyne.AnimationEaseIn
	return anima
}
