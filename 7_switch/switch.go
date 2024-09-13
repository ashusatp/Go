package main

import (
	"fmt"
	"time"
)

func main() {

	// simple switch
	i := 5

	switch i {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	default:
		fmt.Println("other")
	}

	// multiple condition switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it's weekend")
	default:
		fmt.Println("it's workday")
	}

	// type switch
	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("an Integer")
		case string:
			fmt.Println("a string")
		case bool:
			fmt.Println("a boolean")
		default:
			fmt.Println("other", t)
		}
	}

	whoAmI("golang")

}
