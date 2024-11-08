package widgets

import (
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type Link struct {
	widget.BaseWidget
	Alignment fyne.TextAlign
	Title     *canvas.Text
	Url       *url.URL

	TextSize float32
	hovered  bool
}

func NewLink(text, url_string string) *Link {
	t, err := url.Parse(url_string)
	if err != nil {
		t = nil
	}
	elem := &Link{
		Title: canvas.NewText(text, nil),
		Url:   t,
	}

	return elem
}

func (elem *Link) CreateRenderer() fyne.WidgetRenderer {
	elem.ExtendBaseWidget(elem)

	th := elem.Theme()
	h := th.Color(theme.ColorNameHyperlink, 0)

	elem.Title.Color = h

	c := container.NewWithoutLayout(elem.Title)

	return widget.NewSimpleRenderer(c)
}

func (elem *Link) Cursor() desktop.Cursor {
	return desktop.PointerCursor
}

func (elem *Link) Tapped(e *fyne.PointEvent) {
	if elem.Url != nil {
		err := fyne.CurrentApp().OpenURL(elem.Url)
		if err != nil {
			fyne.LogError("Failed to open url", err)
		}
	}
}

func (elem *Link) TappedSecondary(_ *fyne.PointEvent) {}

func (elem *Link) MouseIn(e *desktop.MouseEvent) {
	elem.hovered = true
	th := fyne.CurrentApp().Settings().Theme()
	h := th.Color(theme.ColorNameHeaderBackground, 0)
	elem.Title.Color = h
	elem.Refresh()
}

func (elem *Link) MouseOut() {
	elem.hovered = false
	th := fyne.CurrentApp().Settings().Theme()
	h := th.Color(theme.ColorNameHyperlink, 0)
	elem.Title.Color = h
	elem.Refresh()

}

func (elem *Link) MouseMoved(e *desktop.MouseEvent) {
	oldHovered := elem.hovered
	if elem.hovered != oldHovered {
		elem.BaseWidget.Refresh()
	}
}

func (elem *Link) Update() {
	th := fyne.CurrentApp().Settings().Theme()
	h := th.Color(theme.ColorNameHyperlink, 0)
	elem.Title.Color = h
	elem.Title.Refresh()
}

func (elem *Link) SetText(text string) {
	elem.Title.Text = text
	elem.Refresh()
}

func (elem *Link) SetSize(size float32) {
	elem.Title.TextSize = size
	elem.Refresh()
}

func (elem *Link) MinSize() fyne.Size {
	return elem.Title.MinSize()
}
