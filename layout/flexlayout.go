package layout

import (
	"math"

	"github.com/fyne-io/fyne"
	"github.com/fyne-io/fyne/theme"
)

type flexLayout struct {
	Cols int
}

func (g *flexLayout) countRows(objects []fyne.CanvasObject) int {
	return int(math.Ceil(float64(len(objects)) / float64(g.Cols)))
}

// Layout is called to pack all child objects into a specified size.
// For a FlexLayout this will pack objects into a table format with the number
// of columns specified in our constructor.
func (g *flexLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	rows := g.countRows(objects)

	padWidth := (g.Cols - 1) * theme.Padding()
	padHeight := (rows - 1) * theme.Padding()

	cellWidth := float64(size.Width-padWidth) / float64(g.Cols)
	cellHeight := float64(size.Height-padHeight) / float64(rows)

	row, col := 0, 0
	for i, child := range objects {
		x1 := getLeading(cellWidth, col)
		y1 := getLeading(cellHeight, row)
		x2 := getTrailing(cellWidth, col)
		y2 := getTrailing(cellHeight, row)

		child.Move(fyne.NewPos(x1, y1))
		colwidth := 200
		if i%2 == 1 {
			colwidth = y2 - y1
		}
		child.Resize(fyne.NewSize(x2-x1, colwidth))

		if (i+1)%g.Cols == 0 {
			row++
			col = 0

		} else {
			col++

		}
	}
}

// MinSize finds the smallest size that satisfies all the child objects.
// For a FlexLayout this is the size of the largest child object multiplied by
// the required number of columns and rows, with appropriate padding between
// children.
func (g *flexLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	rows := g.countRows(objects)
	minSize := fyne.NewSize(0, 0)
	for _, child := range objects {
		minSize = minSize.Union(child.MinSize())
	}

	minContentSize := fyne.NewSize(minSize.Width*g.Cols, minSize.Height*rows)
	return minContentSize.Add(fyne.NewSize(theme.Padding()*(g.Cols-1), theme.Padding()*(rows-1)))
}

// NewFlexLayout returns a new FlexLayout instance
func NewFlexLayout(cols int) fyne.Layout {
	return &flexLayout{cols}
}
