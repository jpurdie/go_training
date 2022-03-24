package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
	Package atomic provides low-level atomic memory primitives useful
	for implementing synchronization algorithms.

	https://pkg.go.dev/sync/atomic

*/

func main() {
	fmt.Println("go run -race main.go")
	runOne()
	runTwo()
}

func runOne() {
	fmt.Println("runOne keeps the atomic counter in sync across goroutines.")
	var ops uint64
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			for c := 0; c < 1000; c++ {
				atomic.AddUint64(&ops, 1)
			}
			wg.Done()
		}()
	}
	wg.Wait() // waiting for all waitgroups to complete
	fmt.Println("runOne :", ops)
}

func runTwo() {
	fmt.Println("runTwo DOES NOT keep the atomic counter in sync.")
	var ops uint64
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			for c := 0; c < 1000; c++ {
				ops++
			}
			wg.Done()
		}()
	}
	wg.Wait() // waiting for all waitgroups to complete
	fmt.Println("runTwo:", ops)
}
