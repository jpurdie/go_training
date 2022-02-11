package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	fmt.Println("Begin")

	done := make(chan string)
	fmt.Println("Main going to call hello go goroutine 1")
	go hello(1, 3, done)
	fmt.Println("Main going to call hello go goroutine 2")
	go hello(2, 1, done)
	y := <-done
	fmt.Println("Main received data x", y)
	y = <-done
	fmt.Println("Main received data y", y)

}

func hello(num int, duration int, done chan string) {
	fmt.Println("Hello(" + strconv.Itoa(num) + ") sleeping for: " + strconv.Itoa(duration))
	time.Sleep(time.Duration(duration) * time.Second)
	fmt.Println("Hello(" + strconv.Itoa(num) + ") awake.")
	done <- strconv.Itoa(num)
}
