package block

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func AnimationBlock() *fyne.Container {
	/** Animation Block */

	elem := canvas.NewRectangle(color.Black)
	elem.Resize(fyne.NewSize(50, 50))

	animationContent := container.NewStack(
		container.NewWithoutLayout(
			elem,
		),
	)

	red := color.NRGBA{R: 0xff, A: 0xff}
	blue := color.NRGBA{G: 0x99, B: 0xff, A: 0xff}
	anima := canvas.NewColorRGBAAnimation(red, blue, time.Second*2, func(c color.Color) {
		elem.FillColor = c
		canvas.Refresh(elem)
	})

	anima.AutoReverse = true
	anima.RepeatCount = 30
	anima.Curve = fyne.AnimationEaseIn
	anima.Start()

	move := canvas.NewPositionAnimation(fyne.NewPos(0, 0), fyne.NewPos(200, 0), time.Second, elem.Move)
	move.AutoReverse = true
	move.RepeatCount = 30
	move.Start()

	return animationContent
}
