package money

// Money : 親struct
type Money struct {
	amount int
}

// Equals : 等価かどうか返す
func (money *Money) Equals(another MoneyInterface) bool {
	return money.amount == another.getAmount()
}

// getAmount : Amountを返す
func (money *Money) getAmount() int {
	return money.amount
}

// MoneyInterface : Amountを取り出すinterfaceを持っているもの
type MoneyInterface interface {
	getAmount() int
}
