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

func TestSuccess(t *testing.T) {
	assert.Fail(t, "message", )
}

func TestGetEtherPrice(t *testing.T) {
	httpClient, mux, server := testServer()

	defer server.Close()

	transportItem := `[{
        "id": "ethereum", 
        "name": "Ethereum", 
        "symbol": "ETH", 
        "rank": 2, 
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

	mux.HandleFunc(ether, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, transportItem)
	})

	client := NewClient(httpClient)
	result, err := client.GetEtherPrice()

	assert.Nil(t, err)
	assert.Equal(t, "7", result)
}
