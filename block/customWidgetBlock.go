package block

import (
	"flexbox/widgets"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

/** Custom wibget block*/
func CustomWidgetBlokc() *fyne.Container {

	elem := widgets.NewTappedText("text", func() {
		fmt.Println("work")
	})

	elem.SetSize(32)
	elem.SetText(";sd;d;")

	fmt.Println(fyne.CurrentApp().Settings().ThemeVariant())

	content := container.NewCenter(
		container.New(
			layout.NewCustomPaddedHBoxLayout(32),
			elem,
		),
	)

	return content
}
