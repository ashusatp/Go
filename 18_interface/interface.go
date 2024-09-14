package main

// An interface is a type that specifies a collection of method signatures (without implementation).

// type InterfaceName interface {
//     MethodName(params) returnType
// }

import "fmt"

// PaymentGateway interface defines the methods any payment gateway should implement.
type PaymentGateway interface {
	ProcessPayment(amount float64) string
	RefundPayment(amount float64) string
}

// PayPal struct represents PayPal as a payment gateway.
type PayPal struct {
	AccountEmail string
}

// ProcessPayment processes a payment through PayPal.
func (p PayPal) ProcessPayment(amount float64) string {
	return fmt.Sprintf("Processed payment of $%.2f via PayPal from account %s", amount, p.AccountEmail)
}

// RefundPayment processes a refund through PayPal.
func (p PayPal) RefundPayment(amount float64) string {
	return fmt.Sprintf("Refunded $%.2f via PayPal to account %s", amount, p.AccountEmail)
}

// Stripe struct represents Stripe as a payment gateway.
type Stripe struct {
	AccountID string
}

// ProcessPayment processes a payment through Stripe.
func (s Stripe) ProcessPayment(amount float64) string {
	return fmt.Sprintf("Processed payment of $%.2f via Stripe for account %s", amount, s.AccountID)
}

// RefundPayment processes a refund through Stripe.
func (s Stripe) RefundPayment(amount float64) string {
	return fmt.Sprintf("Refunded $%.2f via Stripe to account %s", amount, s.AccountID)
}

// Order struct represents an e-commerce order.
type Order struct {
	ID     string
	Amount float64
}

// Checkout method processes the payment for an order using the given payment gateway.
func (o Order) Checkout(gateway PaymentGateway) {
	fmt.Println(gateway.ProcessPayment(o.Amount))
}

func main() {
	// Create an order
	order := Order{
		ID:     "ORD1234",
		Amount: 99.99,
	}

	// Use PayPal as the payment gateway
	paypal := PayPal{AccountEmail: "user@example.com"}
	order.Checkout(paypal)

	// Use Stripe as the payment gateway
	stripe := Stripe{AccountID: "acct_12345"}
	order.Checkout(stripe)

	// Refund using PayPal
	fmt.Println(paypal.RefundPayment(order.Amount))

	// Refund using Stripe
	fmt.Println(stripe.RefundPayment(order.Amount))
}
