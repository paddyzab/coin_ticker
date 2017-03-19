package coinmarketcap

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sort"
	"sync"
)

const (
	// Bitcoin is the constant for the Bitcoin currency
	Bitcoin = "bitcoin"
	// Ether is the constant for the Ethereum currency
	Ether   = "ethereum"
	baseURL = "https://api.coinmarketcap.com/v1/ticker/"
)

// CoinMarketClient is the client for the coinmarket API
type CoinMarketClient struct {
	httpClient *http.Client
}

// NewClient Creates new configured Client
func NewClient(httpClient *http.Client) *CoinMarketClient {
	return &CoinMarketClient{
		httpClient: httpClient,
	}
}

// Coin represents data returned from the coinmarketcap API
type Coin struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Symbol           string `json:"symbol"`
	Rank             string `json:"rank"`
	PriceUsd         string `json:"price_usd"`
	PriceBtc         string `json:"price_btc"`
	VolumeUsd24h     string `json:"24h_volume_usd"`
	MarketCapUsd     string `json:"market_cap_usd"`
	TotalSupply      string `json:"total_supply"`
	PercentChange1h  string `json:"percent_change_1h"`
	PercentChange24h string `json:"percent_change_24h"`
	PercentChange7d  string `json:"percent_change_7d"`
	LastUpdated      string `json:"last_updated"`
}

// Coins represents a collection of Coin
type Coins []Coin

// Len to implement the Sort interface
func (c Coins) Len() int {
	return len(c)
}

// Less to implement the Sort interface
func (c Coins) Less(i, j int) bool {
	return c[i].Name < c[j].Name
}

// Swap to implement the Sort interface
func (c Coins) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

// GetCurrenciesQuotes fetches the currencies' quotes
func (c *CoinMarketClient) GetCurrenciesQuotes(currencies ...string) ([]Coin, []error) {
	if len(currencies) == 0 {
		return nil, []error{errors.New("no currencies selected")}
	}

	if len(currencies) == 1 {
		coin, err := c.getCurrencyQuote(currencies[0])
		if err != nil {
			return nil, []error{err}
		}
		return []Coin{coin}, nil
	}

	var wg sync.WaitGroup
	values := make(chan Coin, len(currencies))
	errs := make(chan error, len(currencies))

	for _, currency := range currencies {
		wg.Add(1)
		go func(curr string) {
			defer wg.Done()
			coin, err := c.getCurrencyQuote(curr)
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

	return nil, nil
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
