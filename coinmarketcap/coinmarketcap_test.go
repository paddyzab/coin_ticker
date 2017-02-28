package coinmarketcap

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	relativePath = "/v1/ticker/"
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
		expected        []Coin
		errorString     string
	}{
		{
			title:           "Success request",
			handlerPatterns: []string{relativePath + ether, relativePath + bitcoin},
			handlerFunc: []func(http.ResponseWriter, *http.Request){
				func(w http.ResponseWriter, _ *http.Request) {
					w.Header().Set("Content-Type", "application/json")
					fmt.Fprint(w, etherResponse)
				},
				func(w http.ResponseWriter, _ *http.Request) {
					w.Header().Set("Content-Type", "application/json")
					fmt.Fprint(w, bitcoinResponse)
				},
			},
			request:  []string{ether, bitcoin},
			expected: []Coin{bitcoinCoin, etherCoin},
		},
		{
			title:           "One success - One timeout",
			handlerPatterns: []string{relativePath + ether, relativePath + bitcoin},
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
			},
			request:     []string{ether, bitcoin},
			expected:    []Coin{bitcoinCoin},
			errorString: "Get http://api.coinmarketcap.com/v1/ticker/ethereum: net/http: request canceled (Client.Timeout exceeded while awaiting headers)",
		},
	} {
		t.Run(testCase.title, func(t *testing.T) {
			httpClient, mux, server := testServer()
			for i := range testCase.handlerPatterns {
				mux.HandleFunc(testCase.handlerPatterns[i], testCase.handlerFunc[i])
			}
			defer server.Close()

			client := NewClient(httpClient)
			coins, errs := client.GetCurrenciesQuotes([]string{ether, bitcoin}...)

			assert.Equal(t, testCase.expected, coins)
			if testCase.errorString != "" || errs != nil {
				assert.EqualError(t, errs[0], testCase.errorString)
			}
		})
	}
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
				w.Header().Set("Content-Type", "application/json")
				fmt.Fprint(w, bitcoinResponse)
			},
			expected: "600",
		},
		{
			title:    "Successfull fetch ethereum",
			currency: ether,
			handlerFunc: func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				fmt.Fprint(w, etherResponse)
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
				fmt.Fprint(w, transportItem)
			},
			errorString: "json: cannot unmarshal object into Go value of type []coinmarketcap.Coin",
		},
		{
			title:    "Timeout",
			currency: bitcoin,
			handlerFunc: func(w http.ResponseWriter, r *http.Request) {
				time.Sleep(300 * time.Millisecond)
				w.Header().Set("Content-Type", "application/json")
				fmt.Fprint(w, bitcoinResponse)
			},
			expected:    "",
			errorString: "Get http://api.coinmarketcap.com/v1/ticker/bitcoin: net/http: request canceled (Client.Timeout exceeded while awaiting headers)",
		},
	} {
		t.Run(testCase.title, func(t *testing.T) {
			httpClient, mux, server := testServer()
			mux.HandleFunc("/v1/ticker/"+testCase.currency, testCase.handlerFunc)
			defer server.Close()

			client := NewClient(httpClient)
			resp, err := client.getCurrencyPrice(testCase.currency)

			assert.Equal(t, testCase.expected, resp)
			if testCase.errorString != "" || err != nil {
				assert.EqualError(t, err, testCase.errorString)
			}
		})
	}
}
