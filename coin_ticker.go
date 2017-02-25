package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/urfave/cli"
	"github.com/logrusorgru/aurora"
)

const (
	duration = time.Second * 120
	timeFormat = time.Kitchen
)

// output colorizer
var au aurora.Aurora

func main() {
	au = aurora.NewAurora(true)

	app := cli.NewApp()
	app.Name = "Crypto coin value checker"
	app.Usage = "Tool to check cryptcurrencies prices against coinmarketcap api."

	c := NewCache()
	app.Action = printPrice(c)

	app.Run(os.Args)

	fmt.Println(au.Green("Hello"))
}

func printPrice(c *Cache) func(c *cli.Context) error {
	httpClient := &http.Client{}
	ctClient := NewClient(httpClient)
	ticker := time.NewTicker(duration)

	return printWithInterval(ticker, ctClient, c)
}

func printWithInterval(ticker *time.Ticker, ctClient *Client, c *Cache) func(c *cli.Context) error {
	printCurrent(ctClient, time.Now(), c)

	return func(_ *cli.Context) error {
		for t := range ticker.C {
			printCurrent(ctClient, t, c)
		}
		return nil
	}
}

func printCurrent(ctClient *Client, t time.Time, c *Cache) (int, error) {
	btc, eth, ratio := generateResult(ctClient, c, t)
	return fmt.Printf("%s BTC: %s, ETH: %s, ratio %f \n", t.Format(time.Kitchen), btc, eth, ratio)
}

func generateResult(ctClient *Client, c *Cache, t time.Time) (btc, eth string, ratio float64) {
	btc, _ = ctClient.GetBitcoinPrice()
	eth, _ = ctClient.GetEtherPrice()

	r := calculateRatio(btc, eth)
	le := c.GetLast()
	c.AddEntry(btc, eth, Round(r, .5, 6), t.UTC())

	// When we will have coloring func we will call it from here.
	if r > le.ratio {
		fmt.Printf("%s BTC: %s, ETH: %s, ratio %f \n", t.Format(timeFormat), btc, eth, au.Green(r))
	} else {
		fmt.Printf("%s BTC: %s, ETH: %s, ratio %f \n", t.Format(timeFormat), btc, eth, au.Red(r))
	}

	return btc, eth, r
}

func calculateRatio(bitcoinPrice string, ethereumPrice string) float64 {
	btcPrice, err := strconv.ParseFloat(bitcoinPrice, 64)
	if err != nil {
		return 0
	}

	etherPrice, err := strconv.ParseFloat(ethereumPrice, 64)
	if err != nil {
		return 0
	}

	return etherPrice / btcPrice
}
