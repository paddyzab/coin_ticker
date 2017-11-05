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

// CoinTicker contains every information required to generate, store and prepare results.
type CoinTicker struct {
	Client *cmcap.CoinMarketClient
	Cache  *storage.Cache
	au     aurora.Aurora
}

//NewCoinTicker returns new client.
func NewCoinTicker(client *cmcap.CoinMarketClient, cache *storage.Cache) CoinTicker {
	return CoinTicker{
		Client: client,
		Cache:  cache,
		au:     aurora.NewAurora(true),
	}
}

//GetFormattedPrice returns formatted prices ready to be printed.
func (c CoinTicker) GetFormattedPrice(t time.Time) (string, []error) {

	btc, eth, xmr, neo, ethRatio, xmrRatio, neoRatio, errors := c.generateResult()

	if len(errors) != 0 {
		return "", errors
	}

	lastEntry := c.Cache.GetLast()
	c.Cache.AddEntry(btc, eth, xmr, neo, float.Round(ethRatio), float.Round(xmrRatio), float.Round(neoRatio), t.UTC())

	return fmt.Sprintf("%s BTC: %s, ETH: %s, XMR: %s, NEO: %s \nB/E ratio %f, B/M ratio %f, B/N ratio %f \n\n", t.Format(timeFormat), btc, eth, xmr, neo, decorateRatio(ethRatio, lastEntry.ETHRatio, c)(ethRatio), decorateRatio(xmrRatio, lastEntry.XMRRatio, c)(xmrRatio), decorateRatio(neoRatio, lastEntry.NEORatio, c)(neoRatio)), nil
}

func decorateRatio(r, lr float64, c CoinTicker) func(interface{}) aurora.Value {
	if r > lr {
		return c.au.Green
	}

	return c.au.Red
}

// Results ...
type Results struct {
	Result map[string]float64
}

func (c CoinTicker) generateResult() (btc, eth, mnr, neo string, ethRatio, mnrRatio, neoRatio float64, errors []error) {

	coins, errors := c.Client.GetCurrenciesQuotes()
	if len(errors) != 0 {
		return
	}

	fmt.Println(coins)

	// ethRatio = calculateRatio(btc, eth)
	// mnrRatio = calculateRatio(btc, mnr)
	// neoRatio = calculateRatio(btc, neo)
	return
}

func calculateRatio(bitcoinPrice string, coinPrice string) float64 {
	btcPrice, err := strconv.ParseFloat(bitcoinPrice, 64)
	if err != nil {
		return 0
	}

	cPrice, err := strconv.ParseFloat(coinPrice, 64)
	if err != nil {
		return 0
	}

	return cPrice / btcPrice
}
