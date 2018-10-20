package money

// Money : 親struct
type Money struct {
	amount int
}

// Equals : 等価かどうか返す
func (money *Money) Equals(another MoneyInterface) bool {
	return money.amount == another.getAmount()
}

type MoneyInterface interface {
	getAmount() int
}
