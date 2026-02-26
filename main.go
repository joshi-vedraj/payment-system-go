package main

import (
	"fmt"

	"github.com/joshi-vedraj/payment-system-go/payment"
)

func main() {
	fmt.Println("**** LLD - Payment System - Start ****")

	context := payment.PaymentContext{}
	//credit card payment
	card := payment.CreditCardPayment{
		CardNumber: "1234 5678 9101",
		Name:       "Ved",
	}
	context.SetStrategy(&card)
	context.Pay(12000)

	//upi payment
	upi := payment.UPIPayment{
		UPIID: "ved@xyz",
	}
	context.SetStrategy(&upi)
	context.Pay(5000)

	//paypal payment
	paypal := payment.PayPalPayment{
		Email: "vedjosh@xyz.com",
	}
	context.SetStrategy(&paypal)
	context.Pay(5000)

	fmt.Println("***** LLD - Payment System - End *****")
}
