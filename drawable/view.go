package drawable

import "github.com/gdamore/tcell"

type Point struct {
	x, y uint
}
type Rect struct {
	xleft, xright, ytop, ybot Point
}
type View struct {
	centre Point
	cells  []Cell
}

func NewView(cells []Cell) View {
	// TODO Convert cells' coords to absolute values for display
	return View{
		centre: Point{
			0, 0,
		},
		cells: cells,
	}
}

func (v *View) Draw(screen tcell.Screen) {
	for _, c := range v.cells {
		c.Draw(screen)
	}
}
