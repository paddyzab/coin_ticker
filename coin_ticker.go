package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/urfave/cli"
	"time"
)

func main() {

	var coinToken string

	httpClient := &http.Client{}
	ctClient := NewClient(httpClient)

	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "coin",
			Value:       "bitcoin",
			Usage:       "coin used for testing",
			Destination: &coinToken,
		},
	}

	ticker := time.NewTicker(time.Second * 15)

	app.Name = "bitcoin checker"
	app.Usage = "wrapper testing current cryptcurr price"

	app.Action = func(c *cli.Context) error {
		if coinToken == "b" {
			for t := range ticker.C {
				fmt.Println(t, ctClient.GetBitcoinPrice())
			}

		} else if coinToken == "e" {
			for t := range ticker.C {
				fmt.Println(t, ctClient.GetEtherPrice())
			}
		} else {
			fmt.Println("unkown option")
		}

		return nil
	}

	app.Run(os.Args)
}
