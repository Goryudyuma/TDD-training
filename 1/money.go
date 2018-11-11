package money

// Money : 親struct
type Money struct {
	amount   int
	class    string
	currency string
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

// Currency : currencyを返す
func (money *Money) Currency() string {
	return money.currency
}

// NewMoney : 新しく(Money)を作る
func NewMoney(amount int, class string, currency string) *Money {
	ret := Money{amount: amount, class: class, currency: currency}
	return &ret
}

// NewDollar : 新しく(*Money)を作る
func NewDollar(amount int) *Money {
	ret := NewMoney(amount, "Dollar", "USD")
	return ret
}

// NewFranc : 新しく(*Money)を作る
func NewFranc(amount int) *Money {
	ret := NewMoney(amount, "Franc", "CHF")
	return ret
}

// Times : multipier倍にして返す
func (money *Money) Times(multipier int) *Money {
	return NewMoney(money.amount*multipier, money.class, money.currency)
}

// Plus : addedndを足して返す
func (money *Money) Plus(addend *Money) *Sum {
	return NewSum(money, addend)
}

// MoneyInterface : Amountを取り出すinterfaceを持っているもの
type MoneyInterface interface {
	getAmount() int
	getClass() string
	Currency() string
}
