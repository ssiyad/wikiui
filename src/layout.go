package src

import (
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
)

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	if v, err := g.SetView("main", 0, 0, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		if _, err := g.SetCurrentView("main"); err != nil {
			log.Panicln(err)
		}
		v.Wrap = true
		v.Title = "wikiui"
		fmt.Fprintln(v, helpText)
	}
	return nil
}

func setLayout(g *gocui.Gui) {
	g.SetManagerFunc(layout)
}
