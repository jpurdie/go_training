package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

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
	a := 3
	timer1 := time.NewTimer(time.Duration(a) * time.Second)
	begin := time.Now()
	fmt.Println("Timer 1 pre-fired")
	<-timer1.C // it will wait here for 3 seconds
	fmt.Println("Timer 1 fired", "Should be ~"+strconv.Itoa(a)+"s. And it is...", time.Since(begin))

	a = 2
	timer2 := time.NewTimer(time.Duration(a) * time.Second)
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
