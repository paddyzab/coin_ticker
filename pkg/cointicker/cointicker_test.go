package cointicker

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/logrusorgru/aurora"
	"github.com/paddyzab/coin_ticker/pkg/coinmarketcap"

	"strings"
	"time"

	"github.com/paddyzab/coin_ticker/pkg/storage"
)

const testString = "test"

type coinMarketCapMock struct {
	coins  []coinmarketcap.Coin
	errors []error
}

func (c coinMarketCapMock) GetCurrenciesQuotes(currencies ...string) ([]coinmarketcap.Coin, []error) {
	return c.coins, c.errors
}

func TestGetFormattedPrice(t *testing.T) {

	cases := []struct {
		title  string
		coins  []coinmarketcap.Coin
		errors []error
		time   time.Time
		expRes string
		expErr []error
	}{
		{
			title: "Bitcoin",
			coins: []coinmarketcap.Coin{coinmarketcap.Coin{
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
			}},
			errors: nil,
			time:   time.Date(2017, 12, 01, 10, 15, 0, 0, time.UTC),
			expRes: fmt.Sprintf("10:15AM BTC: 600, ETH: , XMR: %s \nB/E ratio %f, B/M ratio %f \n\n", "", 0.0, 0.0),
				expErr: nil,
			//%s BTC: %s, ETH: %s, XMR: %s \nB/E ratio %f, B/M ratio %f \n\n
		},
	}

	for _, tc := range cases {
		t.Run(tc.title, func(t *testing.T) {
			ct := NewCoinTicker(&coinMarketCapMock{coins: tc.coins, errors: tc.errors}, storage.NewCache())
			res, err := ct.GetFormattedPrice(tc.time)
			fmt.Printf("<%s>\n", res)


			res = strings.Replace(res, "\n", "", -1)
			res = strings.Replace(res, " ", "", -1)
			fmt.Printf("<%s>\n", res)

	//res = "house"

			if fmt.Sprintf("%s", tc.expRes) != fmt.Sprintf("%s", res) {
				fmt.Printf("res <%s>\n", res)

				for _, runeValue := range []rune(res) {
					fmt.Printf("%#U \n", runeValue)
				}

				fmt.Printf("exp <%s>\n", tc.expRes)

				for _, runeValue := range []rune(tc.expRes) {
					fmt.Printf("%#U \n", runeValue)
				}
				t.Error(strings.Replace(res, "\n", " ", 100))
			}

			if !reflect.DeepEqual(err, tc.expErr) {
				t.Error(err)

			}
		})
	}
}

func TestDecorateRatio(t *testing.T) {

	testCases := []struct {
		name      string
		ratio     float64
		lastRatio float64
		exp       func(interface{}) aurora.Value
	}{
		{
			name:      "Stock going up",
			ratio:     1.0,
			lastRatio: 0.9,
			exp:       aurora.Green,
		},
		{
			name:      "Stock flat",
			ratio:     1.0,
			lastRatio: 1.0,
			exp:       aurora.Red,
		},
		{
			name:      "Stock going down",
			ratio:     1.0,
			lastRatio: 1.2,
			exp:       aurora.Red,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			au := aurora.NewAurora(true)
			got := decorateRatio(tc.ratio, tc.lastRatio, au)

			if !reflect.DeepEqual(got(testString), tc.exp(testString)) {
				t.Errorf("got unexpected value. exp: %v, got: %v", tc.exp(testString), got(testString))
			}
		})
	}

}

func TestCalculateRatio(t *testing.T) {

	testCases := []struct {
		name  string
		base  string
		value string
		exp   float64
	}{
		{
			name:  "Smaller than",
			base:  "1.0",
			value: "0.9",
			exp:   0.9,
		},
		{
			name:  "Equal",
			base:  "1.0",
			value: "1.0",
			exp:   1.0,
		},
		{
			name:  "Bigger than",
			base:  "2.0",
			value: "3.2",
			exp:   1.6,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := calculateRatio(tc.base, tc.value)

			if tc.exp != got {
				t.Errorf("got unexpected value. exp: %v, got: %v", tc.exp, got)
			}
		})
	}

}
