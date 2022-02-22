package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	/*
		Note that the channel is buffered, so the send in the goroutine is nonblocking.
		This is a common pattern to prevent goroutine leaks in case the channel is never read
	*/

	c1 := make(chan string, 1)
	go myAsyncFunc(1, 2, c1)

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go myAsyncFunc(2, 2, c2)
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}

func myAsyncFunc(num int, duration int, done chan string) {
	fmt.Println("#" + strconv.Itoa(num) + " sleeping for: " + strconv.Itoa(duration))
	time.Sleep(time.Duration(duration) * time.Second)
	fmt.Println("#" + strconv.Itoa(num) + " awake.")
	done <- "Hello from #" + strconv.Itoa(num)
	fmt.Println("End of #" + strconv.Itoa(num))
}
