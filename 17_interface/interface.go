package main

import "fmt"

type paymenter interface { // struct implementing this interface, must implement all methods
	pay(amount float32)
	refund(amount float32, account string)
}

type payment struct{
	gateway paymenter // the gateway must be a struct that implements the paymenter interface
}
func (p payment) makePayment(amount float32) {
	p.gateway.pay(amount)
}

type razorpay struct{}
func (r razorpay) pay(amount float32) { // In go, interfaces are implemented implicitly automatically
	// logic to make payment
	fmt.Println("Making payment using razorpay", amount)
}
func (r razorpay) refund(amount float32, account string) {
	// logic to make refund
	fmt.Println("Refunding payment using razorpay", amount, "to account", account)
}

type stripe struct{}
func (s stripe) pay(amount float32) {
	// logic to make payment
	fmt.Println("Making payment using stripe", amount)
}
func (s stripe) refund(amount float32, account string) {
	// logic to make refund
	fmt.Println("Refunding payment using stripe", amount, "to account", account)
}

func main() {
	p := payment{
		gateway: stripe{},
	}
	p.makePayment(1000.0)
}