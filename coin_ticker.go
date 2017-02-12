package main

import (
	"net/http"
	"os"

	"github.com/urfave/cli"
	"github.com/jroimartin/gocui"
	"log"
)

func main() {
	app := cli.NewApp()
	app.Name = "Crypto value checker"
	app.Usage = "Tool to check cryptcurrencies prices against coinmarketcap api."

	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panic(err)
	}
	defer g.Close()

	bitcoinPrice := NewPriceWidget("bitcoin", 1, 10, 50, "here be price")
	etherPrice := NewPriceWidget("ether", 1, 30, 50, "here be price")
	currentButton := NewButtonWidget("fetch", 52, 7, "fetch price", displayPrice(bitcoinPrice))
	g.SetManager(bitcoinPrice, etherPrice, currentButton)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}

	app.Run(os.Args)
}

func displayPrice(pw *PriceWidget) func(g *gocui.Gui, v *gocui.View) error {
	httpClient := &http.Client{}
	ctClient := NewClient(httpClient)

	return func(g *gocui.Gui, v *gocui.View) error {
		return labelSet(pw, ctClient.GetBitcoinPrice())
	}
}

func labelSet(pw *PriceWidget, label string) error {
	return pw.SetVal(label)
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}