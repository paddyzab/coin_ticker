package main

import (
	"net/http"

	"github.com/dghubble/sling"
)

const (
	baseURL = "https://api.coinmarketcap.com/v1/ticker/"
	ether   = "ethereum"
	bitcoin = "bitcoin"
)

// Client CoinTicker api client
type Client struct {
	sling *sling.Sling
}

// NewClient Creates new configured Client
func NewClient(client *http.Client) *Client {
	return &Client{
		sling: sling.New().Client(client).Base(baseURL),
	}
}

// Coin represents data resturned from the coinmarketcap API
type Coin struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Symbol           string `json:"symbol"`
	Rank             int    `json:"rank"`
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

// GetEtherPrice returns current price of Ethereum in USD
func (client *Client) GetEtherPrice() (string, error) {
	coin := new([]Coin)
	_, err := client.sling.New().Get(ether).ReceiveSuccess(&coin)

	if err != nil {
		return (*coin)[0].PriceUsd, err
	}

	return (*coin)[0].PriceUsd, err
}

// GetBitcoinPrice returns current price of Bitcoin in USD
func (client *Client) GetBitcoinPrice() (string, error) {
	coin := new([]Coin)
	_, err := client.sling.New().Get(bitcoin).ReceiveSuccess(&coin)

	if err != nil {
		return (*coin)[0].PriceUsd, err
	}

	return (*coin)[0].PriceUsd, err
}
