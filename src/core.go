package src

import (
	"log"
	"os"
	"path"

	"github.com/jroimartin/gocui"
)

// Init function
func Init() {
	g, err := gocui.NewGui(gocui.Output256)
	if err != nil {
		panic(err)
	}
	defer g.Close()

	setProperties(g)
	setLayout(g)
	setBindings(g)
	setNavigation(g)

	var homePath, _ = os.UserHomeDir()
	var savePath = path.Join(homePath, saveDir)

	_ = os.Mkdir(savePath, 0700)

	if err := g.MainLoop(); err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

const saveDir = `.wikiui`

const helpText = `
CLI for Wikipedia!
Made with <3 by Sabu Siyad
https://github.com/ssiyad/wikiui 

C-O 		Open saved
C-S 		Search
C-W			Save for offline
C-Q 		Quit`
