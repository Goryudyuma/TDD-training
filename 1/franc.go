package money

// Franc : フラン
type Franc struct {
	Money
}

// NewFranc : 新しく(*Franc)を作る
func NewFranc(amount int) *Franc {
	ret := Franc{Money{amount: amount}}
	return &ret
}

// Times : multipier倍にして返す
func (Franc *Franc) Times(multipier int) *Franc {
	return NewFranc(Franc.amount * multipier)
}
