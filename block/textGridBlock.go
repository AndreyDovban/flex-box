package block

import (
	"flexbox/styles"
	"flexbox/widgets"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

/** Text Grig */
func TextGridBlock() *fyne.Container {

	loremIpsum := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque quis consectetur nisi. Suspendisse id interdum felis. Sed egestas eget tellus eu pharetra. Praesent pulvinar sed massa id placerat. Etiam sem libero."

	textGrid := widget.NewTextGridFromString("Пароль *")
	start := &widget.CustomTextGridStyle{FGColor: styles.Red}
	textGrid.Rows[0].Cells[7].Style = start

	labelTest := widget.NewLabel(loremIpsum)
	labelTest.Wrapping = fyne.TextWrapWord

	labelTest2 := widget.NewLabel(loremIpsum)
	labelTest2.Wrapping = fyne.TextWrapWord

	textGrigContent := container.NewCenter(
		container.NewStack(
			widgets.Fon(color.Transparent, color.White, 8),
			container.New(
				layout.NewGridWrapLayout(fyne.NewSize(400, 400)),
				container.New(layout.NewCustomPaddedVBoxLayout(24),
					textGrid,
					container.NewThemeOverride(labelTest, &styles.Dark2{}),
					labelTest2,
				),
			),
		),
	)

	return textGrigContent
}
