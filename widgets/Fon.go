package widgets

import (
	"image/color"

	"fyne.io/fyne/v2/canvas"
)

func Fon(background color.Color, cornerRadius float32, strokeColor color.Color) *canvas.Rectangle {
	elem := canvas.NewRectangle(background)
	elem.CornerRadius = cornerRadius
	elem.StrokeWidth = 0
	elem.StrokeColor = strokeColor

	return elem
}
