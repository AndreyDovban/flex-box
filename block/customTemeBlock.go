package block

import (
	"flexbox/styles"
	"flexbox/widgets"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

/** Custom Theme block */
func CustomThemeBlock(colorTheme binding.Bool) *fyne.Container {

	testLabel := canvas.NewText("GENERATION", nil)
	testLabel.TextSize = 24
	testLabel.TextStyle.Bold = true

	testRectangle := widgets.Fon(nil, color.Black, styles.Radius)
	testRectangle.SetMinSize(fyne.NewSquareSize(200))

	changeThemeBut := widget.NewButton("click", func() {
		v, _ := colorTheme.Get()
		colorTheme.Set(!v)

	})

	customThemeContent := container.NewCenter(
		container.New(
			layout.NewGridWrapLayout(fyne.NewSize(400, 600)),
			container.New(layout.NewCustomPaddedVBoxLayout(24),
				testLabel,
				testRectangle,
				changeThemeBut,
			),
		))

	return customThemeContent
}
