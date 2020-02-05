package cui

import (
	"github.com/jroimartin/gocui"
	"log"
)

func Start() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Highlight = true
	g.Cursor = true
	g.SelFgColor = gocui.ColorGreen

	g.SetManagerFunc(layout)

	if err := keybindings(g); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop();  err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}