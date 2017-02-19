package main

import (
	"net/http"
	"os"

	"github.com/urfave/cli"
	"fmt"
	"time"
	"strconv"
)

type result struct {
	bitcoin string
	ether string
	ratio float64
}

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
			res := generateResult(ctClient)
			fmt.Printf("%s BTC: %s, ETH: %s, ratio: %d \n", t.Format(time.Kitchen), res.bitcoin, res.ether, res.ratio)
		}
		return nil
	}
}
func printCurrent(ctClient *Client) (int, error) {
	res := generateResult(ctClient)
	return fmt.Printf("%s BTC: %s, ETH: %s, ratio %f \n", time.Now().Format(time.Kitchen), res.bitcoin, res.ether, res.ratio)
}
func generateResult(ctClient *Client) result {
	btc := ctClient.GetBitcoinPrice()
	eth := ctClient.GetEtherPrice()

	return result{bitcoin:btc, ether:eth, ratio:calculateRatio(btc, eth)}
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
