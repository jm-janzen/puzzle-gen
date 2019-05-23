package drawable

import (
	"github.com/gdamore/tcell"
)

type World struct {
	Dimensions
	rooms []Room
}

func (r *World) Draw(screen  tcell.Screen) {
	for _,room := range r.rooms {
		room.Draw(screen)
	}
}

func NewWorld(dim Dimensions) World {
	var w World
	w = World{
		Dimensions: dim,
	}
	// FIXME Gonna need a room size. Probably should just
	//fix it to symmetrical boxes

	return w
}
