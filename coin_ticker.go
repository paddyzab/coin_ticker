package main

import (
	"net/http"
	"os"

	"github.com/urfave/cli"
	"fmt"
	"time"
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
			fmt.Printf("%s BTC: %s, ETH: %s \n", t.Format(time.Kitchen), ctClient.GetBitcoinPrice(), ctClient.GetEtherPrice())
		}
		return nil
	}
}
func printCurrent(ctClient *Client) (int, error) {
	return fmt.Printf("%s BTC: %s, ETH: %s \n", time.Now().Format(time.Kitchen), ctClient.GetBitcoinPrice(), ctClient.GetEtherPrice())
}
