package cointicker

import (
	"testing"

	"reflect"

	"github.com/logrusorgru/aurora"
)

const testString = "test"

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
