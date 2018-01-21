package coinmarketcap

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/paddyzab/coin_ticker/pkg/parsers"

	"github.com/stretchr/testify/assert"
)

const (
	relativePath = "/v1/ticker/"
)

// Use the client to make requests on the server.
// Register handlers on mux to handle requests.
// Caller must close test server.
func newMockServer() (*http.Client, *http.ServeMux, *httptest.Server) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	transport := &RewriteTransport{&http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			fmt.Printf("url: %s \n", req.URL)
			return url.Parse(server.URL)
		},
	}}

	client := &http.Client{Transport: transport, Timeout: 200 * time.Millisecond}
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

func TestGetCurrencyConcurrently(t *testing.T) {

	for _, testCase := range []struct {
		title           string
		handlerPatterns []string
		handlerFunc     []func(http.ResponseWriter, *http.Request)
		request         []string
		expected        CoinsMap
		errorString     string
	}{
		{
			title: "Success request",
			handlerPatterns: []string{
				relativePath + Ether,
				relativePath + Bitcoin,
				relativePath + Monero},
			handlerFunc: []func(http.ResponseWriter, *http.Request){
				func(w http.ResponseWriter, _ *http.Request) {
					w.Header().Set("Content-Type", "application/json")
					fmt.Fprint(w, etherResponse)
				},
				func(w http.ResponseWriter, _ *http.Request) {
					w.Header().Set("Content-Type", "application/json")
					fmt.Fprint(w, bitcoinResponse)
				},
				func(w http.ResponseWriter, _ *http.Request) {
					w.Header().Set("Content-Type", "application/json")
					fmt.Fprint(w, moneroResponse)
				},
			},
			request: []string{Ether, Bitcoin, Monero},
			expected: CoinsMap{
				"BTC": bitcoinCoin,
				"ETH": etherCoin,
				"XMR": moneroCoin},
		},
		{
			title: "One success - One timeout",
			handlerPatterns: []string{
				relativePath + Ether,
				relativePath + Bitcoin,
				relativePath + Monero},
			handlerFunc: []func(http.ResponseWriter, *http.Request){
				func(w http.ResponseWriter, _ *http.Request) {
					w.Header().Set("Content-Type", "application/json")
					time.Sleep(300 * time.Millisecond)
					fmt.Fprint(w, etherResponse)
				},
				func(w http.ResponseWriter, _ *http.Request) {
					w.Header().Set("Content-Type", "application/json")
					fmt.Fprint(w, bitcoinResponse)
				},
				func(w http.ResponseWriter, _ *http.Request) {
					w.Header().Set("Content-Type", "application/json")
					fmt.Fprint(w, moneroResponse)
				},
			},
			request: []string{Ether, Bitcoin, Monero},
			expected: CoinsMap{
				"BTC": bitcoinCoin,
				"XMR": moneroCoin},
			errorString: "Get http://api.coinmarketcap.com/v1/ticker/ethereum: net/http: request canceled (Client.Timeout exceeded while awaiting headers)",
		},
	} {
		t.Run(testCase.title, func(t *testing.T) {
			httpClient, mux, server := newMockServer()
			for i := range testCase.handlerPatterns {
				mux.HandleFunc(testCase.handlerPatterns[i], testCase.handlerFunc[i])
			}
			defer server.Close()
			c := parsers.Conf{Description: "Description", CoinsSymbols: map[string]float64{"BTC": 1, "ETH": 0, "XMR": 100}}

			client := NewClient(httpClient, c)
			coins, errs := client.GetCurrenciesQuotes()

			assert.Equal(t, testCase.expected, coins)
			if testCase.errorString != "" || errs != nil {
				assert.EqualError(t, errs[0], testCase.errorString)
			}
		})
	}
}

func TestGetCurrencyQuote(t *testing.T) {

	for _, testCase := range []struct {
		title       string
		currency    string
		handlerFunc func(http.ResponseWriter, *http.Request)
		expected    Coin
		errorString string
	}{
		{
			title:    "Successfull fetch bitcoin",
			currency: Bitcoin,
			handlerFunc: func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				fmt.Fprint(w, bitcoinResponse)
			},
			expected: bitcoinCoin,
		},
		{
			title:    "Successfull fetch ethereum",
			currency: Ether,
			handlerFunc: func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				fmt.Fprint(w, etherResponse)
			},
			expected: etherCoin,
		},
		{
			title:    "Successfull fetch monero",
			currency: Monero,
			handlerFunc: func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				fmt.Fprint(w, moneroResponse)
			},
			expected: moneroCoin,
		},
		{
			title:    "No content",
			currency: Bitcoin,
			handlerFunc: func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
			},
			errorString: "EOF",
		},
		{
			title:    "Wrong JSON",
			currency: Bitcoin,
			handlerFunc: func(w http.ResponseWriter, r *http.Request) {
				transportItem := `{
					"percent_change_24h": "7.93",
					"percent_change_7d": "-8.13",
					"last_updated": "1481134760"
					}`
				w.Header().Set("Content-Type", "application/json")
				fmt.Fprint(w, transportItem)
			},
			errorString: "json: cannot unmarshal object into Go value of type []coinmarketcap.Coin",
		},
		{
			title:    "Timeout",
			currency: Bitcoin,
			handlerFunc: func(w http.ResponseWriter, r *http.Request) {
				time.Sleep(300 * time.Millisecond)
				w.Header().Set("Content-Type", "application/json")
				fmt.Fprint(w, bitcoinResponse)
			},
			errorString: "Get http://api.coinmarketcap.com/v1/ticker/bitcoin: net/http: request canceled (Client.Timeout exceeded while awaiting headers)",
		},
	} {
		t.Run(testCase.title, func(t *testing.T) {
			httpClient, mux, server := newMockServer()
			mux.HandleFunc("/v1/ticker/"+testCase.currency, testCase.handlerFunc)
			defer server.Close()
			var c parsers.Conf

			client := NewClient(httpClient, c)
			resp, err := client.getCurrencyQuote(testCase.currency)

			assert.Equal(t, testCase.expected, resp)
			if testCase.errorString != "" || err != nil {
				assert.EqualError(t, err, testCase.errorString)
			}
		})
	}
}
