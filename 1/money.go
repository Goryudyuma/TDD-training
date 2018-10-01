package money

type Dollar struct {
	Amount int
}

func NewDollar(amount int) *Dollar {
	ret := Dollar{}
	return &ret
}

func (dollar *Dollar) Times(multipier int) {

}
