package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	t1 := time.Now()

	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)

	go myAsyncFunc(1, 5, c1)
	go myAsyncFunc(2, 3, c2)
	go myAsyncFunc(3, 1, c3)

	for i := 0; i < 3; i++ {
		fmt.Println("begin i", i)
		select {
		case msg1 := <-c1:
			fmt.Println("received msg1 ", msg1)
		case msg2 := <-c2:
			fmt.Println("received msg2 ", msg2)
		case msg3 := <-c3:
			fmt.Println("received msg2 ", msg3)
		}
		fmt.Println("end i", i)
	}
	t2 := time.Now()
	diff := t2.Sub(t1)
	fmt.Println(diff)
}

func myAsyncFunc(num int, duration int, done chan string) {
	fmt.Println("#" + strconv.Itoa(num) + " sleeping for: " + strconv.Itoa(duration))
	time.Sleep(time.Duration(duration) * time.Second)
	fmt.Println("#" + strconv.Itoa(num) + " awake.")
	done <- "Hello from #" + strconv.Itoa(num)
	fmt.Println("End of #" + strconv.Itoa(num))
}
