package money

type Dollar struct{}

func NewDollar(amount int) *Dollar {
	ret := Dollar{}
	return &ret
}

func (dollar *Dollar) Times(multipier int) {

}
