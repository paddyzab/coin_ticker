package coin_ticker

import "testing"

func TestRoundsFloat(t *testing.T) {
	r := Round(10.001, 5, 2)
	if r != 10 {
		t.Error("Rounding was wrong.", r)
	}

	r2 := Round(10.00123, 5, 3)
	if r2 != 10.001 {
		t.Error("Rounding was wrong.", r2)
	}

	r3 := Round(10.69, 5, 1)
	if r3 != 10.6 {
		t.Error("Rounding was wrong.", r3)
	}

}
