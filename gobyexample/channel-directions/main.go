package main

import (
	"fmt"
	"strconv"
	"time"
)

func hello(num int, duration int, done chan string) {
	fmt.Println("Hello(" + strconv.Itoa(num) + ") sleeping for: " + strconv.Itoa(duration))
	time.Sleep(time.Duration(duration) * time.Second)
	fmt.Println("Hello(" + strconv.Itoa(num) + ") awake.")
	done <- strconv.Itoa(num)
	fmt.Println("Exiting hello(" + strconv.Itoa(num) + ")")
}

func takesReadonly(c <-chan string) {
	// c is now receive-only inside the function and anywhere else it might go from here
}

func returnsReadOnly() <-chan string {
	fmt.Println("Entering returnsReadOnly()")
	c := make(chan string)
	go func() {
		fmt.Println("Entering anon func")
		//go hello(1, 2, c) // This also works as it starts a sub sub gorouting
		hello(1, 2, c)
		fmt.Println("Exiting anon func")
	}()
	fmt.Println("Exiting returnsReadOnly()")
	return c
}

func main() {
	fmt.Println("Begin Here")
	readOnly := returnsReadOnly()
	fmt.Println("After read")                // "<-chan string" (different type)
	fmt.Println("From channel:", <-readOnly) // "hello" (same underlying channel!)
}
