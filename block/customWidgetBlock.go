package block

import (
	"flexbox/widgets"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
)

/** Custom wibget block*/
func CustomWidgetBlokc(colorTheme binding.Bool) *fyne.Container {

	// light := &styles.Light{}
	// dark := &styles.Dark{}

	elem := widgets.NewTappedText("text", func() {
		fmt.Println("work")
	}, "")

	elem.SetSize(32)
	// th := fyne.CurrentApp().Settings().Theme()
	// v := fyne.CurrentApp().Settings().ThemeVariant()
	// elem.SetColor(th.Color(theme.ColorNameHyperlink, v))

	content := container.NewCenter(
		container.New(
			layout.NewCustomPaddedHBoxLayout(32),
			elem,
		),
	)

	colorTheme.AddListener(binding.NewDataListener(func() {
		elem.Update()
	}))

	return content
}
