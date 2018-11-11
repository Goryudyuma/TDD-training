package money

// Sum : 足し算を表す構造体
type Sum struct {
	Augend, Addend *Money
}

// NewSum : 足し算を表す構造体を返す
func NewSum(augend, addend *Money) *Sum {
	return &Sum{Augend: augend, Addend: addend}
}
