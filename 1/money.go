package money

// Money : 親struct
type Money struct {
	amount   int
	currency string
}

// Equals : 等価かどうか返す
func (money *Money) Equals(another MoneyInterface) bool {
	return money.amount == another.getAmount() &&
		money.currency == another.Currency()
}

// getAmount : Amountを返す
func (money *Money) getAmount() int {
	return money.amount
}

// Currency : currencyを返す
func (money *Money) Currency() string {
	return money.currency
}

// NewMoney : 新しく(Money)を作る
func NewMoney(amount int, currency string) *Money {
	ret := Money{amount: amount, currency: currency}
	return &ret
}

// NewDollar : 新しく(*Money)を作る
func NewDollar(amount int) *Money {
	ret := NewMoney(amount, "USD")
	return ret
}

// NewFranc : 新しく(*Money)を作る
func NewFranc(amount int) *Money {
	ret := NewMoney(amount, "CHF")
	return ret
}

// Times : multipier倍にして返す
func (money *Money) Times(multipier int) *Money {
	return NewMoney(money.amount*multipier, money.currency)
}

// Plus : addedndを足して返す
func (money *Money) Plus(addend *Money) *Sum {
	return NewSum(money, addend)
}

// Reduce : 変換して返す
func (money *Money) Reduce(to string) *Money {
	return money
}

// MoneyInterface : Amountを取り出すinterfaceを持っているもの
type MoneyInterface interface {
	getAmount() int
	Currency() string
}
