package money

// Franc : フラン
type Franc struct {
	Money
}

// Times : multipier倍にして返す
func (Franc *Franc) Times(multipier int) *Franc {
	return NewFranc(Franc.amount * multipier)
}
