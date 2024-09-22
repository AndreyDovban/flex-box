package mytheme

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type MyTheme struct{}

func (MyTheme) Color(c fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	switch c {
	case theme.ColorNameBackground:
		return color.RGBA{41, 41, 41, 255}
	case fyne.ThemeColorName(theme.IconNameNavigateNext):
		return color.RGBA{60, 60, 63, 255}
	case theme.ColorNamePrimary:
		return color.RGBA{225, 225, 0, 255}
	case theme.ColorNameSelection:
		return color.RGBA{100, 100, 110, 255}
	case theme.ColorNameShadow:
		return color.RGBA{25, 25, 30, 255}
	default:
		return theme.DefaultTheme().Color(c, v)
	}
}

func (MyTheme) Font(s fyne.TextStyle) fyne.Resource {
	if s.Monospace {
		return theme.DefaultTheme().Font(s)
	}
	if s.Bold {
		if s.Italic {
			return theme.DefaultTheme().Font(s)
		}
		return theme.DefaultTheme().Font(s)
	}
	if s.Italic {
		return theme.DefaultTheme().Font(s)
	}
	return theme.DefaultTheme().Font(s)
}

func (MyTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(n)
}

func (MyTheme) Size(s fyne.ThemeSizeName) float32 {
	switch s {
	case theme.SizeNameCaptionText:
		return 10
	case theme.SizeNameInlineIcon:
		return 22
	case theme.SizeNamePadding:
		return 2
	case theme.SizeNameInnerPadding:
		return 4
	case theme.SizeNameScrollBar:
		return 8
	case theme.SizeNameScrollBarSmall:
		return 3
	case theme.SizeNameSeparatorThickness:
		return 1
	case theme.SizeNameText:
		return 13
	case theme.SizeNameInputBorder:
		return 2
	case theme.SizeNameInputRadius:
		return 4
	default:
		return theme.DefaultTheme().Size(s)
	}
}
