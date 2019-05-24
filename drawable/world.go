package drawable

import (
	"github.com/gdamore/tcell"
)

type World struct {
	Dimensions
	rooms []Room
}

func (r *World) Draw(screen tcell.Screen) {
	for _, room := range r.rooms {
		room.Draw(screen)
	}
}

/*
 * TODO Reimplement as branching, rather than grid
 *      , or maybe just rethink entirely ...
 */
func NewWorld(noRooms int, roomSize int) World {
	var w World
	for noRow := 0; noRow < noRooms/2; noRow++ {
		for noColumn := 0; noColumn < 2; noColumn++ {
			room := NewRoom(
				Dimensions{
					X1: (roomSize * noRow),
					Y1: (roomSize * noColumn),
					X2: roomSize + (roomSize * noRow),
					Y2: roomSize + (roomSize * noColumn),
				},
			)
			w.rooms = append(w.rooms, room)
		}
	}
	return w
}

func (w *World) GetRect(dim Dimensions) []Cell {
	var slice []Cell
	for _, r := range w.rooms {
		for _, c := range r.cells {
			if c.In(dim) {
				slice = append(slice, c)
			}
		}
	}
	return slice
}

func (w *World) GetRooms() []Room {
	return w.rooms
}

func (w *World) AddRoom(r *Room, c Cardinality) {
	// TODO Attach to given room
}
