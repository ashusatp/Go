package main

import (
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

func main() {

	//  if we don't set any field then that field will we assigned with Zero value
	myOrder := Order{
		id:     1,
		amount: 20,
		status: "received",
	}

	myOrder.createdAt = time.Now()
	myOrder.setStatus("Confirmend")

	fmt.Println(myOrder)
	fmt.Println(myOrder.getAmount())

	//-> Anonymous Structs
	person := struct {
		Name  string
		Age   int
		Email string
	}{
		Name:  "Jane Doe",
		Age:   28,
		Email: "janedoe@example.com",
	}
	fmt.Println(person)

}
