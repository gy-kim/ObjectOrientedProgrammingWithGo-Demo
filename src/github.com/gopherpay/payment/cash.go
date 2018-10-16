package payment

import "fmt"

type Cash struct{}

func CreateCashAccount() *Cash {
	return &Cash{}
}

func (a Cash) ProcessPayment(amount float32) bool {
	fmt.Println("Process a cash transaction...")
	return true
}
