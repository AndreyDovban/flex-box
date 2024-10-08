package widgets

import (
	"image/color"

	"fyne.io/fyne/v2/canvas"
)

func Fon(background, strokeColor color.Color, cornerRarius float32) *canvas.Rectangle {
	b := canvas.NewRectangle(background)
	b.StrokeWidth = 1
	b.StrokeColor = strokeColor
	b.CornerRadius = cornerRarius
	return b
}
