package money

// Dollar : ドル
type Dollar struct {
	amount int
}

// NewDollar : 新しく(*Dollar)を作る
func NewDollar(amount int) *Dollar {
	ret := Dollar{amount: amount}
	return &ret
}

// Times : multipier倍にして返す
func (dollar *Dollar) Times(multipier int) *Dollar {
	return NewDollar(dollar.amount * multipier)
}

// Equals : 等価かどうか返す
func (dollar *Dollar) Equals(another *Dollar) bool {
	return dollar.amount == another.amount
}
