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

func printChannelDeets(pos int, c chan string) {
	fmt.Println("Pos", pos, "Length", len(c), "Capacity", cap(c))
}

func runOne() { //unbuffered. Will not work due to channel being occupied.
	done := make(chan string, 2)
	printChannelDeets(1, done)
	go hello(1, 2, done)
	go hello(2, 1, done)
	go hello(3, 1, done)
	printChannelDeets(2, done)
	fmt.Println(<-done)
	printChannelDeets(3, done)
	fmt.Println(<-done)
	printChannelDeets(4, done)
	fmt.Println(<-done)

	printChannelDeets(5, done)
	fmt.Println("Exiting runOne()")
}

func hello(num int, duration int, done chan string) {
	fmt.Println("Hello(" + strconv.Itoa(num) + ") sleeping for: " + strconv.Itoa(duration))
	time.Sleep(time.Duration(duration) * time.Second)
	fmt.Println("Hello(" + strconv.Itoa(num) + ") awake.")
	done <- strconv.Itoa(num)
	fmt.Println("Exiting hello(" + strconv.Itoa(num) + ")")
}
