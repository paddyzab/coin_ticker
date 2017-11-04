package coinmarketcap

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sort"
	"sync"

	"github.com/paddyzab/coin_ticker/pkg/parsers"
)

const (

	// Bitcoin is a constant for the Bitcoin currency
	Bitcoin = "bitcoin"

	// Ether is a constant for the Ethereum currency
	Ether = "ethereum"

	// Monero is a constant for the Monero currency
	Monero = "monero"

	// Neo is a constant for the NEO token
	Neo = "neo"

	baseURL = "https://api.coinmarketcap.com/v1/ticker/"
)

//Mappings are just that, mappings of the currency symbols to coinmarketcap specific url encodings.
var Mappings = struct {
	Symbols map[string]string
}{
	Symbols: map[string]string{
		"BTC":  "bitcoin",
		"ETH":  "ethereum",
		"XMR":  "monero",
		"NEO":  "neo",
		"DASH": "dash",
		"LTC":  "litecoin",
		"ETC":  "ethereum-classic",
		"BTH":  "bitcoin-cash",
	}}

// CoinMarketClient is the client for the coinmarket API
type CoinMarketClient struct {
	httpClient *http.Client
	config     parsers.Conf
}

// NewClient Creates new configured Client
func NewClient(httpClient *http.Client, conf parsers.Conf) *CoinMarketClient {
	return &CoinMarketClient{
		httpClient: httpClient,
		config:     conf,
	}
}

func getCoinMarketCapCoinID(symbol string) string {
	return Mappings.Symbols[symbol]
}

// GetCurrenciesQuotes fetches the currencies' quotes
func (c *CoinMarketClient) GetCurrenciesQuotes() ([]Coin, []error) {

	curSymbols := make([]string, 0, len(c.config.CoinsSymbols))
	for k := range c.config.CoinsSymbols {
		curSymbols = append(curSymbols, k)
	}

	if len(curSymbols) == 0 {
		return nil, []error{errors.New("no currencies selected")}
	}

	if len(curSymbols) == 1 {
		coin, err := c.getCurrencyQuote(getCoinMarketCapCoinID(curSymbols[0]))
		if err != nil {
			return nil, []error{err}
		}
		return []Coin{coin}, nil
	}

	var wg sync.WaitGroup
	values := make(chan Coin, len(curSymbols))
	errs := make(chan error, len(curSymbols))

	for _, currency := range curSymbols {
		wg.Add(1)
		go func(curr string) {
			defer wg.Done()
			coin, err := c.getCurrencyQuote(getCoinMarketCapCoinID(curr))
			if err != nil {
				errs <- err
				return
			}
			values <- coin
		}(currency)
	}

	wg.Wait()

	var coins Coins
	var err []error
	for {
		select {
		case c := <-values:
			coins = append(coins, c)
		case e := <-errs:
			err = append(err, e)
		default:
			sort.Sort(coins)
			return coins, err
		}
	}
}

func (c *CoinMarketClient) getCurrencyQuote(currency string) (Coin, error) {
	resp, err := c.httpClient.Get(baseURL + currency)
	if err != nil {
		return Coin{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return Coin{}, fmt.Errorf("response from service was not OK: <%v>.\nHeaders: <%v>.\nContent: <%v>",
			resp.StatusCode, resp.Header, resp.Body)
	}

	coins := make([]Coin, 0, 1)
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&coins); err != nil {
		return Coin{}, err
	}

	if len(coins) < 1 {
		return Coin{}, errors.New("no content received")
	}

	return coins[0], nil
}
