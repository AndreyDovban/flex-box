package widgets

import (
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

type Link struct {
	widget.Hyperlink
}

func NewLink(text string) *Link {
	elem := &Link{}
	elem.ExtendBaseWidget(elem)
	elem.Text = text

	return elem
}

func (elem *Link) Cursor() desktop.Cursor {
	return desktop.PointerCursor
}
