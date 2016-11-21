package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

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

	//https://api.coinmarketcap.com/v1/ticker/

	app.Run(os.Args)
}

func callEther() {
	fmt.Println("Hello ether")
}

func callBitcoin() {
	fmt.Println("Hello bitcoin")
}
