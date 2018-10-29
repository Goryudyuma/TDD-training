package money

// Dollar : ドル
type Dollar struct {
	Money
}

// Times : multipier倍にして返す
func (dollar *Dollar) Times(multipier int) *Dollar {
	return NewDollar(dollar.amount * multipier)
}

// Currency : "USD"を返す
func (*Dollar) Currency() string {
	return "USD"
}
