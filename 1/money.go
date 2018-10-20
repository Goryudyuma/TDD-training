package money

// Money : 親struct
type Money struct {
	amount int
}

type MoneyInterface interface {
	getAmount() int
}
