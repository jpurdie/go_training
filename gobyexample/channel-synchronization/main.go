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
		//runTwo()
	default:
		fmt.Println("Enter a val")
	}
}

func printChannelDeets(c *chan string) {
	fmt.Println("Length", len(*c), "Capacity", cap(*c))
}

func runOne() { //unbuffered. Will not work due to channel being occupied.
	done := make(chan string)
	printChannelDeets(&done)
	go hello(1, 1, done)
	printChannelDeets(&done)
	go hello(2, 2, done)
	printChannelDeets(&done)
	//	<-done
	fmt.Println("Exiting runOne()")
}

func hello(num int, duration int, done chan string) {
	fmt.Println("Hello(" + strconv.Itoa(num) + ") sleeping for: " + strconv.Itoa(duration))
	time.Sleep(time.Duration(duration) * time.Second)
	fmt.Println("Hello(" + strconv.Itoa(num) + ") awake.")
	done <- strconv.Itoa(num)
	fmt.Println("Exiting hello(" + strconv.Itoa(num) + ")")
}
