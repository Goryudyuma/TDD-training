package money

import "testing"

func TestMultiplication(t *testing.T) {
	five := NewDollar(5)
	product := five.Times(2)
	if product.Amount != 10 {
		t.Error(`$5 * 2 が $10じゃなかった`)
	}
	product = five.Times(3)
	if product.Amount != 15 {
		t.Error(`$5 * 3 が $15じゃなかった`)
	}
}
