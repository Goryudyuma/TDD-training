package money

// Money : 親struct
type Money struct {
	amount int
	class  string
}

// Equals : 等価かどうか返す
func (money *Money) Equals(another MoneyInterface) bool {
	return money.amount == another.getAmount() &&
		money.class == another.getClass()
}

// getAmount : Amountを返す
func (money *Money) getAmount() int {
	return money.amount
}

// getClass : Classを返す
func (money *Money) getClass() string {
	return money.class
}

// MoneyInterface : Amountを取り出すinterfaceを持っているもの
type MoneyInterface interface {
	getAmount() int
	getClass() string
}
