package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func doInts() {
	jobs := make(chan int, 2)
	done := make(chan bool)

	go func() {
		fmt.Println("---")
		for {
			j, more := <-jobs
			if more {
				fmt.Println("Received job", j)
			} else {
				fmt.Println("Received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		fmt.Println("Preparing to send job", j)
		jobs <- j
		fmt.Println("Sent job", j)
	}
	fmt.Println("Sent all jobs. Preparing to close.")
	close(jobs)
	fmt.Println("Jobs closed")
	fmt.Println("<-done", <-done) //blocking using the synchronization technique
}

func main() {
	fmt.Println("Begin")
	doInts()
}

func worker(num int, done chan string) {
	rand.Seed(time.Now().UnixNano())
	a, b := 1, 5
	n := a + rand.Intn(b-a+1)
	fmt.Println("#" + strconv.Itoa(num) + " sleeping for: " + strconv.Itoa(n))
	time.Sleep(time.Duration(n) * time.Second)
	fmt.Println("#" + strconv.Itoa(num) + " awake.")
	done <- "Hello from #" + strconv.Itoa(num)
	fmt.Println("End of #" + strconv.Itoa(num))
}
