package main

import "fmt"

// Define a struct for general Vehicle
type Vehicle struct {
	Brand string
	Year  int
}

// Define a Car struct that embeds Vehicle
type Car struct {
	Vehicle // Embedding Vehicle struct
	Model   string
	Mileage int
}

func main() {
	// Create an instance of Car

	// [+] general way
	// myVehicles := Vehicle{
	// 	Brand: "Toyota",
	// 	Year:  2015,
	// }

	// myCar := Car{
	// 	Vehicle: myVehicles,
	// 	Model:   "Corolla",
	// 	Mileage: 45000,
	// }

	// [+] short hand
	myCar := Car{
		Vehicle: Vehicle{
			Brand: "Toyota",
			Year:  2015,
		},
		Model:   "Corolla",
		Mileage: 45000,
	}

	myCar.Vehicle.Year = 2020

	// Access fields from the embedded Vehicle struct
	fmt.Println(myCar.Brand) // Output: Toyota
	fmt.Println(myCar.Year)  // Output: 2015
	fmt.Println(myCar.Model) // Output: Corolla
}
