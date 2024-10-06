package mywidgets

import (
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

type Button struct {
	widget.BaseWidget
	Title      *widget.Label
	background *canvas.Rectangle
}

func NewButton() *Button {
	elem := &Button{Title: widget.NewLabel("click")}
	elem.ExtendBaseWidget(elem)
	return elem
}

func (elem *Button) CreateRenderer() fyne.WidgetRenderer {

	but := container.NewCenter(elem.Title)

	elem.background = canvas.NewRectangle(color.RGBA{255, 0, 0, 255})
	elem.background.StrokeWidth = 1
	elem.background.StrokeColor = color.RGBA{0, 0, 0, 255}
	elem.background.CornerRadius = 8
	// elem.background.Hide()
	elem.Resize(fyne.NewSize(200, 35))

	contain := container.NewStack(
		elem.background,
		but,
	)

	return widget.NewSimpleRenderer(contain)
}

func (elem *Button) Tapped(_ *fyne.PointEvent) {
	log.Println("I have been tapped")
}

func (elem *Button) TappedSecondary(_ *fyne.PointEvent) {
}

func (elem *Button) Cursor() desktop.Cursor {
	return desktop.PointerCursor
}

func (elem *Button) MouseIn(*desktop.MouseEvent) {
	elem.background.FillColor = color.RGBA{140, 0, 0, 255}
	elem.Refresh()
	log.Println("IN")
}

func (elem *Button) MouseMoved(*desktop.MouseEvent) {
	log.Println("MOVE")
}

func (elem *Button) MouseOut() {
	elem.background.FillColor = color.RGBA{255, 0, 0, 255}
	elem.Refresh()
	log.Println("OUT")
}
