package drawable

import "github.com/gdamore/tcell"

// FIXME These are not so much dimensions as coordinates, or bounds
type Dimensions struct {
	Y1, X1, Y2, X2 int
}

type cells []Cell
type Room struct {
	Dimensions
	cells
}

func (r *Room) Draw(screen  tcell.Screen) {
	for _,cell := range r.cells {
		cell.Draw(screen)
	}
}

func NewRoom(dim Dimensions) Room {
	var r Room
	r = Room{
		Dimensions: dim,
	}
	for i := dim.X1 ; i < dim.X2 ; i++ {
		for j := dim.Y1 ; j < dim.Y2 ; j++ {
			var t Type
			/* West, East, North, South */
			if ( dim.X1 == i || dim.X2 == i + 1 || dim.Y1 == j || dim.Y2 == j + 1) {
				t = Wall
			} else {
				t = Floor
			}
			r.cells = append(r.cells, NewCell(i,j,t))
		}
	}
	return r
}
