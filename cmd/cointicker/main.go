package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/logrusorgru/aurora"
	"github.com/urfave/cli"

	. "github.com/paddyzab/coin_ticker"
	cmcap "github.com/paddyzab/coin_ticker/coinmarketcap"
)

const (
	duration   = time.Second * 120
	timeout    = 3 * time.Second
	timeFormat = time.Kitchen
)

// output colorizer
var au aurora.Aurora

func init() {
	au = aurora.NewAurora(true)
}

func main() {

	app := cli.NewApp()
	app.Name = "Crypto coin value checker"
	app.Usage = "Tool to check cryptcurrencies prices against coinmarketcap api."

	c := NewCache()
	app.Action = printPrice(c)

	app.Run(os.Args)
}

func printPrice(c *Cache) func(c *cli.Context) error {
	httpClient := &http.Client{Timeout: timeout}
	ctClient := cmcap.NewClient(httpClient)
	ticker := time.NewTicker(duration)

	return printWithInterval(ticker, ctClient, c)
}

func printWithInterval(ticker *time.Ticker, ctClient *cmcap.CoinMarketClient, c *Cache) func(c *cli.Context) error {
	printCurrent(ctClient, c, time.Now())

	return func(_ *cli.Context) error {
		for t := range ticker.C {
			printCurrent(ctClient, c, t)
		}
		return nil
	}
}

func printCurrent(ctClient *cmcap.CoinMarketClient, c *Cache, t time.Time) {
	btc, eth, ratio := generateResult(ctClient)

	le := c.GetLast()
	c.AddEntry(btc, eth, Round(ratio, .5, 6), t.UTC())

	var f func(interface{}) aurora.Value
	if ratio > le.Ratio {
		f = au.Green
	} else {
		f = au.Red
	}
	fmt.Printf("%s BTC: %s, ETH: %s, ratio %f \n", t.Format(timeFormat), btc, eth, f(ratio))
}

func generateResult(ctClient *cmcap.CoinMarketClient) (btc, eth string, ratio float64) {
	btc, _ = ctClient.GetBitcoinPrice()
	eth, _ = ctClient.GetEtherPrice()
	ratio = calculateRatio(btc, eth)
	return
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
