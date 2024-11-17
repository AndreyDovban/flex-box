package widgets

import (
	"flexbox/styles"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var arrowIconRed = &fyne.StaticResource{
	StaticName:    "arrow-forward.svg",
	StaticContent: []byte("<svg xmlns=\"http://www.w3.org/2000/svg\" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\"><path d=\"M0 0h24v24H0z\" fill=\"none\"/><path fill=\"red\" d=\"M12 4l-1.41 1.41L16.17 11H4v2h12.17l-5.58 5.59L12 20l8-8z\"/></svg>\n"),
}

type ArrowLabel struct {
	widget.BaseWidget
	Icon  *widget.Icon
	Sel   *widget.Icon
	Label *widget.Label
	Fon   *canvas.Rectangle
	Bas   *canvas.Rectangle

	Selected bool
}

func NewArrowLabel(text string) *ArrowLabel {
	elem := &ArrowLabel{
		Icon:  widget.NewIcon(theme.NavigateNextIcon()),
		Sel:   widget.NewIcon(arrowIconRed),
		Label: widget.NewLabel(text),
		Fon:   canvas.NewRectangle(nil),
		Bas:   canvas.NewRectangle(nil),
	}
	elem.ExtendBaseWidget(elem)

	return elem
}

func (elem *ArrowLabel) CreateRenderer() fyne.WidgetRenderer {

	icon := elem.Icon
	icon.Resize(fyne.NewSize(22, 22))
	icon.Move(fyne.NewPos(20, 10))

	sel := elem.Sel
	sel.Resize(fyne.NewSize(22, 22))
	sel.Move(fyne.NewPos(20, 10))

	label := elem.Label
	label.Move(fyne.NewPos(52, 0))

	fon := elem.Fon
	fon.StrokeWidth = 1
	fon.StrokeColor = nil
	fon.CornerRadius = 8
	fon.Resize(fyne.NewSize(215, 42))

	bas := elem.Bas
	bas.Resize(fyne.NewSize(10, 42))
	bas.Move(fyne.NewPos(201, 0))

	c := container.NewWithoutLayout(
		icon,
		sel,
		label,
		fon,
		bas,
	)
	elem.Icon.Hide()
	elem.Sel.Hide()
	elem.Bas.Hide()

	return widget.NewSimpleRenderer(c)
}

func (elem *ArrowLabel) Cursor() desktop.Cursor {
	return desktop.PointerCursor
}

// MouseIn is called when a desktop pointer enters the widget
func (elem *ArrowLabel) MouseIn(*desktop.MouseEvent) {
	elem.Icon.Show()
	elem.Refresh()
}

// MouseMoved is called when a desktop pointer hovers over the widget
func (elem *ArrowLabel) MouseMoved(*desktop.MouseEvent) {
}

// MouseOut is called when a desktop pointer exits the widget
func (elem *ArrowLabel) MouseOut() {
	elem.Icon.Hide()
	elem.Refresh()
}

func (elem *ArrowLabel) SetSelect() {
	elem.Selected = true
	elem.Sel.Show()
	elem.Fon.StrokeColor = styles.Grey_600
	elem.Bas.Show()

	th := fyne.CurrentApp().Settings().Theme()
	h := th.Color(theme.ColorNameOverlayBackground, 0)
	elem.Bas.FillColor = h
	elem.Refresh()
}

func (elem *ArrowLabel) UnSelect() {
	elem.Selected = false
	elem.Sel.Hide()
	elem.Fon.StrokeColor = nil
	elem.Bas.Hide()

	th := fyne.CurrentApp().Settings().Theme()
	h := th.Color(theme.ColorNameOverlayBackground, 0)
	elem.Bas.FillColor = h
	elem.Refresh()
}
