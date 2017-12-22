package storage

import (
	"testing"
	"time"
)

func TestCacheLastEntry(t *testing.T) {

	for _, tc := range []struct {
		title    string
		times    []time.Time
		expected []map[string]float64
	}{
		{
			title: "One addition",
			times: []time.Time{time.Now()},
			expected: []map[string]float64{map[string]float64{
				"BTC": 1000,
				"ETH": 900,
				"XMR": 500,
			}},
		},
		{
			title: "Two additions",
			times: []time.Time{time.Now(), time.Now().Add(time.Hour)},
			expected: []map[string]float64{
				map[string]float64{
					"BTC": 1000,
					"ETH": 900,
					"XMR": 500,
				},
				map[string]float64{
					"BTC": 1100,
					"ETH": 950,
					"XMR": 530,
				},
			},
		},
	} {
		t.Run(tc.title, func(t *testing.T) {
			cache := NewCache()

			for i := range tc.expected {
				cache.AddEntry(Results{tc.expected[i], nil}, tc.times[i])
			}

			if cache.Size() != len(tc.expected) {
				t.Errorf("invalid size of the cache - want: %v got: %v", len(tc.expected), cache.Size())
			}

			lastValue := cache.GetLast()

			if lastValue.CoinData.Result["BTC"] != tc.expected[len(tc.expected)-1]["BTC"] {
				t.Errorf("bitcoin price does not match - want: %v got: %v", tc.expected[0]["BTC"], lastValue.CoinData.Result["BTC"])
			}

			if lastValue.CoinData.Result["ETH"] != tc.expected[len(tc.expected)-1]["ETH"] {
				t.Errorf("ethereum price does not match - want: %v got: %v", tc.expected[0]["ETH"], lastValue.CoinData.Result["ETH"])
			}

			if lastValue.CoinData.Result["XMR"] != tc.expected[len(tc.expected)-1]["XMR"] {
				t.Errorf("monero price does not match - want: %v got: %v", tc.expected[0]["XMR"], lastValue.CoinData.Result["XMR"])
			}

			if lastValue.Timestamp != tc.times[len(tc.times)-1] {
				t.Errorf("timestamp does not match - want: %v got: %v", tc.times[0], lastValue.Timestamp)
			}
		})
	}
}

func TestClearsCache(t *testing.T) {
	cache := NewCache()

	r := map[string]float64{
		"BTC": 1000,
		"ETH": 900,
		"XMR": 500,
		"LTC": 300,
	}

	r2 := map[string]float64{
		"BTC": 1100,
		"ETH": 950,
		"XMR": 530,
		"LTC": 310,
	}

	res := Results{r, nil}
	res2 := Results{r2, nil}

	cache.AddEntry(res, time.Now())
	cache.AddEntry(res2, time.Now())

	if cache.Size() != 2 {
		t.Errorf("invalid size of the cache - want: %v got: %v", 2, cache.Size())
	}

	cache.Clear()

	if cache.Size() != 0 {
		t.Errorf("invalid size of the cache - want: %v got: %v", 0, cache.Size())
	}
}
