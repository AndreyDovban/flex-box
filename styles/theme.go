package styles

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type Light struct {
}

func (l *Light) Color(c fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	switch c {
	case theme.ColorNameBackground:
		return Grey_100
	case theme.ColorNameForeground:
		ColorNameForeground = Grey_800
		return Grey_800

	case theme.ColorNameInputBackground:
		return White
	case theme.ColorNameInputBorder:
		return Grey_300
	case theme.ColorNameMenuBackground:
		return Grey_100
		/*

		 */

	case theme.ColorNameButton:
		return Grey_300
	case theme.ColorNameHover:
		return Grey_300
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
	case theme.ColorNameDisabled:
		return color.RGBA{255, 255, 255, 100}
	case theme.ColorNameDisabledButton:
		return color.RGBA{255, 255, 255, 100}

	default:
		return theme.DefaultTheme().Color(c, v)
	}
}

func (l *Light) Font(s fyne.TextStyle) fyne.Resource {
	// if s.Monospace {
	// 	return static.ResInterRegularTtf
	// }
	// if s.Bold {
	// 	if s.Italic {
	// 		return static.ResInterSemiBoldTtf
	// 	}
	// 	return static.ResInterSemiBoldTtf
	// }
	// if s.Italic {
	// 	return static.ResInterRegularTtf
	// }
	// return static.ResInterRegularTtf
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
		return 4
	case theme.SizeNameScrollBarRadius:
		return Radius
	case theme.SizeNameSubHeadingText:
		return 28
	default:
		return theme.DefaultTheme().Size(s)
	}
}

type Dark struct {
}

func (l *Dark) Color(c fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	switch c {
	case theme.ColorNameBackground:
		return Grey_900
	case theme.ColorNameForeground:
		ColorNameForeground = Grey_200
		return Grey_100

	case theme.ColorNameInputBackground:
		return White
	case theme.ColorNameInputBorder:
		return Grey_300
	case theme.ColorNameMenuBackground:
		return Grey_800
		/*

		 */

	case theme.ColorNameButton:
		return Grey_300
	case theme.ColorNameHover:
		return Grey_300
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
	case theme.ColorNameDisabled:
		return color.RGBA{255, 255, 255, 100}
	case theme.ColorNameDisabledButton:
		return color.RGBA{255, 255, 255, 100}

	default:
		return theme.DefaultTheme().Color(c, v)
	}
}

func (l *Dark) Font(s fyne.TextStyle) fyne.Resource {
	// if s.Monospace {
	// 	return static.ResInterRegularTtf
	// }
	// if s.Bold {
	// 	if s.Italic {
	// 		return static.ResInterSemiBoldTtf
	// 	}
	// 	return static.ResInterSemiBoldTtf
	// }
	// if s.Italic {
	// 	return static.ResInterRegularTtf
	// }
	// return static.ResInterRegularTtf
	return theme.DefaultTheme().Font(s)
}

func (l *Dark) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(n)
}

func (l *Dark) Size(s fyne.ThemeSizeName) float32 {
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

type Dark2 struct {
	*Dark
}

func (l *Dark2) Size(s fyne.ThemeSizeName) float32 {
	switch s {
	case theme.SizeNameInnerPadding:
		return 0
	default:
		return l.Dark.Size(s)
	}
}
