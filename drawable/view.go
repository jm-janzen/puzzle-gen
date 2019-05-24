package drawable

import (
	"../logical"
	"github.com/gdamore/tcell"
)

type View struct {
	Centre logical.Point
	cells  []Cell
}

func NewView(centre logical.Point, world World) *View {
	const viewArea = 10
	slice := world.GetRect(
		Dimensions{
			Y1: centre.Y - viewArea,
			X1: centre.X - viewArea,
			Y2: centre.Y + viewArea,
			X2: centre.X + viewArea,
		},
	)
	return &View{
		Centre: centre,
		cells:  toAbsolute(slice),
	}
}

func (v *View) Draw(screen tcell.Screen) {
	for _, c := range v.cells {
		// FIXME This is a completely wrong-headed approach
		//to displaying player
		if c.X == v.Centre.X && c.Y == v.Centre.Y {
			c.Rune = '@' // Player :p
		}
		c.Draw(screen)
	}
	screen.Show()
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
