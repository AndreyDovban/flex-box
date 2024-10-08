package widgets

import (
	"flexbox/styles"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/dialog"
)

func AuthForm(w fyne.Window) *dialog.CustomDialog {

	content := canvas.NewRectangle(styles.Grey_200)
	content.CornerRadius = styles.Radius
	content.Resize(fyne.NewSize(400, 500))

	elem := dialog.NewCustom(
		"Titme",
		"disming",
		content,
		w,
	)

	return elem
}
