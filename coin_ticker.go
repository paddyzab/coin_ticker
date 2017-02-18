package main

import (
	"net/http"
	"os"

	"github.com/urfave/cli"
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
