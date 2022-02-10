package main

import (
	"fmt"
	"math"
)

type shape interface {
	surfaceArea() float64
	volume() float64
}

type cylinder struct {
	radius float64
	height float64
}

func (c *cylinder) surfaceArea() float64 {
	return (2 * math.Pi * c.radius * c.height) + (2 * math.Pi * math.Pow(c.radius, 2))
}

func (c *cylinder) volume() float64 {
	return math.Pi * math.Pow(c.radius, 2) * c.height
}

type sphere struct {
	radius float64
}

func (s *sphere) surfaceArea() float64 {
	return 4 * math.Pi * math.Pow(s.radius, 2)
}

func (s *sphere) volume() float64 {
	return (4 / 3) * math.Pi * math.Pow(s.radius, 2)
}

type cuboid struct {
	width  float64
	height float64
	length float64
}

func (c *cuboid) surfaceArea() float64 {
	return (2 * c.length * c.width) + (2 * c.length * c.height) + (2 * c.height * c.width)
}

func (c *cuboid) volume() float64 {
	return c.length * c.width * c.height
}

func measure(s shape) {
	fmt.Printf("%v %T\n", s, s)
	fmt.Printf("Surface Area %.2f\n", s.surfaceArea())
	fmt.Printf("Volume %.2f\n", s.volume())
	fmt.Println()
}

func main() {
	fmt.Println("Begin")

	c := &cuboid{2, 3, 3}
	measure(c)

	s := &sphere{3}
	measure(s)

	cyl := &cylinder{3, 4}
	measure(cyl)

}
