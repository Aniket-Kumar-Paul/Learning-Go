package main

import (
	"fmt"
	// "math/rand"
	"time"
)

func processNum(numChan chan int) {
	for num:=range numChan {
		fmt.Println("Processing number:", num)
		time.Sleep(time.Second)
	}
}

func sum(result chan int, num1 int, num2 int) {
	sum := num1 + num2
	result <- sum
}

// goroutine synchronizer
func task(done chan bool) {
	defer func() { done <- true }() // will run after function executes, irrespective of any errors or exceptions
	fmt.Println("Processing...")
}

func emailSender(emailChan <-chan string, done chan<- bool) { // emailChan is receive only channel & done is send only channel here
	defer func() { done <- true }()
	for email:=range emailChan {
		fmt.Println("Sending email to..", email)
		time.Sleep(time.Second)
	}
}

// channels -> used to send/receive data b/w different go routines
func main() {
	// messageChannel := make(chan string) // unbuffered channel
	// messageChannel <- "data" // sending
	// myData := <- messageChannel // receiving
	// fmt.Println(msg)
	// The above will give DEADLOCK error as we're using unbuffered channels &
	// A send (channel <- value) will block until another goroutine is ready to receive.
	// A receive (<-channel) will block until another goroutine sends.

	// Soln -> 
	// 1. Using go routine:
		// go func() {
		// 		messageChannel <- "data" // sending in a goroutine
		// }()
	// 2. Using a buffered channel (Buffered channels don’t require an immediate receiver — the message can "sit" in the channel buffer until it's received):
		// messageChannel := make(chan string, 1) // buffer size of 1

	// numChan := make(chan int)
	// go processNum(numChan)
	// for {
	// 	numChan <- rand.Intn(100)
	// }

	result := make(chan int)
	go sum(result, 4, 5)
	res := <- result
	fmt.Println(res)

	done := make(chan bool)
	go task(done)
	<- done // receive (block). Thus, we can use both WAITGROUPS or CHANNELS for synchroniztion/blocking
	// Preferred is, to use channels if single goroutine, for multiple, use waitgroups

	// Buffered Channel
	emailChan := make(chan string, 100)
	emailChan <- "ani1@gmail.com"
	emailChan <- "ani2@gmail.com"
	fmt.Println(<-emailChan)
	fmt.Println(<-emailChan)

	go emailSender(emailChan, done)
	for i:=0; i<100; i++ {
		emailChan <- fmt.Sprintf("%d@gmail.com", i)
	}
	fmt.Println("Done sending!")
	close(emailChan) // close channel (important)
	<- done

	// multiple channels
	chan1 := make(chan int)
	chan2 := make(chan string)
	go func() {
		chan1 <- 10
	}()
	go func() {
		chan2 <- "hi"
	}()
	for i:=0;i<2;i++ {
		select {
			case chan1Val := <- chan1 :
				fmt.Println("Received data from chan1", chan1Val)
			case chan2Val := <- chan2 :
				fmt.Println("Received data from chan2", chan2Val)
		}
	}
}