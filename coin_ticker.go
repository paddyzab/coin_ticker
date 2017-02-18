package main

import (
	"net/http"
	"os"

	"github.com/urfave/cli"
	"github.com/jroimartin/gocui"
	"fmt"
)

func main() {
	app := cli.NewApp()
	app.Name = "Crypto coin value checker"
	app.Usage = "Tool to check cryptcurrencies prices against coinmarketcap api."

	app.Action = printPrice()

	app.Run(os.Args)
}

func printPrice() func(c* cli.Context) error {
	return func(c* cli.Context) error {
		httpClient := &http.Client{}
		ctClient := NewClient(httpClient)

		fmt.Printf("\nBTC: %s, ETH: %s \n", ctClient.GetBitcoinPrice(), ctClient.GetEtherPrice())
		return nil
	}
}

func displayPrice(pw *PriceWidget) func(g *gocui.Gui, v *gocui.View) error {
	httpClient := &http.Client{}
	ctClient := NewClient(httpClient)

	return func(g *gocui.Gui, v *gocui.View) error {
		if pw.name == ether {
			return labelSet(pw, ctClient.GetEtherPrice())
		} else {
			return labelSet(pw, ctClient.GetBitcoinPrice())
		}
	}
}

func labelSet(pw *PriceWidget, label string) error {
	return pw.SetVal(label)
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}