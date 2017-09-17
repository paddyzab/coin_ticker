package main

import (
	"net/http"
	"os"
	"time"

	"github.com/paddyzab/coin_ticker/pkg/parsers"

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

	var c parsers.Conf
	c.GetConfiguration()

	app.Action = func(c *cli.Context) error {
		fetchAndDisplay(c, ct, time.Now())
		for t := range time.NewTicker(duration).C {
			fetchAndDisplay(c, ct, t)
		}
		return nil
	}

	app.Run(os.Args)
}

func fetchAndDisplay(c *cli.Context, ct cointicker.CoinTicker, t time.Time) {
	str, errors := ct.GetFormattedPrice(t)
	if len(errors) != 0 {
		for _, err := range errors {
			c.App.Writer.Write([]byte(err.Error()))
			c.App.Writer.Write([]byte{'\n'})
		}
	}
	c.App.Writer.Write([]byte(str))
}
