package main

import (
	"net/http"
	"os"
	"time"

	cmcap "github.com/paddyzab/coin_ticker/pkg/coinmarketcap"
	"github.com/paddyzab/coin_ticker/pkg/cointicker"
	"github.com/paddyzab/coin_ticker/pkg/storage"
	"github.com/urfave/cli"
)

const (
	duration = time.Minute * 2
	timeout  = 3 * time.Second
)

func main() {

	app := cli.NewApp()
	app.Name = "Crypto coin value checker"
	app.Usage = "Tool to check cryptcurrencies prices against coinmarketcap api."

	client := cmcap.NewClient(&http.Client{Timeout: timeout})
	cache := storage.NewCache()
	ct := cointicker.NewCoinTicker(client, cache)

	app.Action = func(c *cli.Context) error {
		fetchAndDisplay(c, ct)
		for range time.NewTicker(duration).C {
			fetchAndDisplay(c, ct)
		}
		return nil
	}

	app.Run(os.Args)
}

func fetchAndDisplay(c *cli.Context, ct cointicker.CoinTicker) {
	str, err := ct.GetFormattedPrice()
	if err != nil {
		c.App.Writer.Write([]byte(err.Error()))
	}
	c.App.Writer.Write([]byte(str))
}
