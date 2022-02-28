package main

import (
	"fmt"
	"strconv"
)

func printChannelDeets(pos int, c chan string) {
	fmt.Println(strconv.Itoa(pos)+") Length:", len(c), "Capacity:", cap(c))
}

func main() {
	messages := make(chan string)
	//messages := make(chan string, 1)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		// This is not hit because there is no *value* in messages.
		fmt.Println("received message", msg)
	default:
		// Channel is empty so it came here.
		fmt.Println("no message received")
	}

	msg := "hi"

	/* Uncommenting this will make it send a message into messages */
	// go func() { messages <- msg }()
	// fmt.Println(<-messages)

	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		/*
			A non-blocking send works similarly. Here msg cannot be sent to the messages channel,
			because the channel has no buffer and there is no receiver.
			Therefore the default case is selected.

			There is no sender because there was no value written to it.
		*/
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
	printChannelDeets(4, messages)

}
