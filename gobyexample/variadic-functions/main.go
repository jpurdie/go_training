package main

import "fmt"

func main() {
	fmt.Println("Begin")

	// Called with individual arguments
	printPax("Even Lelio", "Colette Ann", "Jaci Yeshua", "Lilyana Eka")
	// Called with slice of arguments
	passengers := []string{"Even Lelio", "Colette Ann", "Jaci Yeshua", "Lilyana Eka"}
	printPax(passengers...)
}

func printPax(pax ...string) {
	fmt.Printf("Type of variable pax: %T\n", pax)
	fmt.Println("Full list:", pax)
	for i, passenger := range pax {
		fmt.Println(i+1, passenger)
	}
}
