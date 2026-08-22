package main

import (
	"fmt"
)

type paymenter interface {
	pay(amount float32)
}

type payment struct {
	// gateway stripe
	gateway paymenter
}

func (p payment) makePayment(amount float32) {
	/*
		razorpayPaymentGw := razorpay{}
		stripePaymentGw := stripe{}
		razorpayPaymentGw.pay(amount)
		stripePaymentGw.pay(amount)
	*/

	p.gateway.pay(amount)

}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	//apis
	fmt.Println("making payment using razorpay", amount)
}

type fakepayment struct{}

func (s fakepayment) pay(amount float32) {
	//apis
	fmt.Println("making payment using fakepayment", amount)
}

type stripe struct{}

func (s stripe) pay(amount float32) {
	//apis
	fmt.Println("making payment using stripe", amount)
}

type paypal struct{}

func (p paypal) pay(amount float32) {
	fmt.Println("paypal implenmented")
}

type paytm struct{}

func (p paytm) pay(amount float32) {
	fmt.Println("i am using paytm payment")
}

func main() {
	// newPayment := payment{}
	// newPayment.makePayment(100)
	// fakeGw := fakepayment{}
	// paypalGw := paypal{}
	// paytmGw := paytm{}
	// razorpayGw := razorpay{}
	stripepayGw := stripe{}
	newPayment := payment{
		// gateway: fakeGw,
		// gateway: paypalGw,
		// gateway: paytmGw,
		// gateway: razorpayGw,
		gateway: stripepayGw,
	}
	newPayment.makePayment(100)
}
