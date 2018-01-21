package cointicker

import (
	"fmt"
	"strconv"
	"time"
	"bytes"

	"github.com/logrusorgru/aurora"

	cmcap "github.com/paddyzab/coin_ticker/pkg/coinmarketcap"
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

	res, errors := c.generateResult()

	if len(errors) != 0 {
		return "", errors
	}
	c.Cache.AddEntry(res, t.UTC())
	lastEntry := c.Cache.GetLast()

	keys := make([]string, len(res.Result))

	i := 0
	for k := range res.Result {
		keys[i] = k
		i++
	}

	var buffer bytes.Buffer
	buffer.WriteString("\n" + t.Format(timeFormat) + "\n")
	for i := 0; i < len(keys); i++ {
		buffer.WriteString(fmt.Sprintf(keys[i] + ": %f, ", decorateRatio(res.Result[keys[i]], lastEntry.CoinData.Result[keys[i]], c)(res.Result[keys[i]])))
	}

	return buffer.String(), nil
}

func decorateRatio(r, lr float64, c CoinTicker) func(interface{}) aurora.Value {
	if r > lr {
		return c.au.Green
	}

	return c.au.Red
}

func (c CoinTicker) generateResult() (res storage.Results, errors []error) {

	res.Result = make(map[string]float64)
	coinsMap, errors := c.Client.GetCurrenciesQuotes()
	if len(errors) != 0 {
		res.Errors = errors
		return
	}

	for k, v := range coinsMap {
		if k != "BTC" {
			res.Result[k] = calculateRatio(coinsMap["BTC"].PriceUsd, v.PriceUsd)
		}
	}

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
