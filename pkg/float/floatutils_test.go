package float

import "testing"

func TestRoundsFloat(t *testing.T) {
	r := Round(10.76868)
	if r != 10.77 {
		t.Error("Rounding was wrong.", r)
	}

	r2 := Round(10.00123)
	if r2 != 10 {
		t.Error("Rounding was wrong.", r2)
	}

	r3 := Round(10.69)
	if r3 != 10.69 {
		t.Error("Rounding was wrong.", r3)
	}

}
