package main

import (
	"fmt"
)

type name struct {
	givenName  string
	middleName string
	familyName string
}

type person struct {
	name
	birthPlace string
}

func (n name) describe() string {
	return fmt.Sprintf("Name=%v", n.givenName)
}

// Commenting this out makes the program use the embedded struct's describe method
func (p person) describe() string {
	return fmt.Sprintf("Birth Place=%v", p.birthPlace)
}

type describer interface {
	describe() string
}

func main() {
	fmt.Println("Begin")

	n := name{givenName: "Marge", middleName: "Jane", familyName: "Simpson"}
	p := person{name: n, birthPlace: "Tempe"}

	fmt.Printf("%v %T\n", p, p)

	/*
		Since container embeds base, the methods of base also become methods of a container.
		Here we invoke a method that was embedded from base directly on co.
	*/
	fmt.Println(p.describe())

	var d describer = p
	fmt.Println(d.describe())
}
