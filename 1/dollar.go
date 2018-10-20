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

// Equals : 等価かどうか返す
func (dollar *Dollar) Equals(another MoneyInterface) bool {
	return dollar.getAmount() == another.getAmount()
}

// getAmount : Amountを返す
func (dollar *Dollar) getAmount() int {
	return dollar.amount
}
