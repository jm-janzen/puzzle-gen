package main

import (
	"log"
	"os"

	"github.com/gdamore/tcell"
	"github.com/jm-janzen/puzzle-gen/drawable"
	"github.com/jm-janzen/puzzle-gen/logger"
	"github.com/jm-janzen/puzzle-gen/logical"
	"github.com/jm-janzen/puzzle-gen/screen"
	"github.com/jm-janzen/puzzle-gen/types"
)

const (
	noRooms  = 16
	roomSize = 25
)

func main() {

	logger.InitLogger("logs/debug.log", false)

	player := logical.Player{
		Name: "you",
		Pos:  types.Point{1, 1},
	}

	// FIXME Need to be able to update world
	world := drawable.NewWorld(noRooms, roomSize)

	s := screen.InitScreen()

	// TODO Just set view to max size of display - buf
	x, y := s.Size()
	log.Printf("Display dimensions: %+v, %+v\n", x, y)

	for {
		playerTarget := player.Pos

		// FIXME Update by reference ?
		view := drawable.NewView(player.Pos, &world)
		view.Draw(s)

		ev := s.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEscape, tcell.KeyEnter:
				return

			case tcell.KeyDown:
				playerTarget = types.Point{player.Pos.X, (player.Pos.Y + 1)}
			case tcell.KeyUp:
				playerTarget = types.Point{player.Pos.X, (player.Pos.Y - 1)}
			case tcell.KeyRight:
				playerTarget = types.Point{(player.Pos.X + 1), player.Pos.Y}
			case tcell.KeyLeft:
				playerTarget = types.Point{(player.Pos.X - 1), player.Pos.Y}
			}

		}
		if playerTarget != player.Pos {
			targetCell := *world.GetCell(playerTarget)
			log.Printf("Before: %v", targetCell)
			player.MoveTo(&targetCell)
			log.Printf("After : %v", targetCell)
		}
	}

	os.Exit(0)
}
