package src

import (
	"log"

	"github.com/jroimartin/gocui"
)

func setNavigation(g *gocui.Gui) {
	// Keys on 'main' buffer
	if err := g.SetKeybinding("main", gocui.KeyCtrlS, gocui.ModNone, jumpToSearchBox); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("main", gocui.KeyCtrlO, gocui.ModNone, jumpToSaved); err != nil {
		log.Panicln(err)
	}

	// Keys on 'result' buffer
	if err := g.SetKeybinding("result", gocui.KeyArrowUp, gocui.ModNone, moveUp); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("result", gocui.KeyArrowDown, gocui.ModNone, moveDown); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("result", gocui.KeyEnter, gocui.ModNone, showPage); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("result", gocui.KeyCtrlQ, gocui.ModNone, jumpToMain); err != nil {
		log.Panicln(err)
	}

	// keys on 'searchbox' buffer
	if err := g.SetKeybinding("searchbox", gocui.KeyEnter, gocui.ModNone, showResult); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("searchbox", gocui.KeyCtrlQ, gocui.ModNone, jumpToMain); err != nil {
		log.Panicln(err)
	}

	// Keys on 'saved' buffer
	if err := g.SetKeybinding("saved", gocui.KeyArrowUp, gocui.ModNone, moveUp); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("saved", gocui.KeyArrowDown, gocui.ModNone, moveDown); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("saved", gocui.KeyEnter, gocui.ModNone, readSaved); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("saved", gocui.KeyCtrlQ, gocui.ModNone, jumpToMain); err != nil {
		log.Panicln(err)
	}
}
