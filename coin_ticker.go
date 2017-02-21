package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Crypto coin value checker"
	app.Usage = "Tool to check cryptcurrencies prices against coinmarketcap api."

	app.Action = printPrice()

	app.Run(os.Args)
}

func printPrice() func(c *cli.Context) error {
	httpClient := &http.Client{}
	ctClient := NewClient(httpClient)
	ticker := time.NewTicker(time.Second * 60)

	return printWithInterval(ticker, ctClient)
}

func printWithInterval(ticker *time.Ticker, ctClient *Client) func(c *cli.Context) error {
	printCurrent(ctClient, time.Now())

	return func(_ *cli.Context) error {
		for t := range ticker.C {
			printCurrent(ctClient, t)
		}
		return nil
	}
}

func printCurrent(ctClient *Client, t time.Time) (int, error) {
	btc, eth, ratio := generateResult(ctClient)
	return fmt.Printf("%s BTC: %s, ETH: %s, ratio %f \n", t.Format(time.Kitchen), btc, eth, ratio)
}

func generateResult(ctClient *Client) (btc, eth string, ratio float64) {
	btc, _ = ctClient.GetBitcoinPrice()
	eth, _ = ctClient.GetEtherPrice()

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
