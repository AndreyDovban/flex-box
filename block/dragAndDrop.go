package block

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

type DrawRect struct {
	widget.Label
}

func (e *DrawRect) Dragged(d *fyne.DragEvent) {
	e.Move(fyne.NewPos(d.AbsolutePosition.X, d.AbsolutePosition.Y))
	e.Refresh()
}

func (e *DrawRect) DragEnd() {
	fmt.Println("B")
}

func (e *DrawRect) Cursor() desktop.Cursor {
	return desktop.PointerCursor
}

/** Drag End Drop Block */
func DragEndDrop() *fyne.Container {

	// elem := canvas.NewRectangle(color.Black)
	// elem.Resize(fyne.NewSquareSize(100))

	elem := &DrawRect{}
	elem.Text = "Click"
	elem.Resize(fyne.NewSquareSize(100))

	dragendDropContent := container.NewWithoutLayout(
		elem,
	)

	// dragendDropContent := container.NewCenter(
	// 	container.NewStack(
	// 		widgets.Fon(color.Transparent, color.White, 8),
	// 		container.New(
	// 			layout.NewGridWrapLayout(fyne.NewSize(400, 400)),
	// 			container.New(layout.NewCustomPaddedLayout(16, 16, 16, 16),
	// 				container.New(layout.NewCustomPaddedVBoxLayout(24),
	// 					elem,
	// 				),
	// 			),
	// 		),
	// 	),
	// )

	return dragendDropContent
}
