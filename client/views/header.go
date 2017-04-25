package views

import (
	"fmt"
	"github.com/jroimartin/gocui"
)

type HeaderView struct {
	Name  string
	title string
}

func NewHeaderView(title string) *HeaderView {
	return &HeaderView{
		Name:  "header",
		title: title,
	}
}

func (hv *HeaderView) Layout(g *gocui.Gui) error {
	maxX, _ := g.Size()
	v, err := g.SetView(hv.Name, -1, -1, maxX, 1)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Autoscroll = false
		v.Editable = false
		v.Wrap = false
		v.Frame = false
		v.FgColor = gocui.ColorWhite
		v.BgColor = gocui.ColorBlue

		fmt.Fprint(v, hv.title)
	}
	return nil
}

func (hv *HeaderView) Update(g *gocui.Gui, content string) {
	g.Execute(func(g *gocui.Gui) error {
		v, err := g.View(hv.Name)
		if err != nil {
			return err
		}
		v.Clear()
		fmt.Fprint(v, content)

		return nil
	})
}
