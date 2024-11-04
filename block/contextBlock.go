package block

import (
	"fmt"
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func PrecessLong(flag *bool) {
	if *flag {
		return
	}
	*flag = true
	fmt.Println("start")

	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err.Error())
	}
	for {
		time.Sleep(time.Second)
		file.WriteString("test string\n")
		fmt.Println(*flag)
		if !*flag {
			return
		}

	}
}

/** Notify Block */
func ContextBlock() *fyne.Container {
	flag := false

	butStart := widget.NewButton("Start", func() {
		go PrecessLong(&flag)
	})
	butStart.Importance = widget.HighImportance

	butStop := widget.NewButton("Stop", func() {
		flag = false
	})
	butStop.Importance = widget.HighImportance

	textGrigContent := container.NewCenter(
		container.New(
			layout.NewCustomPaddedHBoxLayout(32),
			butStart,
			butStop,
		),
	)

	return textGrigContent
}
