package drawable

import (
	"fmt"

	"github.com/gdamore/tcell"
	"github.com/jm-janzen/puzzle-gen/types"
)

// FIXME These are not so much dimensions as coordinates, or bounds
type Dimensions struct {
	Y1, X1, Y2, X2 int
}

func (d Dimensions) String() string {
	return fmt.Sprintf("{%v,%v;%v,%v}", d.Y1, d.X1, d.Y2, d.X2)
}

type Room struct {
	Dimensions
	cells []Cell
}

func (r *Room) Draw(screen tcell.Screen) {
	for _, cell := range r.cells {
		cell.Draw(screen)
	}
}

// Checks if given Point is at half-way coordinates
//along an edge of the given Dimensions rectangle ...
func isMiddleEdge(dim Dimensions, p types.Point) bool {
	isMidEdge := false
	var (
		topMiddle   = types.Point{(dim.X2 + dim.X1) / 2, dim.Y1}
		botMiddle   = types.Point{(dim.X2 + dim.X1) / 2, dim.Y2 - 1}
		rightMiddle = types.Point{dim.X2 - 1, (dim.Y2 + dim.Y1) / 2}
		leftMiddle  = types.Point{dim.X1, (dim.Y2 + dim.Y1) / 2}
	)
	if p == topMiddle {
		isMidEdge = true
	} else if p == botMiddle {
		isMidEdge = true
	} else if p == rightMiddle {
		isMidEdge = true
	} else if p == leftMiddle {
		isMidEdge = true
	}
	return isMidEdge
}
func NewRoom(dim Dimensions) Room {
	var r Room
	r = Room{
		Dimensions: dim,
	}
	for x := dim.X1; x < dim.X2; x++ {
		for y := dim.Y1; y < dim.Y2; y++ {
			var t Type
			/* West, East, North, South */
			if dim.X1 == x || dim.X2 == x+1 || dim.Y1 == y || dim.Y2 == y+1 {
				if isMiddleEdge(dim, types.Point{x, y}) {
					t = Door
				} else {
					t = Wall
				}
			} else {
				t = Floor
			}

			r.cells = append(r.cells, NewCell(x, y, t))
		}
	}
	return r
}
