package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	fmt.Println("Begin")
	passengers := []string{}
	fmt.Println("Before: ", len(passengers))

	addPax(&passengers) //aaaaaand give me the address
	fmt.Println("After: ", len(passengers))

}

func addPax(pax *[]string) {

	/*
		https://pkg.go.dev/math/rand
		Random numbers are generated by a Source. Top-level functions, such as Float64 and Int,
		use a default shared Source that produces a deterministic sequence of values each time
		a program is run. Use the Seed function to initialize the default Source if different behavior
		is required for each run. The default Source is safe for concurrent use by multiple goroutines,
		but Sources created by NewSource are not.
	*/
	rand.Seed(time.Now().UnixNano())
	numPax := rand.Intn(100) + 1

	for i := 0; i < numPax; i++ {
		*pax = append(*pax, "pax"+strconv.Itoa(i))
	}
}
