package mywidgets

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type myWidget struct {
	widget.Label
}

func NewMyWidget(text string) *myWidget {
	v := &myWidget{}
	v.ExtendBaseWidget(v)
	v.Text = text
	return v
}

func (v *myWidget) Tapped(p *fyne.PointEvent) {
	log.Println(p.Position)
	log.Println(p.AbsolutePosition)
	log.Println("I have been tapped")
}

func (v *myWidget) TappedSecondary(_ *fyne.PointEvent) {
	log.Println("I have been tapped 2")
}
