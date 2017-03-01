package main

import (
	"testing"
	"time"
)

func TestLastEntry(t *testing.T) {
	tc := NewCache()
	curt := time.Now().UTC()

	tc.AddEntry("100", "10", 0.1, curt)
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

	if !(lv.timestamp == curt) {
		t.Error("Timestamp is not matching expected value.", lv.timestamp)
	}
}

func TestLastEntryTwoValues(t *testing.T) {
	tc := NewCache()
	lvt := time.Now().UTC()
	cvt := time.Now().UTC()

	tc.AddEntry("100", "10", 0.1, lvt)
	tc.AddEntry("110", "9", 0.08, cvt)

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

	if !(lv.timestamp == cvt) {
		t.Error("Timestamp is not matching expected value.", lv.timestamp)
	}
}

func TestCacheHasExpectedSize(t *testing.T) {
	tc := NewCache()
	lvt := time.Now().UTC()

	tc.AddEntry("100", "10", 0.1, lvt)
	tc.AddEntry("110", "9", 0.08, lvt)

	s := tc.Size()

	if !(s == 2) {
		t.Error("Size does not match expectations.", s)
	}
}

func TestClearsCache(t *testing.T) {
	tc := NewCache()
	lvt := time.Now().UTC()

	tc.AddEntry("100", "10", 0.1, lvt)
	tc.AddEntry("110", "9", 0.08, lvt)

	tc.Clear()

	s := tc.Size()

	if !(s == 0) {
		t.Error("Size after clearing does not match expectations.", s)
	}
}
