package block

import (
	"log"
	"os/exec"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

/** Notify Block */
func NotifyBlock() *fyne.Container {
	but := widget.NewButton("send notify", func() {
		cmd := exec.Command("notify-send", "hello")
		log.Printf("Running command and waiting for it to finish...")
		err := cmd.Run()
		log.Printf("Command finished with error: %v", err)
	})
	but.Importance = widget.HighImportance

	textGrigContent := container.NewCenter(
		but,
	)

	return textGrigContent
}
