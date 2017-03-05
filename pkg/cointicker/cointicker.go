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

func (c CoinTicker) GetFormattedPrice(t time.Time) (string, error) {

	btc, eth, ratio, errors := c.generateResult()

	if len(errors) != 0 {
		return "", errors[0]
	}

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

func (c CoinTicker) generateResult() (btc, eth string, ratio float64, errors []error) {

	coins, errors := c.Client.GetCurrenciesQuotes([]string{cmcap.Bitcoin, cmcap.Ether}...)
	if len(errors) != 0 {
		return
	}

	for i := range coins {
		switch coins[i].ID {
		case cmcap.Bitcoin:
			btc = coins[i].PriceUsd
		case cmcap.Ether:
			eth = coins[i].PriceUsd
		}
	}
	ratio = calculateRatio(btc, eth)
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
