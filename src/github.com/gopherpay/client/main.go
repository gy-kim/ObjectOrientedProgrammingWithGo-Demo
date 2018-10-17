/* Interface

package main

import (
	"github.com/gopherpay/payment"
)

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

*/

/*  // Message passing with Channel
package main

import "fmt"

type CreditAccount struct{}

func (c *CreditAccount) processPayment(amount float32) {
	fmt.Println("Processing credit card payment...")
}

func CreateCreditAccount(chargeCh chan float32) *CreditAccount {
	creditAccount := &CreditAccount{}
	go func(chargeCh chan float32) {
		for amount := range chargeCh {
			creditAccount.processPayment(amount)
		}
	}(chargeCh)

	return creditAccount
}
func main() {
	chargeCh := make(chan float32)
	CreateCreditAccount(chargeCh)
	chargeCh <- 500
	var a string
	fmt.Scanln(&a)
}

*/

package main

import "fmt"

type Account struct{}

func (a *Account) AvailableFunds() float32 {
	fmt.Println("Listening available funds")
	return 0
}

func (a *Account) ProcessPayment(amount float32) bool {
	fmt.Println("Processing payment")
	return true
}

type CreditAccount struct {
	Account
}

func main() {
	ca := &CreditAccount{}
	ca.AvailableFunds()
	ca.ProcessPayment(500)
}
