package main

import (
	"fmt"
	"time"
)

func processNum(numChan chan int) {
	fmt.Println("Process number: ", <-numChan)
}

func processNumMany(numChan chan int) {
	for num := range numChan {
		fmt.Println("Process number: ", num)
		time.Sleep(time.Second)
	}
}

func sum(result chan int, num1, num2 int) {

	numResult := num1 + num2
	result <- numResult

}

func task(done chan bool) {
	defer func() { done <- true }()
	fmt.Println("Processing...")
}

func emailSender(emailChan <-chan string, done chan<- bool) {
	defer func() { done <- true }()

	for email := range emailChan {
		fmt.Println("Sending email to: ", email)
		time.Sleep(time.Second)
	}
}

func main() {

	// -> sending data into goroutine
	// numChan := make(chan int)
	// go processNum(numChan)
	// numChan <- 5 // blocking
	// time.Sleep(time.Second)

	// -> sending series of data through channel to goroutine
	// numChan := make(chan int)
	// go processNumMany(numChan)
	// for {
	// 	numChan <- rand.Intn(100)
	// }

	// -> Receiving data
	// result := make(chan int)
	// go sum(result, 1, 2)
	// res := <-result //blocking
	// fmt.Println(res)

	// -> sync goroutine using channels / same like wait groups
	// done := make(chan bool)
	// go task(done)
	// <- done //blocking

	// -> Buffered channel (Queue system implementaion)
	// emailChan := make(chan string, 100)
	// done := make(chan bool)

	// go emailSender(emailChan, done)

	// for i := 0; i < 100; i++ {
	// 	emailChan <- fmt.Sprintf("%d@gmail.com", i)
	// }

	// fmt.Println("email sending done")

	// close(emailChan) //close buffered channel (otherwise it will be deadlock)

	// <-done // blocking

	// -> Listening data on multiple channels at a time
	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 10
	}()

	go func() {
		chan2 <- "Ping"
	}()

	for i := 0; i < 2; i++ {
		select {
		case chan1Val := <-chan1:
			println(chan1Val)
		case chan2Val := <-chan2:
			println(chan2Val)
		}
	}
}
