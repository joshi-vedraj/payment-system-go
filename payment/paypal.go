package payment

import "fmt"

type PayPalPayment struct {
	Email string
}

func (p *PayPalPayment) Pay(amount float64) error {
	fmt.Printf("Paid %.2f using Paypal (%s)\n", amount, p.Email)
	return nil
}
