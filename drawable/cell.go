package drawable

import (
	"fmt"
	"log"

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
		"hall",
		"wall",
		"door",
	}[t]
}

type Cell struct {
	X, Y int
	Type Type

	// FIXME Cell should only contain logical data values
	//not drawable members like before
	Rune  rune
	Style tcell.Style
}

func NewCell(x int, y int, t Type) Cell {
	var (
		style tcell.Style
		rune  rune
	)
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
		Type:  t,
		Style: style,
		Rune:  rune,
	}
}

func (c Cell) String() string {
	return fmt.Sprintf("{x:%v,y:%v,typ:%v}", c.X, c.Y, c.Type)
}

func (c *Cell) Draw(screen tcell.Screen) {
	screen.SetContent(c.X, c.Y, c.Rune, nil, c.Style)
}

func (c *Cell) In(dim Dimensions) bool {
	if c.X >= dim.X1 && c.X < dim.X2 && c.Y >= dim.Y1 && c.Y < dim.Y2 {
		return true
	} else {
		return false
	}
}

// Pointer-receiver pattern
func (c Cell) GetType() *Type {
	return &c.Type
}
func (c *Cell) SetType(newType Type) {
	t := c.GetType()
	*t = newType
	c.Type = *t
	log.Printf("me am: %#v, wanna be: %#v", c, *t)
}
