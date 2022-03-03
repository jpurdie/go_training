package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2 // It is multiplied by two, why?
	}
}

func main() {

	const numJobs = 17
	jobs := make(chan int, numJobs)    //Making buffered channel
	results := make(chan int, numJobs) //Making buffered channel

	fmt.Println("Before worker calls", fmt.Sprint(time.Now().UnixMicro()))

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results) //this creates 3 "workers" which will continue to process
	}

	fmt.Println("After worker calls", fmt.Sprint(time.Now().UnixMicro()))

	for j := 1; j <= numJobs; j++ {
		jobs <- j //Send an integer to the job buffered channel
	}
	fmt.Println("After writing to jobs channel", fmt.Sprint(time.Now().UnixMicro()))

	close(jobs) // Closing a channel indicates that no more values will be *SENT* on it. Doesn't mean to end processing.

	fmt.Println("After closing jobs channel", fmt.Sprint(time.Now().UnixMicro()))

	for a := 1; a <= numJobs; a++ {
		<-results // blocking operatin on the main() thread until all values are received.
	}

	fmt.Println("End", fmt.Sprint(time.Now().UnixMicro()))

}
