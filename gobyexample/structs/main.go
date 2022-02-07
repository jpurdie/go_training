package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println("Begin")
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 30})

	// Preferred
	p1 := person{name: "Jim", age: 22}
	fmt.Println("p1", p1)

	// Not idiomatic
	p2 := person{}
	p2.age = 3
	p2.name = "Hank"
	fmt.Println("p2", p2)

	p3 := &person{"Greg", 31}
	fmt.Printf("p3 %T %v\n", p3, p3)

	p4 := newPerson("Sara", 66)
	fmt.Printf("p4 %T %v\n", p4, p4)

	p4.age = 65
	fmt.Printf("p4 %T %v\n", p4, p4)

	p5 := person{name: "Jill"}
	fmt.Println("p5", p5)
}

func newPerson(name string, age int) *person {
	return &person{
		name, age,
	}
}
