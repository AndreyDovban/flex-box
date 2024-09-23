package mywidgets

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type myWidget struct {
	widget.Button
}

func NewMyWidget() *myWidget {
	v := &myWidget{}
	v.ExtendBaseWidget(v)
	return v
}

func (m *myWidget) CreateRenderer() fyne.WidgetRenderer {

	but := widget.NewButton(m.Text, nil)

	return widget.NewSimpleRenderer(but)
}

func (v *myWidget) Tapped(p *fyne.PointEvent) {
	log.Println(p.Position)
	log.Println(p.AbsolutePosition)
	log.Println("I have been tapped")
}

func (v *myWidget) TappedSecondary(_ *fyne.PointEvent) {
	log.Println("I have been tapped 2")
}
