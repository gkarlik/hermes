package main

import (
	"flag"
	"fmt"

	"github.com/gkarlik/hermes"
	"github.com/gkarlik/hermes/client/views"
	"github.com/jroimartin/gocui"
)

const (
	appName = "[hermes]"
	version = "0.1"
)

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func keybindings(g *gocui.Gui) error {
	if err := g.SetKeybinding("", gocui.KeyCtrlX, gocui.ModNone, quit); err != nil {
		return err
	}
	return nil
}

func main() {
	nick := flag.String("nick", "unknown", "name of the user")
	server := flag.String("server", "tcp://localhost:8080", "address of the server")
	flag.Parse()

	cfg := hermes.NewConfig(*nick, *server)
	client := hermes.NewClient(cfg)
	defer client.Close()

	gui, err := gocui.NewGui(gocui.Output256)
	if err != nil {
		panic("cannot create gui")
	}
	defer gui.Close()

	hv := views.NewHeaderView(fmt.Sprintf("%s v%s", appName, version))
	mv := views.NewMainView()
	sv := views.NewStatusView()
	iv := views.NewInputView(client, mv, sv)

	gui.Cursor = true
	gui.SetManager(hv, mv, sv, iv)

	iv.Init(gui)

	if err := keybindings(gui); err != nil {
		panic(err)
	}

	if err := gui.MainLoop(); err != nil && err != gocui.ErrQuit {
		panic(err)
	}
}
