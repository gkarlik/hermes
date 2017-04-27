package views

import (
	"fmt"

	"github.com/gkarlik/hermes"
	"github.com/jroimartin/gocui"
)

type InputView struct {
	Name   string
	main   *MainView
	status *StatusView
	client *hermes.Client
}

func NewInputView(c *hermes.Client, mv *MainView, sv *StatusView) *InputView {
	return &InputView{
		Name:   "input",
		main:   mv,
		status: sv,
		client: c,
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

		if err := g.SetKeybinding(iv.Name, gocui.KeyEnter, gocui.ModNone, handleCommand(iv)); err != nil {
			return err
		}

		if _, err := g.SetCurrentView(iv.Name); err != nil {
			return err
		}
	}
	return nil
}

func (iv *InputView) Init(g *gocui.Gui) error {
	messages, err := iv.client.Init()
	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case msg := <-messages:
				iv.main.Update(g, fmt.Sprintf("<%s> %s", msg.Sender, msg.Body))
			}
		}
	}()

	return nil
}

func handleCommand(iv *InputView) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		cmd := v.ViewBuffer()

		msg := &hermes.Message{
			Type:      hermes.BROADCAST,
			Sender:    iv.client.ID,
			Recipient: "*",
			Body:      cmd,
		}
		if err := iv.client.SendMessage(msg); err != nil {
			iv.status.Update(g, "Cannot send command to the server!")

			return err
		}

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
