package payment

import "fmt"

type CreditCardPayment struct {
	CardNumber string
	Name       string
}

func (c *CreditCardPayment) Pay(amount float64) error {
	fmt.Printf("Paid %.2f using Credit Card (%s)\n", amount, c.CardNumber)
	return nil
}
