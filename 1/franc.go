package money

// Franc : フラン
type Franc struct {
	Money
	currency string
}

// Times : multipier倍にして返す
func (franc *Franc) Times(multipier int) *Franc {
	return NewFranc(franc.amount * multipier)
}

// Currency : "CHF"を返す
func (franc *Franc) Currency() string {
	return franc.currency
}
