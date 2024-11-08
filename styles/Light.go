package styles

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type Light struct{}

func (l *Light) Color(c fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {

	switch c {

	case theme.ColorNameForeground:
		return Grey_800
	case "darktext":
		return Grey_600
	case theme.ColorNameHyperlink:
		return Grey_600
	case theme.ColorNameHeaderBackground:
		return Red

	case theme.ColorNameBackground:
		return Grey_200
	case theme.ColorNameMenuBackground:
		return Grey_100
	case theme.ColorNameInputBackground:
		return White
	case theme.ColorNameOverlayBackground:
		return Grey_100

	case theme.ColorNameInputBorder:
		return Grey_300
	case theme.ColorNameSeparator:
		return Grey_300

	case theme.ColorNamePrimary:
		return Red
	case theme.ColorNameForegroundOnPrimary:
		return White
	case theme.ColorNameButton:
		return Grey_300
	case theme.ColorNameSuccess:
		return Grey_300
	case theme.ColorNameForegroundOnSuccess:
		return Grey_800

	case theme.ColorNameHover:
		return Hover
	case theme.ColorNameSelection:
		return Grey_50
	case theme.ColorNamePressed:
		return White
	case theme.ColorNameDisabled:
		return Grey_600
	case theme.ColorNameDisabledButton:
		return Disabled

	case theme.ColorNameScrollBar:
		return Grey_600
	case theme.ColorNameShadow:
		return Shadow

	default:
		return theme.DefaultTheme().Color(c, v)
	}
}

func (l *Light) Font(s fyne.TextStyle) fyne.Resource {
	// if s.Monospace {
	// 	return static.ResRobotoMonoRegularTtf
	// }
	// if s.Bold {
	// 	if s.Italic {
	// 		return static.ResMontserratBoldTtf
	// 	}
	// 	return static.ResRobotoBoldTtf
	// }
	// if s.Italic {
	// 	return static.ResRobotoBoldTtf
	// }
	// return static.ResRobotoRegularTtf
	return theme.DefaultTheme().Font(s)
}

func (l *Light) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(n)
}

func (l *Light) Size(s fyne.ThemeSizeName) float32 {
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
		return 40
	case theme.SizeNameScrollBarRadius:
		return Radius
	case theme.SizeNameSubHeadingText:
		return 40
	case theme.SizeNameHeadingText:
		return 40
	default:
		return theme.DefaultTheme().Size(s)
	}
}
