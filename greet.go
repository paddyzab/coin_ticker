package main

import (
	"fmt"
	"os"

	"github.com/bndr/gopencils"
	"github.com/urfave/cli"
)

type coin struct {
	id                 string
	name               string
	symbol             string
	price_usd          string
	price_btc          string
	percent_change_24h string
	last_updated       string
}

type coins struct {
	data []coin
}

func main() {

	var coin string

	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "coin",
			Value:       "bitcoin",
			Usage:       "coin used for testing",
			Destination: &coin,
		},
	}

	app.Name = "bitcoin checker"
	app.Usage = "wrapper testing current cryptcurr price"

	app.Action = func(c *cli.Context) error {
		if coin == "b" {
			callBitcoin()
		} else if coin == "e" {
			callEther()
		} else {
			fmt.Println("unkown option")
		}

		return nil
	}

	app.Run(os.Args)

	api := gopencils.Api("https://api.coinmarketcap.com/v1/ticker/bitcoin")
	c := api.Res()

	res := new(coins)

	_, err := c.Id("bitcoin", res).Get()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}

func callEther() {
	fmt.Println("Hello ether")
}

func callBitcoin() {
	fmt.Println("Hello bitcoin")
}
