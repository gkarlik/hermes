package views

import (
	"github.com/jroimartin/gocui"
)

type View interface {
	gocui.Manager

	Update(g *gocui.Gui, content string)
}
