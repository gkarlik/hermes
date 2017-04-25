package views

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"time"
)

type InputView struct {
	Name string
	nick string
}

func NewInputView(nick string) *InputView {
	return &InputView{
		Name: "input",
		nick: nick,
	}
}

func (iv *InputView) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	v, err := g.SetView(iv.Name, -1, maxY-2, maxX, maxY)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Autoscroll = false
		v.Editable = true
		v.Wrap = false
		v.Frame = false
		v.FgColor = gocui.ColorWhite

		if err := g.SetKeybinding(iv.Name, gocui.KeyEnter, gocui.ModNone, handleCommand(iv.nick)); err != nil {
			return err
		}

		if _, err := g.SetCurrentView(iv.Name); err != nil {
			return err
		}
	}
	return nil
}

func handleCommand(nick string) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		view, err := g.View("main")
		if err != nil {
			return err
		}

		fmt.Fprintf(view, "%s <%s> %s", time.Now().Format("15:04:05"), nick, v.ViewBuffer())

		v.Clear()
		v.SetCursor(0, 0)

		return nil
	}
}

func (iv *InputView) Update(g *gocui.Gui, content string) {
	g.Execute(func(g *gocui.Gui) error {
		v, err := g.View(iv.Name)
		if err != nil {
			return err
		}
		v.Clear()
		fmt.Fprint(v, content)

		return nil
	})
}
