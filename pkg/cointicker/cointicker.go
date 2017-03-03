package cointicker

import (
	"fmt"
	"strconv"
	"time"

	"github.com/logrusorgru/aurora"

	cmcap "github.com/paddyzab/coin_ticker/pkg/coinmarketcap"
	"github.com/paddyzab/coin_ticker/pkg/float"
	"github.com/paddyzab/coin_ticker/pkg/storage"
)

const (
	timeFormat = time.Kitchen
)

type CoinTicker struct {
	Client *cmcap.CoinMarketClient
	Cache  *storage.Cache
	au     aurora.Aurora
}

func NewCoinTicker(client *cmcap.CoinMarketClient, cache *storage.Cache) CoinTicker {
	return CoinTicker{
		Client: client,
		Cache:  cache,
		au:     aurora.NewAurora(true),
	}
}

func (c CoinTicker) GetFormattedPrice() (string, error) {

	btc, eth, ratio, t := c.generateResult()

	lastEntry := c.Cache.GetLast()
	c.Cache.AddEntry(btc, eth, float.Round(ratio, .5, 6), t.UTC())

	var f func(interface{}) aurora.Value
	if ratio > lastEntry.Ratio {
		f = c.au.Green
	} else {
		f = c.au.Red
	}

	return fmt.Sprintf("%s BTC: %s, ETH: %s, ratio %f \n", t.Format(timeFormat), btc, eth, f(ratio)), nil
}

func (c CoinTicker) generateResult() (btc, eth string, ratio float64, t time.Time) {
	btc, _ = c.Client.GetBitcoinPrice()
	eth, _ = c.Client.GetEtherPrice()
	ratio = calculateRatio(btc, eth)
	t = time.Now()
	return
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
