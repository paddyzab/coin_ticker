package main

import (
	"net/http"
	"os"

	"github.com/urfave/cli"
	"fmt"
	"time"
	"strconv"
)

func main() {
	app := cli.NewApp()
	app.Name = "Crypto coin value checker"
	app.Usage = "Tool to check cryptcurrencies prices against coinmarketcap api."

	app.Action = printPrice()

	app.Run(os.Args)
}

func printPrice() func(c* cli.Context) error {
	httpClient := &http.Client{}
	ctClient := NewClient(httpClient)
	ticker := time.NewTicker(time.Second * 60)

	return printWithInterval(ticker, ctClient)
}

func printWithInterval(ticker* time.Ticker, ctClient* Client) func(c* cli.Context) error {
	printCurrent(ctClient)

	return func(c *cli.Context) error {
		for t := range ticker.C {
			btc := ctClient.GetBitcoinPrice()
			eth := ctClient.GetEtherPrice()

			fmt.Printf("%s BTC: %s, ETH: %s, ratio: %d \n", t.Format(time.Kitchen), btc, eth, calculateRatio(btc, eth))
		}
		return nil
	}
}
func printCurrent(ctClient *Client) (int, error) {
	btc := ctClient.GetBitcoinPrice()
	eth := ctClient.GetEtherPrice()

	return fmt.Printf("%s BTC: %s, ETH: %s, ratio %f \n", time.Now().Format(time.Kitchen), ctClient.GetBitcoinPrice(), ctClient.GetEtherPrice(), calculateRatio(btc, eth))
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
