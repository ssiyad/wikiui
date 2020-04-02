package src

import "github.com/jroimartin/gocui"


func setProperties(g *gocui.Gui) {
	g.SelFgColor = gocui.ColorBlue
	g.Highlight = true
	g.Cursor = true
	g.Mouse = false
	g.ASCII = false
}