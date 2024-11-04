package block

import (
	"fmt"
	"log"
	"os/exec"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

/** Notify Block */
func ContextBlock() *fyne.Container {

	var cmd *exec.Cmd
	pid := 0

	butStart := widget.NewButton("Start", func() {
		if pid == 0 {
			cmd = exec.Command("tilix")
			log.Printf("Running command and waiting for it to finish...")
			err := cmd.Start()
			log.Printf("Command finished with error: %v", err)
			pid = cmd.Process.Pid
		}
	})
	butStart.Importance = widget.HighImportance

	butStop := widget.NewButton("Stop", func() {
		if pid != 0 {
			fmt.Println(pid)
			cmd.Process.Kill()
			pid = 0
		}

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
