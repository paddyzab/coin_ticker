package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/urfave/cli"
)

const (
	duration = time.Second * 120
)

func main() {
	app := cli.NewApp()
	app.Name = "Crypto coin value checker"
	app.Usage = "Tool to check cryptcurrencies prices against coinmarketcap api."

	c := New()
	app.Action = printPrice(c)

	app.Run(os.Args)
}

func printPrice(c *Cache) func(c *cli.Context) error {
	httpClient := &http.Client{}
	ctClient := NewClient(httpClient)
	ticker := time.NewTicker(duration)

	return printWithInterval(ticker, ctClient, c)
}

func printWithInterval(ticker *time.Ticker, ctClient *Client, ch *Cache) func(c *cli.Context) error {
	printCurrent(ctClient, time.Now(), ch)

	return func(_ *cli.Context) error {
		for t := range ticker.C {
			printCurrent(ctClient, t, ch)
		}
		return nil
	}
}

func printCurrent(ctClient *Client, t time.Time, ch *Cache) (int, error) {
	btc, eth, ratio := generateResult(ctClient, t, ch)
	return fmt.Printf("%s BTC: %s, ETH: %s, ratio %f \n", t.Format(time.Kitchen), btc, eth, ratio)
}

func generateResult(ctClient *Client, k time.Time, ch *Cache) (btc, eth string, ratio float64) {
	btc, _ = ctClient.GetBitcoinPrice()
	eth, _ = ctClient.GetEtherPrice()

	ch.AddEntry(strconv.FormatInt(k.Unix(), 10), btc, eth)
	return btc, eth, calculateRatio(btc, eth)
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
