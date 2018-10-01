package money

// Dollar : ドル
type Dollar struct {
	Amount int
}

// NewDollar : 新しく(*Dollar)を作る
func NewDollar(amount int) *Dollar {
	ret := Dollar{}
	return &ret
}

// Times : multipier倍にして返す
func (dollar *Dollar) Times(multipier int) {

}
