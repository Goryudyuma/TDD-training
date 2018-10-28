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

// NewDollar : 新しく(*Dollar)を作る
func NewDollar(amount int) *Dollar {
	ret := Dollar{Money{amount: amount, class: "Dollar"}}
	return &ret
}

// NewFranc : 新しく(*Franc)を作る
func NewFranc(amount int) *Franc {
	ret := Franc{Money{amount: amount, class: "Franc"}}
	return &ret
}

// MoneyInterface : Amountを取り出すinterfaceを持っているもの
type MoneyInterface interface {
	getAmount() int
	getClass() string
}
