package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Use the client to make requests on the server.
// Register handlers on mux to handle requests.
// Caller must close test server.
func testServer() (*http.Client, *http.ServeMux, *httptest.Server) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	transport := &RewriteTransport{&http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(server.URL)
		},
	}}

	client := &http.Client{Transport: transport}
	return client, mux, server
}

type RewriteTransport struct {
	Transport http.RoundTripper
}

func (t *RewriteTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.URL.Scheme = "http"
	if t.Transport == nil {
		return http.DefaultTransport.RoundTrip(req)
	}
	return t.Transport.RoundTrip(req)
}

func TestGetCurrencyPrice(t *testing.T) {

	for _, testCase := range []struct {
		title       string
		currency    string
		handlerFunc func(http.ResponseWriter, *http.Request)
		expected    string
		errorString string
	}{
		{
			title:    "Successfull fetch bitcoin",
			currency: bitcoin,
			handlerFunc: func(w http.ResponseWriter, r *http.Request) {
				transportItem := `[{
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
				w.Header().Set("Content-Type", "application/json")
				fmt.Fprintf(w, transportItem)
			},
			expected: "600",
		},
		{
			title:    "Successfull fetch ethereum",
			currency: ether,
			handlerFunc: func(w http.ResponseWriter, r *http.Request) {
				transportItem := `[{
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
				w.Header().Set("Content-Type", "application/json")
				fmt.Fprintf(w, transportItem)
			},
			expected: "7",
		},
		{
			title:    "No content",
			currency: bitcoin,
			handlerFunc: func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
			},
			errorString: "EOF",
		},
		{
			title:    "Wrong JSON",
			currency: bitcoin,
			handlerFunc: func(w http.ResponseWriter, r *http.Request) {
				transportItem := `{
					"percent_change_24h": "7.93",
					"percent_change_7d": "-8.13",
					"last_updated": "1481134760"
					}`
				w.Header().Set("Content-Type", "application/json")
				fmt.Fprintf(w, transportItem)
			},
			errorString: "json: cannot unmarshal object into Go value of type []main.Coin",
		},
		{
			title:    "Timeout",
			currency: bitcoin,
			handlerFunc: func(w http.ResponseWriter, r *http.Request) {
				transportItem := `[{
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
				//time.Sleep(5 * time.Second)
				w.Header().Set("Content-Type", "application/json")
				fmt.Fprintf(w, transportItem)
			},
			expected: "600",
		},
	} {
		t.Run(testCase.title, func(t *testing.T) {
			httpClient, mux, server := testServer()
			defer server.Close()
			mux.HandleFunc("/v1/ticker/"+testCase.currency, testCase.handlerFunc)
			client := NewClient(httpClient)

			resp, err := client.getCurrencyPrice(testCase.currency)
			assert.Equal(t, testCase.expected, resp)
			if testCase.errorString != "" {
				assert.EqualError(t, err, testCase.errorString)
			}
		})
	}
}
