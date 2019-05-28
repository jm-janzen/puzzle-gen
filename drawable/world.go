package drawable

import (
	"log"

	"github.com/gdamore/tcell"
	"github.com/jm-janzen/puzzle-gen/types"
)

type World struct {
	Dimensions
	rooms []Room
}

func (r *World) Draw(screen tcell.Screen) {
	for i := 0; i < len(r.rooms); i++ {
		room := r.rooms[i]
		room.Draw(screen)
	}
}

/*
 * TODO Reimplement as branching, rather than grid
 *      , or maybe just rethink entirely ...
 * 		; yeah, just take an arbitrarily large int for now
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
	for i := 0; i < len(w.rooms); i++ {
		r := w.rooms[i]
		for j := 0; j < len(r.cells); j++ {
			c := r.cells[j]
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

func (w *World) GetCell(p types.Point) *Cell {
	var foundCell *Cell
	for i := 0; i < len(w.rooms); i++ {
		r := w.rooms[i]
		for j := 0; j < len(r.cells); j++ {
			c := r.cells[j]
			if c.X == p.X && c.Y == p.Y {
				foundCell = &c
			}
		}
	}
	return foundCell
}
