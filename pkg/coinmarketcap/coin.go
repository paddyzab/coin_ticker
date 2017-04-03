package coinmarketcap

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
