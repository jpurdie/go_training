package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

/*
Channels are a typed conduit through which you can send and receive values with the channel operator, <-
*/
func main() {
	fmt.Println("Begin")

	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Enter a val")
		return
	}
	switch args[0] {
	case "1":
		runOne()
	case "2":
		runTwo()
	case "3":
		runThree()
	case "4":
		runFour()
	case "5":
		runFive()
	}

}

func runFive() { // removing channel receivers as a test.
	done := make(chan string)
	fmt.Println("Main going to call hello go goroutine 1")
	go hello(1, 5, done)
	fmt.Println("Main going to call hello go goroutine 2")
	go hello(2, 1, done)
}

func runFour() {
	done := make(chan string)
	fmt.Println("Main going to call hello go goroutine 1")
	go hello(1, 5, done)
	fmt.Println("Main going to call hello go goroutine 2")
	go hello(2, 1, done)
	y := <-done
	fmt.Println("Main received data x", y)
	fmt.Println("Calling 3")
	go hello(3, 5, done)
	fmt.Println("After calling 3")
	y = <-done
	fmt.Println("Main received data y", y)
	fmt.Println("Before last done")
	<-done
	fmt.Println("After done")

}

func runThree() {
	done := make(chan string)
	fmt.Println("Main going to call hello go goroutine 1")
	go hello(1, 5, done)
	fmt.Println("Main going to call hello go goroutine 2")
	go hello(2, 1, done)
	y := <-done
	fmt.Println("Main received data x", y)
	fmt.Println("Calling 3")
	go hello(3, 5, done)
	fmt.Println("After calling 3")
	y = <-done
	fmt.Println("Main received data y", y)

	// Why is 3 never received?
	/*
		By default channels are unbuffered, meaning that they will only accept
		sends (chan <-) if there is a corresponding receive (<- chan) ready to receive the sent value.
	*/
}

func runTwo() {
	done := make(chan string)
	fmt.Println("Main going to call hello go goroutine 1")
	go hello(1, 2, done)
	fmt.Println("Main going to call hello go goroutine 2")
	go hello(2, 5, done)
	x := <-done
	fmt.Println("Main received data x", x)
	y := <-done
	fmt.Println("Main received data y", y)
}

func runOne() {
	done := make(chan string)
	fmt.Println("Main going to call hello go goroutine 1")
	go hello(1, 5, done)
	fmt.Println("Main going to call hello go goroutine 2")
	go hello(2, 2, done)
	x := <-done
	fmt.Println("Main received data x", x)
	y := <-done
	fmt.Println("Main received data y", y)
}

func hello(num int, duration int, done chan string) {
	fmt.Println("Hello(" + strconv.Itoa(num) + ") sleeping for: " + strconv.Itoa(duration))
	time.Sleep(time.Duration(duration) * time.Second)
	fmt.Println("Hello(" + strconv.Itoa(num) + ") awake.")
	done <- strconv.Itoa(num)
	fmt.Println("End of hello()")
}
