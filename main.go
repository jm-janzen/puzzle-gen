package main

import (
	"os"

	"./drawable"
	"./screen"
)

const (
	noRooms  = 10
	roomSize = 20
)

func main() {

	screen := screen.InitScreen()

	world := drawable.NewWorld(noRooms, roomSize)
	world.Draw(screen)
	screen.Show()

	os.Exit(0)
}
