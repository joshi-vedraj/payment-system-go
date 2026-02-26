package payment

import "errors"

type PaymentContext struct {
	strategy PaymentStrategy
}

func (pc *PaymentContext) SetStrategy(ps PaymentStrategy) {
	pc.strategy = ps
}

func (pc *PaymentContext) Pay(amount float64) error {
	if pc.strategy == nil {
		return errors.New("Payment Strategy not set")
	}
	return pc.strategy.Pay(amount)
}
