package drawable

import (
	"github.com/gdamore/tcell"
)

type Type uint

const (
	Floor Type = iota
	Hall
	Wall
	Door
)

func (t Type) String() string {
	return [...]string{
		"floor",
		"wall",
		"door",
		"hall",
	}[t]
}

type Cardinality uint

const (
	North Cardinality = iota
	East
	South
	West
)

func (c Cardinality) String() string {
	return [...]string{
		"north",
		"south",
		"east",
		"west",
	}[c]
}

type Cell struct {
	X, Y  int
	Rune  rune
	Style tcell.Style
}

func NewCell(x int, y int, t Type) Cell {
	var style tcell.Style
	var rune rune
	switch t {
	case Floor, Hall:
		style = tcell.StyleDefault.Foreground(tcell.ColorGrey)
		rune = tcell.RuneBullet
	case Wall:
		style = tcell.StyleDefault.Foreground(tcell.ColorWhite)
		rune = tcell.RuneBlock
	case Door:
		style = tcell.StyleDefault.Foreground(tcell.ColorGold)
		rune = tcell.RunePlus
	}
	return Cell{
		X:     x,
		Y:     y,
		Style: style,
		Rune:  rune,
	}
}

func (c *Cell) Draw(screen tcell.Screen) {
	screen.SetContent(c.X, c.Y, c.Rune, nil, c.Style)
}
