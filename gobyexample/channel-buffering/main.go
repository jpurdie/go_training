package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// https://medium.com/a-journey-with-go/go-buffered-and-unbuffered-channels-29a107c00268

/*
Channels are a typed conduit through which you can send and receive values with the channel operator, <-
*/

/*
If the channel is unbuffered, the sender blocks until the receiver has received the value
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
	default:
		fmt.Println("Enter a val")
	}
}

func printChannelDeets(c chan string) {
	fmt.Println("Length", len(c), "Capacity", cap(c))
}

func runTwo() { //buffered
	done := make(chan string, 2)
	printChannelDeets(done)
	done <- "unbuffered 1"
	printChannelDeets(done)
	done <- "unbuffered 2"
	printChannelDeets(done)
	fmt.Println(<-done)
	printChannelDeets(done)
	fmt.Println(<-done)
	printChannelDeets(done)

}

func runOne() { //unbuffered. Will not work due to channel being occupied.
	done := make(chan string)
	printChannelDeets(done)
	done <- "unbuffered 1"
	printChannelDeets(done)
	done <- "unbuffered 2" // will fail here with deadlock error
	fmt.Println(<-done)
	fmt.Println(<-done)
}

func hello(num int, duration int, done chan string) {
	fmt.Println("Hello(" + strconv.Itoa(num) + ") sleeping for: " + strconv.Itoa(duration))
	time.Sleep(time.Duration(duration) * time.Second)
	fmt.Println("Hello(" + strconv.Itoa(num) + ") awake.")
	done <- strconv.Itoa(num)
	fmt.Println("End of hello(" + strconv.Itoa(num) + ")")
}
