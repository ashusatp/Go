package main

import (
	"errors"
	"fmt"
	"time"
)

type Order struct {
	id        int
	amount    int
	status    string
	createdAt time.Time
}

// [+] constructor for struct
func NewOrder(id int, amount int, status string) *Order {
	myOrder := Order{
		id:     id,
		amount: amount,
		status: status,
	}

	return &myOrder
}

func (O *Order) setStatus(status string) {
	O.status = status
}

// if we are not modifing struct so no need to pass O *Order , go with O Order
func (O Order) getAmount() int {
	return O.amount
}

// func (O *Order) getAmount() int {
// 	return O.amount
// }

// [+] Constructor with Validation
type User struct {
	Username string
	Age      int
}

// Constructor function with validation
func NewUser(username string, age int) (*User, error) {
	if age < 0 {
		return nil, errors.New("age cannot be negative")
	}
	return &User{
		Username: username,
		Age:      age,
	}, nil
}

func main() {

	myOrder := NewOrder(1, 20, "Received")
	fmt.Println(myOrder)
	myOrder.setStatus("Confirmed")
	fmt.Println(myOrder.getAmount())

	//-> Validation
	user, err := NewUser("Alice", 25)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(user.Username) // Output: Alice
		fmt.Println(user.Age)      // Output: 25
	}

	// Trying to create a user with invalid age
	invalidUser, err := NewUser("Bob", -5)
	if err != nil {
		fmt.Println("Error:", err) // Output: Error: age cannot be negative
	} else {
		fmt.Println(invalidUser.Username)
		fmt.Println(invalidUser.Age)
	}

}
