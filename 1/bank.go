package money

// Bank : 変換する場所
type Bank struct {
}

// NewBank : Bankを返す
func NewBank() *Bank {
	ret := Bank{}
	return &ret
}

// Reduce : 変換して返す
func (*Bank) Reduce(source Expression, to string) *Money {
	if v, ok := source.(*Money); ok {
		return v
	}
	sum := source.(*Sum)
	return sum.Reduce(to)
}
