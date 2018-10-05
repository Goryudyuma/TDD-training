package money

import (
	"testing"
)

func TestMultiplication(t *testing.T) {
	five := NewDollar(5)
	product := five.Times(2)
	if *product != *NewDollar(10) {
		t.Error(`$5 * 2 が $10じゃなかった`)
	}
	product = five.Times(3)
	if *product != *NewDollar(15) {
		t.Error(`$5 * 3 が $15じゃなかった`)
	}
}

func TestEquality(t *testing.T) {
	if !NewDollar(5).Equals(NewDollar(5)) {
		t.Error(`$5 と $5 が等価ではなかった`)
	}

	if NewDollar(5).Equals(NewDollar(6)) {
		t.Error(`$5 と $6 が等価だった`)
	}
}
