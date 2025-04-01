package block

import (
	"flexbox/widgets"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func submit(s string) {
	fmt.Println(s)
}

/** Custom wibget block*/
func CustomWidgetBlokc(colorTheme binding.Bool) *fyne.Container {

	entry := widget.NewEntry()
	entry.OnSubmitted = submit

	num := widgets.NewNumericalEntry()
	num.OnSubmitted = submit

	content := container.NewCenter(
		container.NewGridWrap(
			fyne.NewSize(300, 44),
			entry,
			num,
		),
	)

	return content
}
