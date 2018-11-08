package money

import (
	"testing"
)

func TestMultiplication(t *testing.T) {
	five := NewDollar(5)
	if !(*five).Times(2).Equals(NewDollar(10)) {
		t.Error(`$5 * 2 が $10じゃなかった`)
	}
	if !(*five).Times(3).Equals(NewDollar(15)) {
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

	if !NewFranc(5).Equals(NewFranc(5)) {
		t.Error(`5 CHF と 5 CHF が等価ではなかった`)
	}

	if NewFranc(5).Equals(NewFranc(6)) {
		t.Error(`5 CHF と 6 CHF が等価だった`)
	}

	if NewFranc(5).Equals(NewDollar(5)) {
		t.Error(`5 CHF と $5 が等価だった`)
	}
}

func TestFrancMultiplication(t *testing.T) {
	five := NewFranc(5)
	if !(*five).Times(2).Equals(NewFranc(10)) {
		t.Error(`$5 * 2 が $10じゃなかった`)
	}
	if !(*five).Times(3).Equals(NewFranc(15)) {
		t.Error(`$5 * 3 が $15じゃなかった`)
	}
}

func TestCurrency(t *testing.T) {
	if "USD" != NewDollar(1).Currency() {
		t.Error(`DollarはUSDではなかった`)
	}
	if "CHF" != NewFranc(1).Currency() {
		t.Error(`FrancはCHFではなかった`)
	}
}
