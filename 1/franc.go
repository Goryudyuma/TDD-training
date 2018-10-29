package money

// Franc : フラン
type Franc struct {
	Money
}

// Times : multipier倍にして返す
func (franc *Franc) Times(multipier int) *Franc {
	return NewFranc(franc.amount * multipier)
}
