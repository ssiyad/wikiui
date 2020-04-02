package src

import (
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"

	strip "github.com/grokify/html-strip-tags-go"
	"github.com/jroimartin/gocui"
)

func jumpToMain(g *gocui.Gui, v *gocui.View) error {
	err := g.DeleteView(v.Name())
	if err != nil {
		log.Panicln(err)
	}
	if v, err = g.SetCurrentView("main"); err != nil {
		if err != gocui.ErrUnknownView {
			log.Panicln(err)
		}
		g.SetViewOnTop(v.Name())
	}
	return nil
}

func jumpToSearchBox(g *gocui.Gui, v *gocui.View) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("searchbox", maxX/2-20, maxY/2-1, maxX/2+20, maxY/2+1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		_, err := g.SetCurrentView("searchbox")
		if err != nil {
			log.Panicln(err)
		}
		v.Title = "Search"
		v.Editable = true
		v.Overwrite = true
		fmt.Fprintln(v, "...")
	}
	return nil
}

func jumpToSaved(g *gocui.Gui, v *gocui.View) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("saved", maxX/2-10, maxY/2-20, maxX/2+20, maxY/2+10); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		_, err := g.SetCurrentView("saved")
		if err != nil {
			log.Panicln(err)
		}
		v.Title = "Saved"
		v.Highlight = true
		v.SelFgColor = gocui.ColorGreen
		var homePath, _ = os.UserHomeDir()
		var savePath = path.Join(homePath, saveDir)
		files, _ := ioutil.ReadDir(savePath)
		for _, f := range files {
			fmt.Fprintln(v, f.Name())
		}
	}
	return nil
}

func showResult(g *gocui.Gui, v *gocui.View) error {
	var keyword string
	var err error

	_, cy := v.Cursor()
	if keyword, err = v.Line(cy); err != nil {
		log.Panicln(err)
	}

	err = g.DeleteView(v.Name())
	if err != nil {
		log.Panicln(err)
	}

	maxX, maxY := g.Size()
	if v, err := g.SetView("result", maxX/2-20, maxY/2-3, maxX/2+20, maxY/2+3); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		result := searchWiki(keyword)
		for _, name := range result.Name {
			fmt.Fprintln(v, name)
		}
		_, err = g.SetCurrentView("result")
		if err != nil {
			log.Panicln(err)
		}
	}

	return nil
}

func showPage(g *gocui.Gui, v *gocui.View) error {
	var keyword string
	var err error
	var title string
	var content string

	_, cy := v.Cursor()
	if keyword, err = v.Line(cy); err != nil {
		log.Panicln(err)
	}

	err = g.DeleteView(v.Name())
	if err != nil {
		log.Panicln(err)
	}

	v, err = g.View("main")
	if err != nil {
		log.Panicln(err)
	}

	_, err = g.SetCurrentView(v.Name())
	if err != nil {
		log.Panicln(err)
	}

	result := getWikiPage(keyword)
	v.Clear()
	v.Editable = true
	for _, data := range result.Query.Pages {
		title += data.Title
		content += data.Extract
	}
	v.Title = title + " - wikiui"
	v.Write([]byte(strip.StripTags(html.UnescapeString(strings.TrimSpace(content)))))

	return nil
}

func readSaved(g *gocui.Gui, v *gocui.View) error {
	var l string
	var err error

	_, cy := v.Cursor()
	if l, err = v.Line(cy); err != nil {
		log.Panicln(err)
	}

	var homePath, _ = os.UserHomeDir()
	var readPath = path.Join(homePath, saveDir, l)
	cont, err := ioutil.ReadFile(readPath)
	if err != nil {
		log.Panicln(err)
	}

	if _, err := g.SetCurrentView("main"); err != nil {
		if err != gocui.ErrUnknownView {
			log.Panicln(err)
		}
	}
	err = g.DeleteView(v.Name())
	if err != nil {
		log.Panicln(err)
	}

	v, err = g.View("main")
	if err != nil {
		log.Panicln(err)
	}
	v.Clear()
	v.Title = l + " - wikiui"
	v.Editable = true
	v.Write(cont)

	return nil
}

func writeArticle(g *gocui.Gui, v *gocui.View) error {
	var homePath, _ = os.UserHomeDir()
	var writePath = path.Join(homePath, saveDir, strings.Trim(v.Title, " - wikiui"))

	err := ioutil.WriteFile(writePath, []byte(v.Buffer()), 0700)
	if err != nil {
		log.Panicln(err)
	}
	return nil
}

func moveDown(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		cx, cy := v.Cursor()
		if err := v.SetCursor(cx, cy+1); err != nil {
			ox, oy := v.Origin()
			if err := v.SetOrigin(ox, oy+1); err != nil {
				return err
			}
		}
	}
	return nil
}

func moveUp(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		ox, oy := v.Origin()
		cx, cy := v.Cursor()
		if err := v.SetCursor(cx, cy-1); err != nil && oy > 0 {
			if err := v.SetOrigin(ox, oy-1); err != nil {
				return err
			}
		}
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
