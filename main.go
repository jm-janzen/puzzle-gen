package main

import (
	"log"
	"os"

	"./drawable"
	"./logger"
	"./screen"
)

const (
	noRooms  = 10
	roomSize = 20
)

func main() {

	logger.InitLogger("logs/debug.log", false)

	world := drawable.NewWorld(noRooms, roomSize)
	for _, r := range world.GetRooms() {
		log.Printf("%#v\n", r.Dimensions)
	}

	// TODO Centre around inner point
	slice := world.GetRect(
		drawable.Dimensions{10, 10, 20, 40},
	)

	view := drawable.NewView(slice)

	s := screen.InitScreen()
	view.Draw(s)

	s.Show()

	os.Exit(0)
}
