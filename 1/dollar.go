package money

// Dollar : ドル
type Dollar struct {
	Money
}

// NewDollar : 新しく(*Dollar)を作る
func NewDollar(amount int) *Dollar {
	ret := Dollar{Money{amount: amount}}
	return &ret
}

// Times : multipier倍にして返す
func (dollar *Dollar) Times(multipier int) *Dollar {
	return NewDollar(dollar.amount * multipier)
}
