package styles

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type Light struct{}

func (Light) Color(c fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	switch c {
	case theme.ColorNameDisabled:
		return color.RGBA{255, 255, 255, 100}
	case theme.ColorNameDisabledButton:
		return color.RGBA{255, 255, 255, 100}

	case theme.ColorNameBackground:
		return Grey_200
	case theme.ColorNameButton:
		return Grey_300
	case theme.ColorNameForeground:
		return Grey_800
	// case theme.ColorNameForegroundOnPrimary:
	// 	return Red
	// case theme.ColorNameHeaderBackground:
	// 	return Red
	case theme.ColorNameInputBorder:
		return Grey_300
	// case theme.ColorNameShadow:
	// 	return Grey_400
	case theme.ColorNameInputBackground:
		return White
	case theme.ColorNameMenuBackground:
		return Grey_100
	// case theme.ColorNameSeparator:
	// 	return Red
	case theme.ColorNameHover:
		return Grey_300
	// case theme.ColorNameFocus:
	// 	return Red
	// case theme.ColorNamePrimary:
	// 	return Red
	// case theme.ColorNameSuccess:
	// 	return color.RGBA{0, 0, 255, 255}
	case theme.ColorNameSelection:
		return White
	case theme.ColorNameScrollBar:
		return Grey_600
	case theme.ColorNameHeaderBackground:
		return Red
	case theme.ColorNamePressed:
		return Red
	case theme.ColorNameOverlayBackground:
		return Grey_100

	default:
		return theme.DefaultTheme().Color(c, v)
	}
}

func (Light) Font(s fyne.TextStyle) fyne.Resource {
	// if s.Monospace {
	// 	return theme.DefaultTheme().Font(s)
	// }
	// if s.Bold {
	// 	if s.Italic {
	// 		return theme.DefaultTheme().Font(s)
	// 	}
	// 	return theme.DefaultTheme().Font(s)
	// }
	// if s.Italic {
	// 	return theme.DefaultTheme().Font(s)
	// }
	return theme.DefaultTheme().Font(s)
}

func (Light) Icon(n fyne.ThemeIconName) fyne.Resource {
	switch n {
	// case theme.IconNameSettings:
	// 	return static.ResSettingsSvg
	}
	return theme.DefaultTheme().Icon(n)
}

func (Light) Size(s fyne.ThemeSizeName) float32 {
	switch s {
	case theme.SizeNameCaptionText:
		return 10
	case theme.SizeNameInlineIcon:
		return 24
	case theme.SizeNamePadding:
		return 0
	case theme.SizeNameInnerPadding:
		return 10
	case theme.SizeNameScrollBar:
		return 10
	case theme.SizeNameScrollBarSmall:
		return 6
	case theme.SizeNameSeparatorThickness:
		return 1
	case theme.SizeNameText:
		return 16
	case theme.SizeNameInputBorder:
		return 1
	case theme.SizeNameInputRadius:
		return Radius
	case theme.SizeNameSelectionRadius:
		return Radius
	case theme.SizeNameLineSpacing:
		return 4
	case theme.SizeNameScrollBarRadius:
		return Radius
	case theme.SizeNameSubHeadingText:
		return 28
	default:
		return theme.DefaultTheme().Size(s)
	}
}
