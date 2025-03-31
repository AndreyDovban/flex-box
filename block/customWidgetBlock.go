package block

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

/** Custom wibget block*/
func CustomWidgetBlokc(colorTheme binding.Bool) *fyne.Container {

	entry := widget.NewEntry()
	entry.OnSubmitted = func(s string) {
		fmt.Println(s)
	}

	content := container.NewCenter(
		container.NewGridWrap(
			fyne.NewSize(300, 44),
			entry,
		),
	)

	return content
}
