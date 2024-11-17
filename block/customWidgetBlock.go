package block

import (
	"flexbox/styles"
	"flexbox/widgets"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

/** Custom wibget block*/
func CustomWidgetBlokc(colorTheme binding.Bool) *fyne.Container {

	// light := &styles.Light{}
	// dark := &styles.Dark{}

	but := widgets.NewButton("click", func() {
		v, _ := colorTheme.Get()
		colorTheme.Set(!v)

	})
	but.Resize(fyne.NewSize(200, 100))

	elem := widget.NewLabel("hjghjgj gjh gj hjg hjg j gjhghjgjhg jg hgjgjgh hhgjgjh hh j gghghghghgh hg hgh gh ghg h gh jhgjhg jgj gjg jgjg hjg ")
	elem.Wrapping = fyne.TextWrapWord
	ttt := container.NewThemeOverride(
		elem,
		nil,
	)

	// th := fyne.CurrentApp().Settings().Theme()
	// v := fyne.CurrentApp().Settings().ThemeVariant()
	// elem.SetColor(th.Color(theme.ColorNameHyperlink, v))

	content := container.NewCenter(
		container.New(
			layout.NewCustomPaddedVBoxLayout(0),
			but,
			container.NewGridWrap(
				fyne.NewSize(200, elem.Size().Height),
				container.NewStack(
					widgets.Fon(styles.Grey_800, styles.Grey_600, 8),
					ttt,
				),
			),
		),
	)

	colorTheme.AddListener(binding.NewDataListener(func() {
		v, _ := colorTheme.Get()
		if v {
			ttt.Theme = &styles.Light2{}
		} else {
			ttt.Theme = &styles.Dark2{}
		}
	}))

	return content
}
