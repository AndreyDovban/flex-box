package mywidgets

import (
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Vid struct {
	widget.Label
}

func NewVid(text string) *Vid {
	v := &Vid{}
	v.ExtendBaseWidget(v)
	v.Text = text
	return v
}

func (m *Vid) CreateRenderer() fyne.WidgetRenderer {
	// background := canvas.NewRectangle(color.Transparent)
	// background.StrokeWidth = 1
	// background.StrokeColor = color.RGBA{200, 200, 200, 255}

	text := widget.NewLabel(m.Text)
	// text.Alignment = fyne.TextAlignCenter

	log.Println(text.Size())

	line := canvas.NewLine(color.RGBA{60, 60, 70, 255})
	line.Position1 = fyne.NewPos(0, 14)
	line.Position2 = fyne.NewPos(-30, 14)

	c := container.NewWithoutLayout(
		line,
		text,
	)

	return widget.NewSimpleRenderer(c)
}

func (v *Vid) Tapped(p *fyne.PointEvent) {
	log.Println(p.Position)
	log.Println(p.AbsolutePosition)
	log.Println("I have been tapped")
}

func (v *Vid) TappedSecondary(_ *fyne.PointEvent) {
	log.Println("I have been tapped 2")
}
