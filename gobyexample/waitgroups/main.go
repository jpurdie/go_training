package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func randInt() int {
	rand.Seed(time.Now().Unix())
	rangeLower := 1
	rangeUpper := 5
	return rangeLower + rand.Intn(rangeUpper-rangeLower+1)
}

func worker(num int, duration int) {
	fmt.Printf("%s worker(%d) sleeping for: %d\n", time.Now().Format("04:05"), num, duration)
	time.Sleep(time.Duration(duration) * time.Second)
	fmt.Printf("%s worker(%d) awake.\n", time.Now().Format("04:05"), num)
}

func callWorker(wg *sync.WaitGroup, i int) {
	defer wg.Done()
	worker(i, randInt())
}

func main() {
	fmt.Printf("-- Non-Closure Method d--\n")
	var wg sync.WaitGroup // A WaitGroup waits for a collection of goroutines to finish
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		i := i // https://go.dev/doc/faq#closures_and_goroutines
		go callWorker(&wg, i)
	}

	fmt.Printf("%s Waiting.\n", time.Now().Format("04:05"))
	wg.Wait()

	// This is the same, but with an anonymous function.
	fmt.Printf("-- Closure Method d--\n")
	var wg2 sync.WaitGroup
	for j := 1; j <= 5; j++ {
		wg2.Add(1)
		j := j
		go func() {
			defer wg2.Done()
			worker(j, randInt())
		}()
	}
	fmt.Printf("%s Waiting.\n", time.Now().Format("04:05"))
	wg2.Wait()
}
