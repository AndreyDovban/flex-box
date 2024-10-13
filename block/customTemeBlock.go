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

	testMenuItem := widgets.Fon(styles.ColorNameForeground, color.Black, styles.Radius)
	testMenuItem.SetMinSize(fyne.NewSquareSize(200))

	testLabel := canvas.NewText("GENERATION", styles.ColorNameForeground)
	testLabel.TextSize = 24
	testLabel.TextStyle.Bold = true

	changeThemeBut := widget.NewButton("click", func() {
		v, _ := colorTheme.Get()
		colorTheme.Set(!v)

	})

	customThemeContent := container.NewCenter(
		container.New(
			layout.NewGridWrapLayout(fyne.NewSize(400, 600)),
			container.New(layout.NewCustomPaddedVBoxLayout(24),
				testLabel,
				testMenuItem,
				changeThemeBut,
			),
		))

	return customThemeContent
}
