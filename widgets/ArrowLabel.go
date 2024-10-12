package widgets

import (
	"flexbox/styles"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type List struct {
	widget.BaseWidget
	List []*ArrowLabel
}

// func NewList([]string) *List {

// }

type ArrowLabel struct {
	widget.BaseWidget
	Icon *widget.Icon
	Text string
	F    fyne.CanvasObject

	hovered bool

	selected     bool
	selectedIcon *widget.Icon

	OnTapped func() `json:"-"`
}

func NewArrowLabel(text string, tapped func()) *ArrowLabel {
	elem := &ArrowLabel{
		Text:         text,
		Icon:         widget.NewIcon(theme.NavigateNextIcon()),
		OnTapped:     tapped,
		F:            container.NewGridWrap(fyne.NewSize(24, 24)),
		selectedIcon: widget.NewIcon(theme.HistoryIcon()),
	}
	elem.ExtendBaseWidget(elem)

	return elem
}

func (elem *ArrowLabel) CreateRenderer() fyne.WidgetRenderer {
	f := elem.F

	icon := elem.Icon
	icon.Resize(fyne.NewSize(24, 24))

	selIc := elem.selectedIcon

	c := container.NewStack(
		Fon(color.Transparent, styles.Grey_400, 8),
		container.NewBorder(nil, nil,
			container.NewStack(
				f,
				icon,
				selIc,
			),
			nil,
			widget.NewLabel(elem.Text),
		),
	)
	elem.Icon.Hide()
	selIc.Hide()

	return widget.NewSimpleRenderer(c)
}

func (elem *ArrowLabel) Cursor() desktop.Cursor {
	return desktop.PointerCursor
}

// MouseIn is called when a desktop pointer enters the widget
func (elem *ArrowLabel) MouseIn(*desktop.MouseEvent) {
	elem.hovered = true
	if !elem.selected {
		elem.Icon.Show()
		elem.Refresh()
	}
}

// MouseMoved is called when a desktop pointer hovers over the widget
func (elem *ArrowLabel) MouseMoved(*desktop.MouseEvent) {
}

// MouseOut is called when a desktop pointer exits the widget
func (elem *ArrowLabel) MouseOut() {
	elem.hovered = false
	elem.Icon.Hide()
	elem.Refresh()
}

func (elem *ArrowLabel) Tapped(*fyne.PointEvent) {
	elem.Refresh()
	if onTapped := elem.OnTapped; onTapped != nil {
		elem.selected = true
		elem.selectedIcon.Show()
		onTapped()
		elem.Refresh()
	}
}
