package main

/* TODO
[ ] 1. Draw simple grid
[ ] 2. Implement cell (floor, wall, door)
[ ] 3. Implement room (made of up cells)
[ ] 4. Implement world (made of up rooms)
		World has meta-info: Path through rooms
[ ] 5. Implement: Key, Lock, Start, End
[ ] 6. Implement: Procedurally generate
		path from start to end with detours
		to fetch key to unlock door.
*/

import (
	"./drawable"
	"github.com/gdamore/tcell"

	"os"
)

const (
	noRooms  = 10
	roomSize = 10
)

func main() {
	screen, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}
	if err := screen.Init(); err != nil {
		panic(err)
	}

	// TODO Finish this logic and move it to World
	// Draw row of rooms
	for noRoom := range [noRooms]int{} {
		room := drawable.NewRoom(
			drawable.Dimensions{
				X1: (roomSize * noRoom),
				Y1: 0,
				X2: roomSize + (roomSize * noRoom),
				Y2: roomSize + 0,
			},
		)
		room.Draw(screen)
	}
	// TODO Draw rest of rows

	os.Exit(0)
}
