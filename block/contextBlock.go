package block

import (
	"fmt"
	"log"
	"os/exec"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func f(stop chan bool) {
	cmd := exec.Command("tilix")
	log.Printf("Running command and waiting for it to finish...")
	err := cmd.Start()
	log.Printf("Command finished with error: %v", err)

	for {
		select {
		default:
			// Выполнение работы в горутине
			fmt.Println("Работаю...")
			time.Sleep(1 * time.Second)
		case <-stop:
			// Получен сигнал об остановке
			fmt.Println("Остановлен")
			cmd.Wait()
			return
		}
	}
}

/** Notify Block */
func ContextBlock() *fyne.Container {
	stop := make(chan bool)

	butStart := widget.NewButton("Start", func() {
		go f(stop)
	})
	butStart.Importance = widget.HighImportance

	butStop := widget.NewButton("Stop", func() {
		stop <- true
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
