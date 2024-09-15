package flex

import "fyne.io/fyne/v2"

type FlexBox struct {
	Dir     string  // Direction "columns" or "rows"
	Align   string  // Arrangement of elements along the main axis
	Justify string  // Aligning elements along the secondary axis
	Gap     float32 // Distance between elements
	Padding float32 // Inner padding
}

func (d *FlexBox) MinSize(objects []fyne.CanvasObject) fyne.Size {
	w, h := float32(0), float32(0)
	for i, o := range objects {
		// Check is exist Size of elems
		size := o.MinSize()
		if !o.Size().IsZero() {
			size = o.Size()
		}

		// If dicection is "row" calculate sum width all elems and max heigth of elem
		if d.Dir == "row" {
			if i == len(objects)-1 {
				w += size.Width
			} else {
				w += size.Width + d.Gap
			}
			if h < size.Height {
				h = size.Height
			}
		}
		// If dicection is "columns" calculate sum heigth all elems and max width of elem
		if d.Dir == "column" {
			if i == len(objects)-1 {
				h += size.Height
			} else {
				h += size.Height + d.Gap
			}
			if w < size.Width {
				w = size.Width
			}
		}

	}

	// Return fyne.Size object with internal paddings
	return fyne.NewSize(w+d.Padding*2, h+d.Padding*2)
}

func (d *FlexBox) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	// Count elems
	l := float32(len(objects))

	// Full width fnd height with inner padding
	fullWidth := d.MinSize(objects).Width
	fullHeight := d.MinSize(objects).Height

	// Setting the initial position taking into account internal margins
	pos := fyne.NewPos(d.Padding, d.Padding)

	// Initial position taking into account aligin
	startDrawW := float32(0)
	startDrawH := float32(0)
	// Сдвиг начальной точки элемента с учётом aligin
	shiftW := float32(0)
	shiftH := float32(0)
	// Distance between elements
	Gap := d.Gap

	for i, o := range objects {
		// Check is exist Size of elems
		size := o.MinSize()
		if !o.Size().IsZero() {
			size = o.Size()
		}
		o.Resize(size)

		if d.Dir == "row" {

			switch d.Align {
			case "center":
				startDrawH = containerSize.Height/2 - fullHeight/2
				shiftH = fullHeight/2 - size.Height/2 - d.Padding
			case "end":
				startDrawH = containerSize.Height - fullHeight
				shiftH = fullHeight - size.Height - d.Padding*2
			}

			switch d.Justify {
			case "center":
				startDrawW = containerSize.Width/2 - fullWidth/2
			case "end":
				startDrawW = containerSize.Width - fullWidth
			case "around":
				Gap = (containerSize.Width - fullWidth + (l-1)*d.Gap) / (l - 1)
			case "between":
				Gap = (containerSize.Width - fullWidth + (l+1)*d.Gap) / (l + 1)
				startDrawW += Gap - d.Padding
			}

			if i == 0 {
				pos = pos.Add(fyne.NewPos(startDrawW, startDrawH))
			}
			pos = pos.Add(fyne.NewPos(0, shiftH))

			o.Move(pos)
			pos = pos.Add(fyne.NewPos(size.Width+Gap, -shiftH))
		}

		if d.Dir == "column" {

			switch d.Align {
			case "center":
				startDrawW = containerSize.Width/2 - fullWidth/2
				shiftW = fullWidth/2 - size.Width/2 - d.Padding
			case "end":
				startDrawW = containerSize.Width - fullWidth
				shiftW = fullWidth - size.Width - d.Padding*2
			}

			switch d.Justify {
			case "center":
				startDrawH = containerSize.Height/2 - fullHeight/2
			case "end":
				startDrawH = containerSize.Height - fullHeight
			case "around":
				Gap = (containerSize.Height - fullHeight + (l-1)*d.Gap) / (l - 1)
			case "between":
				Gap = (containerSize.Height - fullHeight + (l+1)*d.Gap) / (l + 1)
				startDrawH += Gap - d.Padding
			}

			if i == 0 {
				pos = pos.Add(fyne.NewPos(startDrawW, startDrawH))
			}
			pos = pos.Add(fyne.NewPos(shiftW, 0))

			o.Move(pos)
			pos = pos.Add(fyne.NewPos(-shiftW, size.Height+Gap))
		}
	}

}
