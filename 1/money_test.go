package money

import "testing"

func TestMultiplication(t *testing.T) {
	five := NewDollar(5)
	five.Times(2)
	if five.Amount != 10 {
		t.Error(`$5 * 2 が $10じゃなかった`)
	}
	five.Times(3)
	if five.Amount != 15 {
		t.Error(`$5 * 3 が $15じゃなかった`)
	}
}
