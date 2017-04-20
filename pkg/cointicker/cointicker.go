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
	timeFormat   = time.Kitchen
	formatString = "%s BTC: %s, ETH: %s, XMR: %s \nB/E ratio %f, B/M ratio %f \n\n"
)

// CoinTicker is the base struct for accessing the functionality of the CoinTicker package
type CoinTicker struct {
	Client *cmcap.CoinMarketClient
	Cache  *storage.Cache
	au     aurora.Aurora
}

// NewCoinTicker creates a CoinTicker struct with the passed in API client and storage
func NewCoinTicker(client *cmcap.CoinMarketClient, cache *storage.Cache) CoinTicker {
	return CoinTicker{
		Client: client,
		Cache:  cache,
		au:     aurora.NewAurora(true),
	}
}

// GetFormattedPrice returns the formatted string to display
func (c CoinTicker) GetFormattedPrice(t time.Time) (string, []error) {

	btc, eth, xmr, ethRatio, xmrRatio, errors := c.generateResult()
	if len(errors) != 0 {
		return "", errors
	}

	lastEntry := c.Cache.GetLast()
	c.Cache.AddEntry(btc, eth, xmr, float.Round(ethRatio, .5, 6), float.Round(xmrRatio, .5, 6), t.UTC())

	return fmt.Sprintf(formatString, t.Format(timeFormat), btc, eth, xmr,
		decorateRatio(ethRatio, lastEntry.ETHRatio, c.au)(ethRatio),
		decorateRatio(xmrRatio, lastEntry.XMRRatio, c.au)(xmrRatio)), nil
}

func decorateRatio(ratio, lastRatio float64, au aurora.Aurora) func(interface{}) aurora.Value {
	if ratio > lastRatio {
		return au.Green
	}

	return au.Red
}

func (c CoinTicker) generateResult() (btc, eth, mnr string, ethRatio, mnrRatio float64, errors []error) {

	coins, errors := c.Client.GetCurrenciesQuotes(cmcap.Bitcoin, cmcap.Ether, cmcap.Monero)
	if len(errors) != 0 {
		return
	}

	for i := range coins {
		switch coins[i].ID {
		case cmcap.Bitcoin:
			btc = coins[i].PriceUsd
		case cmcap.Ether:
			eth = coins[i].PriceUsd
		case cmcap.Monero:
			mnr = coins[i].PriceUsd
		}
	}

	ethRatio = calculateRatio(btc, eth)
	mnrRatio = calculateRatio(btc, mnr)
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
