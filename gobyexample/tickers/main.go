package main

import (
	"fmt"
	"time"
)

/*

	Tickers are for when you want to do something repeatedly at regular intervals.

*/

func main() {
	/*
		NewTicker returns a new Ticker containing a channel that
		will send the time on the channel after each tick.

		A Ticker holds a channel that delivers “ticks” of a clock at intervals

		The ticker will adjust the time interval or drop ticks to make up for slow receivers.
	*/
	/*
		type Ticker struct {
			C <-chan Time // The channel on which the ticks are delivered.
			// contains filtered or unexported fields
		}
	*/
	ticker := time.NewTicker(1000 * time.Millisecond)
	done := make(chan bool)

	go foo(done, ticker)

	time.Sleep(3000)
	ticker.Stop()
	fmt.Println("Pausing")
	time.Sleep(6000 * time.Millisecond) // While this is paused, foo will be executing its for loop.
	ticker.Stop()                       // stops the ticker to close the channel
	done <- true                        // Send true to the done channel to have the for loop break;
	fmt.Println("Ticker stopped")
}

func foo(done chan bool, ticker *time.Ticker) {
	fmt.Println("foo")
	for {
		// This will continue until the time.sleep is paused.
		fmt.Println("For loop start")
		select {
		case t := <-ticker.C: // this will hit when t is of type receiver
			fmt.Println("Prints the current time", t) // Printing the time on the channel
		case <-done: // this will hit when done is of type receiver
			return
		}
	}

}
