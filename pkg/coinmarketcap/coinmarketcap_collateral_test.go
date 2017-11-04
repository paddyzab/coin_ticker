package coinmarketcap

var (
	etherResponse = `[{
				"id": "ethereum",
				"name": "Ethereum",
				"symbol": "ETH",
				"rank": "2",
				"price_usd": "7",
				"price_btc": "0.1",
				"24h_volume_usd": "220",
				"market_cap_usd": "420",
				"available_supply": "86695896.0",
				"total_supply": "800",
				"percent_change_1h": "0.2",
				"percent_change_24h": "7.93",
				"percent_change_7d": "-8.13",
				"last_updated": "1481134760"
			}]`

	etherCoin = Coin{
		ID:               "ethereum",
		Name:             "Ethereum",
		Symbol:           "ETH",
		Rank:             "2",
		PriceUsd:         "7",
		PriceBtc:         "0.1",
		VolumeUsd24h:     "220",
		MarketCapUsd:     "420",
		TotalSupply:      "800",
		PercentChange1h:  "0.2",
		PercentChange24h: "7.93",
		PercentChange7d:  "-8.13",
		LastUpdated:      "1481134760",
	}

	bitcoinResponse = `[{
				"id": "bitcoin",
				"name": "Bitcoin",
				"symbol": "BTC",
				"rank": "1",
				"price_usd": "600",
				"price_btc": "1.0",
				"24h_volume_usd": "220",
				"market_cap_usd": "420",
				"available_supply": "86695896.0",
				"total_supply": "800",
				"percent_change_1h": "0.2",
				"percent_change_24h": "7.93",
				"percent_change_7d": "-8.13",
				"last_updated": "1481134760"
			}]`

	bitcoinCoin = Coin{
		ID:               "bitcoin",
		Name:             "Bitcoin",
		Symbol:           "BTC",
		Rank:             "1",
		PriceUsd:         "600",
		PriceBtc:         "1.0",
		VolumeUsd24h:     "220",
		MarketCapUsd:     "420",
		TotalSupply:      "800",
		PercentChange1h:  "0.2",
		PercentChange24h: "7.93",
		PercentChange7d:  "-8.13",
		LastUpdated:      "1481134760",
	}

	moneroResponse = `[{
			"id": "monero", 
			"name": "Monero", 
			"symbol": "XMR", 
			"rank": "6", 
			"price_usd": "21.1549", 
			"price_btc": "0.0196459", 
			"24h_volume_usd": "8139100.0", 
			"market_cap_usd": "300852142.0", 
			"available_supply": "14221393.0", 
			"total_supply": "14221393.0", 
			"percent_change_1h": "0.24", 
			"percent_change_24h": "6.71", 
			"percent_change_7d": "4.67", 
			"last_updated": "1491112149"
		}]`

	litecoinResponse = `[{
        "id": "litecoin", 
        "name": "Litecoin", 
        "symbol": "LTC", 
        "rank": "5", 
        "price_usd": "54.2022", 
        "price_btc": "0.00766044", 
        "24h_volume_usd": "198344000.0", 
        "market_cap_usd": "2908021603.0", 
        "available_supply": "53651357.0", 
        "total_supply": "53651357.0", 
        "percent_change_1h": "-0.44", 
        "percent_change_24h": "-0.54", 
        "percent_change_7d": "-2.89", 
        "last_updated": "1509658442"
    }]`

	moneroCoin = Coin{
		ID:               "monero",
		Name:             "Monero",
		Symbol:           "XMR",
		Rank:             "6",
		PriceUsd:         "21.1549",
		PriceBtc:         "0.0196459",
		VolumeUsd24h:     "8139100.0",
		MarketCapUsd:     "300852142.0",
		TotalSupply:      "14221393.0",
		PercentChange1h:  "0.24",
		PercentChange24h: "6.71",
		PercentChange7d:  "4.67",
		LastUpdated:      "1491112149",
	}

	litecoinCoin = Coin{
		ID:               "litecoin",
		Name:             "Litecoin",
		Symbol:           "LTC",
		Rank:             "5",
		PriceUsd:         "54.2022",
		PriceBtc:         "0.00766044",
		VolumeUsd24h:     "198344000.0",
		MarketCapUsd:     "2908021603.0",
		TotalSupply:      "53651357.0",
		PercentChange1h:  "-0.44",
		PercentChange24h: "-0.54",
		PercentChange7d:  "-2.89",
		LastUpdated:      "1509658442",
	}
)
