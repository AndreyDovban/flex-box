package block

import (
	"flexbox/widgets"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

/** Custom wibget block*/
func CustomWidgetBlokc() *fyne.Container {

	customWidget := widgets.NewLabel("Example cloud.google.com/go v0.44.1/go.mod h1:iSa0Kza dfgd dfg ddfg df  dfgdfgdfg  dfgdfgdfg dfgdfgdfgd dfgdfgdfg dfgdfgdf")
	// customWidget.Title.Resize(fyne.NewSize(200, 400))
	customWidget.Resize(fyne.NewSize(200, 400))

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
