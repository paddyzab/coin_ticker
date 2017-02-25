package main

import (
	"testing"
)

func TestLastEntry(t *testing.T) {
	tc := New()

	tc.AddEntry("100", "10", 0.1)
	lv := tc.GetLast()

	if !(lv.bitcoinPrice == "100") {
		t.Error("Bitcoin price does not match expected value.", lv.bitcoinPrice)
	}

	if !(lv.etherPrice == "10") {
		t.Error("Ether price does not match expected value.", lv.etherPrice)
	}

	if !(lv.ratio == 0.1) {
		t.Error("Ratio does not match expected value.", lv.ratio)
	}
}

func TestLastEntryTwoValues(t *testing.T) {
	tc := New()

	tc.AddEntry("100", "10", 0.1)
	tc.AddEntry("110", "9", 0.08)

	lv := tc.GetLast()

	if !(lv.bitcoinPrice == "110") {
		t.Error("Bitcoin price does not match expected value.", lv.bitcoinPrice)
	}

	if !(lv.etherPrice == "9") {
		t.Error("Ether price does not match expected value.", lv.etherPrice)
	}

	if !(lv.ratio == 0.08) {
		t.Error("Ratio does not match expected value.", lv.ratio)
	}
}

func TestCacheHasExpectedSize(t *testing.T) {
	tc := New()

	tc.AddEntry("100", "10", 0.1)
	tc.AddEntry("110", "9", 0.08)

	s := tc.Size()

	if !(s == 2) {
		t.Error("Size does not match expectations.", s)
	}
}

func TestClearsCache(t *testing.T) {
	tc := New()

	tc.AddEntry("100", "10", 0.1)
	tc.AddEntry("110", "9", 0.08)

	tc.Clear()

	s:= tc.Size()

	if !(s == 0) {
		t.Error("Size after clearing does not match expectations.", s)
	}
}
