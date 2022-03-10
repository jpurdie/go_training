package main

import (
	"fmt"
	"strconv"
	"time"
)

/*
Channels are a typed conduit through which you can send and receive values with
the channel operator, <-
*/
func main() {
	fmt.Println("Begin")
	runOne()
	runTwo()
	fmt.Println("End")
}

func runTwo() {
	done := make(chan string)
	go worker(1, 2, done)
	go worker(2, 5, done)
	x := <-done
	fmt.Println("Main received data x", x)
	y := <-done
	fmt.Println("Main received data y", y)
}

func runOne() {
	done := make(chan string)
	go worker(1, 15, done)
	go worker(2, 5, done)
	x := <-done
	fmt.Println("Main received data x", x)
	y := <-done
	fmt.Println("Main received data y", y)
}

func worker(num int, duration int, done chan string) {
	fmt.Println("worker(" + strconv.Itoa(num) + ") sleeping for: " + strconv.Itoa(duration))
	time.Sleep(time.Duration(duration) * time.Second)
	fmt.Println("worker(" + strconv.Itoa(num) + ") awake.")
	done <- strconv.Itoa(num)
	fmt.Println("End of worker()")
}
