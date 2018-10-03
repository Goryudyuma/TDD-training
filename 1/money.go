package money

// Dollar : ドル
type Dollar struct {
	Amount int
}

// NewDollar : 新しく(*Dollar)を作る
func NewDollar(amount int) *Dollar {
	ret := Dollar{Amount: amount}
	return &ret
}

// Times : multipier倍にして返す
func (dollar *Dollar) Times(multipier int) *Dollar {
	return NewDollar(dollar.Amount * multipier)
}

// Equals : 等価かどうか返す
func (dollar *Dollar) Equals(another *Dollar) bool {
	return dollar.Amount == another.Amount
}
