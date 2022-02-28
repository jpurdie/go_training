package main

import (
	"fmt"
	"time"
)

func main() {

	timer1 := time.NewTimer(3 * time.Second)
	/*
		Timers represent a single event in the future.
		You tell the timer how long you want to wait, and it provides a channel
		 that will be notified at that time. This timer will wait 2 seconds.
	*/

	/*
		If you just wanted to wait, you could have used time.Sleep.
		One reason a timer may be useful is
		that you can cancel the timer before it fires. Here’s an example of that.
	*/
	fmt.Println("Timer 1 pre-fired", time.Now())
	<-timer1.C
	fmt.Println("Timer 1 fired", time.Now())

	timer2 := time.NewTimer(2 * time.Second)
	go func() {
		fmt.Println("Timer 2 pre-fired", time.Now())
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	time.Sleep(5 * time.Second)
}
