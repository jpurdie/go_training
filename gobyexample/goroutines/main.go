package main

import (
	"fmt"
	"time"
)

func myFunc(s string) {
	for i, c := range s {
		fmt.Println(i, string(c))
	}
}
func myFunc2(s string) {
	fmt.Println(s)

}

func main() {
	fmt.Println("Begin")

	myFunc("foobar")

	go myFunc("select")

	go myFunc2("going")

	time.Sleep(time.Second)

	fmt.Println("Done")

}
