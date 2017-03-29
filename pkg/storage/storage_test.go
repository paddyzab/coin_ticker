package storage

import (
	"testing"
	"time"
)

func TestCacheLastEntry(t *testing.T) {

	for _, tc := range []struct {
		title     string
		bitcoins  []string
		ethereums []string
		moneros   []string
		bers      []float64
		mers      []float64
		times     []time.Time
		expected  []Entry
	}{
		{
			bitcoins:  []string{"100"},
			ethereums: []string{"10"},
			moneros:   []string{"60"},
			bers:      []float64{0.1},
			mers:      []float64{2.3},
			times:     []time.Time{time.Now()},
			title:     "One addition",
		},
		{
			bitcoins:  []string{"100", "10"},
			ethereums: []string{"10", "5"},
			moneros:   []string{"60", "1"},
			bers:      []float64{0.1, 1.5},
			mers:      []float64{2.3, 5.2},
			times:     []time.Time{time.Now(), time.Now().Add(time.Hour)},
			title:     "Two additions",
		},
	} {
		t.Run(tc.title, func(t *testing.T) {
			cache := NewCache()

			for i := range tc.bitcoins {
				cache.AddEntry(tc.bitcoins[i], tc.ethereums[i], tc.moneros[i], tc.bers[i], tc.mers[i], tc.times[i])
			}

			if cache.Size() != len(tc.bitcoins) {
				t.Errorf("invalid size of the cache - want: %v got: %v", len(tc.bitcoins), cache.Size())
			}

			lastValue := cache.GetLast()

			if lastValue.BitcoinPrice != tc.bitcoins[len(tc.bitcoins)-1] {
				t.Errorf("bitcoin price does not match - want: %v got: %v", tc.bitcoins[0], lastValue.BitcoinPrice)
			}

			if lastValue.EtherPrice != tc.ethereums[len(tc.ethereums)-1] {
				t.Errorf("ethereum price does not match - want: %v got: %v", tc.ethereums[0], lastValue.EtherPrice)
			}

			if lastValue.MoneroPrice != tc.moneros[len(tc.moneros)-1] {
				t.Errorf("monero price does not match - want: %v got: %v", tc.moneros[0], lastValue.MoneroPrice)
			}

			if lastValue.ETHRatio != tc.bers[len(tc.bers)-1] {
				t.Errorf("ethereum/bitcoin ratio does not match - want: %v got: %v", tc.bers[0], lastValue.ETHRatio)
			}

			if lastValue.XMRRatio != tc.mers[len(tc.mers)-1] {
				t.Errorf("monero/bitcoin ratio does not match - want: %v got: %v", tc.mers[0], lastValue.XMRRatio)
			}

			if lastValue.Timestamp != tc.times[len(tc.times)-1] {
				t.Errorf("timestamp does not match - want: %v got: %v", tc.times[0], lastValue.Timestamp)
			}
		})
	}
}

func TestClearsCache(t *testing.T) {
	cache := NewCache()

	cache.AddEntry("100", "10", "22", 0.01, 0.1, time.Now())
	cache.AddEntry("110", "9", "12", 9.99, 0.08, time.Now())

	if cache.Size() != 2 {
		t.Errorf("invalid size of the cache - want: %v got: %v", 2, cache.Size())
	}

	cache.Clear()

	if cache.Size() != 0 {
		t.Errorf("invalid size of the cache - want: %v got: %v", 0, cache.Size())
	}
}
