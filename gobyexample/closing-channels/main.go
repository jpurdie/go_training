package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func doInts() {

	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	fmt.Println(<-done)
}

func doStrings() {

	jobs := make(chan string, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		fmt.Println("sent job", j)
		myAsyncFunc(j, jobs)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	fmt.Println(<-done) //blocking using the synchronization technique
}

func main() {
	fmt.Println("Begin")

	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Enter a val")
		return
	}
	switch args[0] {
	case "1":
		doInts()
	case "2":
		doStrings()
	}
}

func myAsyncFunc(num int, done chan string) {
	rand.Seed(time.Now().UnixNano())
	a, b := 1, 5
	n := a + rand.Intn(b-a+1)
	fmt.Println("#" + strconv.Itoa(num) + " sleeping for: " + strconv.Itoa(n))
	time.Sleep(time.Duration(n) * time.Second)
	fmt.Println("#" + strconv.Itoa(num) + " awake.")
	done <- "Hello from #" + strconv.Itoa(num)
	fmt.Println("End of #" + strconv.Itoa(num))
}
