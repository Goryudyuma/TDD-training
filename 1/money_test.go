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

	if NewFranc(5).Equals(NewDollar(5)) {
		t.Error(`5 CHF と $5 が等価だった`)
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

func TestSimpleAddition(t *testing.T) {
	five := NewDollar(5)
	sum := five.Plus(five)
	bank := NewBank()
	reduced := bank.Reduce(sum, "USD")
	if !NewDollar(10).Equals(reduced) {
		t.Error(`$5 + $5 をUSDに変換すると、 $10 となる`)
	}
}

func TestPlusReturnsSum(t *testing.T) {
	five := NewDollar(5)
	result := five.Plus(five)
	sum := result
	if !five.Equals(sum.Augend) {
		t.Error(`$5 は Sum.Augendではなかった`)
	}
	if !five.Equals(sum.Addend) {
		t.Error(`$5 は Sum.Addendではなかった`)
	}
}

func TestReduceSum(t *testing.T) {
	sum := NewSum(NewDollar(3), NewDollar(4))
	bank := NewBank()
	result := bank.Reduce(sum, "USD")
	if !NewDollar(7).Equals(result) {
		t.Error(`$3 と $4 を足して、USDで取り出すと、 $7 となる`)
	}
}
