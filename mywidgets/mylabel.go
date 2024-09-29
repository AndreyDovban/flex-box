package mywidgets

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// MyLabel widget is a label component with appropriate padding and layout.
type MyLabel struct {
	widget.Label
}

// NewLabel creates a new label widget with the set text content
func NewLabel(text string) *MyLabel {
	elem := &MyLabel{}
	elem.Text = text
	elem.ExtendBaseWidget(elem)

	return elem
}

func (elem *MyLabel) MinSize() fyne.Size {
	elem.ExtendBaseWidget(elem)
	return elem.BaseWidget.MinSize()
}

// func (elem *MyLabel) Tapped(_ *fyne.PointEvent) {
// 	elem.Text += "1"
// 	elem.Refresh()
// }

// func (elem *MyLabel) TappedSecondary(_ *fyne.PointEvent) {
// 	log.Println("I have been tapped 2")
// 	elem.Refresh()
// }
