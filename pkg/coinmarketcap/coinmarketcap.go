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
	Ether = "ethereum"
	// Monero is the constant for the Monero currency
	Monero  = "monero"
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
