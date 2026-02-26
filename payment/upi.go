package payment

import "fmt"

type UPIPayment struct {
	UPIID string
}

func (u *UPIPayment) Pay(amount float64) error {
	fmt.Printf("Paid %.2f using UPI (%s)\n", amount, u.UPIID)
	return nil
}
