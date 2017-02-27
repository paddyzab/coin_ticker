package main

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
		ID: "ethereum",
		Name: "Ethereum",
		Symbol: "ETH",
		Rank: "2",
		PriceUsd: "7",
		PriceBtc: "0.1",
		VolumeUsd24h: "220",
		MarketCapUsd: "420",
		TotalSupply: "800",
		PercentChange1h: "0.2",
		PercentChange24h: "7.93",
		PercentChange7d: "-8.13",
		LastUpdated: "1481134760",
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
		ID: "bitcoin",
		Name: "Bitcoin",
		Symbol: "BTC",
		Rank: "1",
		PriceUsd: "600",
		PriceBtc: "1.0",
		VolumeUsd24h: "220",
		MarketCapUsd: "420",
		TotalSupply: "800",
		PercentChange1h: "0.2",
		PercentChange24h: "7.93",
		PercentChange7d: "-8.13",
		LastUpdated: "1481134760",
	}
)
