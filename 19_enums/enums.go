package main

import "fmt"

// Define OrderStatus as a set of enumerated constants
type OrderStatus int

const (
	Pending    OrderStatus = iota // 0
	Processing                    // 1
	Shipped                       // 2
	Delivered                     // 3
)

// [+] string enum
type OrderStringStatus string

const (
	PendingString    OrderStringStatus = "pending"
	ProcessingString                   = "processing"
	ShippedString                      = "shipped"
	DeliveredString                    = "delivered"
)

//[+] Other Use Cases of Enums

// Payment Methods:
type PaymentMethod int

const (
	CreditCard PaymentMethod = iota
	DebitCard
	PayPal
	BankTransfer
)

// User Roles:
type UserRole int

const (
	Admin UserRole = iota
	Moderator
	Member
	Guest
)

// Product Categories:
type ProductCategory int

const (
	Electronics ProductCategory = iota
	Clothing
	HomeAppliances
	Books
)

// String method for pretty-printing the enum values
func (status OrderStatus) String() string {
	switch status {
	case Pending:
		return "Pending"
	case Processing:
		return "Processing"
	case Shipped:
		return "Shipped"
	case Delivered:
		return "Delivered"
	default:
		return "Unknown"
	}
}

func main() {
	var status OrderStatus = Processing

	fmt.Println(status)  // Output: Processing
	fmt.Println(Shipped) // Output: Shipped

	// Example of changing status
	status = Delivered
	fmt.Println(status) // Output: Delivered
}
