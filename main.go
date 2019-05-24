package main

import (
	"log"
	"os"

	"./drawable"
	"./logger"
	"./logical"
	"./screen"
	"github.com/gdamore/tcell"
)

const (
	noRooms  = 10
	roomSize = 20
)

func main() {

	logger.InitLogger("logs/debug.log", false)

	player := logical.Player{
		Pos: logical.Point{20, 20},
	}

	world := drawable.NewWorld(noRooms, roomSize)

	s := screen.InitScreen()

	for {
		log.Printf("%#v\n", player)

		// FIXME Update by reference ?
		view := *drawable.NewView(player.Pos, world)
		view.Draw(s)

		ev := s.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEscape, tcell.KeyEnter:
				return
			case tcell.KeyDown:
				player.Pos.Y += 1
			case tcell.KeyUp:
				player.Pos.Y -= 1
			case tcell.KeyRight:
				player.Pos.X += 1
			case tcell.KeyLeft:
				player.Pos.X -= 1
			}
		}
	}

	os.Exit(0)
}
