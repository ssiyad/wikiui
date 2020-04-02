package src

import (
	"log"

	"github.com/jroimartin/gocui"
)

func setBindings(g *gocui.Gui) {
	if err := g.SetKeybinding("main", gocui.KeyCtrlQ, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("main", gocui.KeyArrowDown, gocui.ModNone, moveDown); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("main", gocui.KeyArrowUp, gocui.ModNone, moveUp); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("main", gocui.KeyCtrlW, gocui.ModNone, writeArticle); err != nil {
		log.Panicln(err)
	}
}
