package views

import (
	"fmt"
	"github.com/jroimartin/gocui"
)

type MainView struct {
	Name string
}

func NewMainView() *MainView {
	return &MainView{
		Name: "main",
	}
}

func (mv *MainView) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	v, err := g.SetView(mv.Name, -1, 0, maxX, maxY-2)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Autoscroll = true
		v.Editable = false
		v.Wrap = true
		v.Frame = false
		v.FgColor = gocui.ColorWhite
	}
	return nil
}

func (mv *MainView) Update(g *gocui.Gui, content string) {
	g.Execute(func(g *gocui.Gui) error {
		v, err := g.View(mv.Name)
		if err != nil {
			return err
		}
		v.Clear()
		fmt.Fprint(v, content)

		return nil
	})
}
