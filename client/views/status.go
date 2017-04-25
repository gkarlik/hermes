package views

import (
	"fmt"
	"github.com/jroimartin/gocui"
)

type StatusView struct {
	Name string
}

func NewStatusView() *StatusView {
	return &StatusView{
		Name: "status",
	}
}

func (sv *StatusView) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	v, err := g.SetView(sv.Name, -1, maxY-3, maxX, maxY-1)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Autoscroll = false
		v.Editable = false
		v.Wrap = false
		v.Frame = false
		v.FgColor = gocui.ColorWhite
		v.BgColor = gocui.ColorGreen
	}
	return nil
}

func (sv *StatusView) Update(g *gocui.Gui, content string) {
	g.Execute(func(g *gocui.Gui) error {
		v, err := g.View(sv.Name)
		if err != nil {
			return err
		}
		v.Clear()
		fmt.Fprint(v, content)

		return nil
	})
}
