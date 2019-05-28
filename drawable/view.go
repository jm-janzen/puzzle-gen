package drawable

import (
	"github.com/gdamore/tcell"
	"github.com/jm-janzen/puzzle-gen/types"
)

type View struct {
	Centre types.Point
	cells  []Cell
}

func NewView(centre types.Point, world *World) View {
	const (
		verticalView   = 48
		horizontalView = 106
	)
	slice := world.GetRect(
		Dimensions{
			Y1: centre.Y - verticalView,
			X1: centre.X - horizontalView,
			Y2: centre.Y + verticalView,
			X2: centre.X + horizontalView,
		},
	)
	return View{
		Centre: centre,
		cells:  toAbsolute(slice),
	}
}

func (v *View) Draw(screen tcell.Screen) {
	screen.Clear()
	for _, c := range v.cells {
		// FIXME This is a completely wrong-headed approach
		//to displaying player
		if c.X == v.Centre.X && c.Y == v.Centre.Y {
			c.Rune = '@' // Player :p
		}
		c.Draw(screen)
	}
	screen.Sync()
}

// Hug any offsets to (0,0) and so on
func toAbsolute(cells []Cell) []Cell {
	minX := cells[0].X
	minY := cells[0].Y
	for i := 0; i < len(cells); i++ {
		c := cells[i]
		cells[i].X = c.X - minX
		cells[i].Y = c.Y - minY
	}
	return cells
}
