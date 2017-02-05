package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/urfave/cli"
	"time"
	"github.com/jroimartin/gocui"
	"log"
)

func main() {

	var coinToken string

	httpClient := &http.Client{}
	ctClient := NewClient(httpClient)

	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "coin",
			Value:       "bitcoin",
			Usage:       "coin used for testing",
			Destination: &coinToken,
		},
	}
	app.Name = "Crypto value checker"
	app.Usage = "Tool to check cryptcurrencies prices against coinmarketcap api."

	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panic(err)
	}
	defer g.Close()

	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}


	//todo	move this action to the layout.
	ticker := time.NewTicker(time.Second * 15)
	app.Action = func(c *cli.Context) error {
		if coinToken == "b" {
			for t := range ticker.C {
				fmt.Println(t, ctClient.GetBitcoinPrice())
			}

		} else if coinToken == "e" {
			for t := range ticker.C {
				fmt.Println(t, ctClient.GetEtherPrice())
			}
		} else {
			fmt.Println("unkown option")
		}

		return nil
	}

	app.Run(os.Args)
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("hello", maxX/2-7, maxY/2, maxX/2+7, maxY/2+2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, "Hello world!")
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}