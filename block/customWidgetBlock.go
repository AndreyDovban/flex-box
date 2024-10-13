package block

import (
	"flexbox/widgets"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

/** Custom wibget block*/
func CustomWidgetBlokc() *fyne.Container {

	customWidget := widget.NewEntry()

	customWidgetContent := container.NewCenter(
		container.NewStack(
			widgets.Fon(color.Transparent, color.White, 8),
			container.New(
				layout.NewGridWrapLayout(fyne.NewSize(400, 400)),
				container.New(layout.NewCustomPaddedLayout(16, 16, 16, 16),
					container.New(layout.NewCustomPaddedVBoxLayout(24),
						customWidget,
					),
				),
			),
		),
	)

	return customWidgetContent
}
