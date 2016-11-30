package main

import (
	"fmt"
	"os"

	"github.com/dghubble/sling"
	"github.com/urfave/cli"
)

// Coin represents data resturned from the coinmarketcap API
type Coin struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Symbol           string `json:"symbol"`
	Rank             int    `json:"rank"`
	PriceUsd         string `json:"price_usd"`
	PriceBtc         string `json:"price_btc"`
	VolumeUsd24h     string `json:"24h_volume_usd"`
	MarketCapUsd     string `json:"market_cap_usd"`
	TotalSupply      string `json:"total_supply"`
	PercentChange1h  string `json:"percent_change_1h"`
	PercentChange24h string `json:"percent_change_24h"`
	PercentChange7d  string `json:"percent_change_7d"`
	LastUpdated      string `json:"last_updated"`
}

const baseURL = "https://api.coinmarketcap.com/v1/ticker/"

func main() {

	var coinToken string

	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "coin",
			Value:       "bitcoin",
			Usage:       "coin used for testing",
			Destination: &coinToken,
		},
	}

	app.Name = "bitcoin checker"
	app.Usage = "wrapper testing current cryptcurr price"

	app.Action = func(c *cli.Context) error {
		if coinToken == "b" {
			callBitcoin()
		} else if coinToken == "e" {
			callEther()
		} else {
			fmt.Println("unkown option")
		}

		return nil
	}

	app.Run(os.Args)
}

func callEther() {
	res := new([]Coin)
	resp, err := sling.New().Get(baseURL).Path("ethereum").ReceiveSuccess(res)
	if err != nil {
		fmt.Println(err, resp)
	} else {
		fmt.Println((*res)[0].PriceUsd)
	}
}

func callBitcoin() {
	res := new([]Coin)
	resp, err := sling.New().Get(baseURL).Path("bitcoin").ReceiveSuccess(res)
	if err != nil {
		fmt.Println(err, resp)
	} else {
		fmt.Println((*res)[0].PriceUsd)
	}
}
