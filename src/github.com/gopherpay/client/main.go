package main

import (
	"github.com/gopherpay/payment"
)

// func main() {
// 	credit := payment.CreateCreditAccount(
// 		"Debra Ingram",
// 		"1111-2222-3333-4444",
// 		5,
// 		2021,
// 		123,
// 	)

// 	fmt.Printf("Owner name: %v\n", credit.OwnerName())
// 	fmt.Printf("Card number: %v\n", credit.CardNumber())
// 	fmt.Println("Trying to change card number")
// 	err := credit.SetCardNumber("Invalid")
// 	if err != nil {
// 		fmt.Printf("That didn't work: %v\n", err)
// 	}
// 	fmt.Printf("Available credit: %v\n", credit.AvailableCredit())
// }

func main() {
	var option payment.PaymentOption

	option = payment.CreateCreditAccount(
		"Debra Ingram",
		"1111-2222-3333-4444",
		5,
		2021,
		123,
	)

	option.ProcessPayment(500)

	option = payment.CreateCashAccount()

	option.ProcessPayment(500)
}
