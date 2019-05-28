package logical

import (
	"log"

	"github.com/jm-janzen/puzzle-gen/drawable"
	"github.com/jm-janzen/puzzle-gen/types"
)

type Player struct {
	Pos  types.Point
	Name string
}

func (p *Player) MoveTo(cell *drawable.Cell) *Player {
	if string(*cell.GetType()) == string(drawable.Floor) {
		log.Printf("moving '%v' to %v", p, cell)
		p.Pos = types.Point{
			X: cell.X,
			Y: cell.Y,
		}
	} else if string(*cell.GetType()) == string(drawable.Door) {
		cell.SetType(drawable.Floor)
		p.Pos = types.Point{
			X: cell.X,
			Y: cell.Y,
		}
	} else {
		log.Printf("Can't move %v to %v (type %v)",
			p,
			cell,
			cell.Type.String())
	}
	return p
}
