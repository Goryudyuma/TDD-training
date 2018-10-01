package money

import "testing"

func TestMultiplication(t *testing.T) {
	five := NewDollar(5)
	five.times(2)
	if five.Amount != 10 {
		t.Error(`$5 * 2 が $10じゃなかった。`)
	}
}
