package main

import (
	"flexbox/block"
	"flexbox/styles"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/driver/desktop"
)

func main() {
	app := app.New()
	mainWindow := app.NewWindow("Learn")

	colorTheme := binding.NewBool()
	colorTheme.Set(false)
	light := &styles.Light{}
	// dark := &styles.Dark{}

	/** Shortcat example */

	ctrlTab := &desktop.CustomShortcut{KeyName: fyne.KeyTab, Modifier: fyne.KeyModifierControl}

	mainWindow.Canvas().AddShortcut(ctrlTab, func(shortcut fyne.Shortcut) {
		log.Println("We tapped Ctrl+Tab")
	})

	/** App Tabs block */

	tab := container.NewAppTabs(
		container.NewTabItem("DragEndDrop", block.DragEndDrop()),
		container.NewTabItem("TextGrig", block.TextGridBlock()),
		container.NewTabItem("Theme", block.CustomThemeBlock(colorTheme)),
		container.NewTabItem("Widget", block.CustomWidgetBlokc()),
		container.NewTabItem("Animation", block.AnimationBlock()),
		container.NewTabItem("Tree", block.TreeBlock()),
		container.NewTabItem("List", block.ListBlock()),
		container.NewTabItem("Flex", block.FlexBlock()),
		container.NewTabItem("Notify", block.NotifyBlock()),
		container.NewTabItem("Context", block.ContextBlock()),
	)
	tab.SelectIndex(9)

	mainWindow.SetContent(tab)
	colorTheme.AddListener(binding.NewDataListener(func() {
		v, _ := colorTheme.Get()
		if v {
			app.Settings().SetTheme(light)
		} else {
			app.Settings().SetTheme(&styles.Dark{})
		}
		// testMenuItem.FillColor = styles.ColorNameForeground
	}))

	mainWindow.CenterOnScreen()
	mainWindow.Resize(fyne.NewSize(1000, 600))
	mainWindow.Show()

	app.Run()
}
