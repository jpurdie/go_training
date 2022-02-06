package main

import "fmt"

func main() {
	result := fib(10)
	fmt.Print(result)
}

func fib(value int) int {
	fmt.Println("got", value)
	if value == 0 || value == 1 {
		return value
	}
	fmt.Println("Sending ", value-2, value-1)
	return fib(value-2) + fib(value-1)
}
