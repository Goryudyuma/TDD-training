package money

// Dollar : ドル
type Dollar struct {
	Money
	currency string
}

// Times : multipier倍にして返す
func (dollar *Dollar) Times(multipier int) *Dollar {
	return NewDollar(dollar.amount * multipier)
}

// Currency : "USD"を返す
func (dollar *Dollar) Currency() string {
	return dollar.currency
}
