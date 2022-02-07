package main

import (
	"fmt"
)

func main() {
	fmt.Println("Begin")

	myBox := box{length: 3, width: 5, height: 10}
	v := myBox.volume()
	fmt.Println("Volume is: ", v)
	fmt.Println("----")
	myBoxTwo := boxTwo{length: 3, width: 5, height: 10}
	myBoxTwo.calcVolume()
	fmt.Println("Volume is: ", myBoxTwo.volume)

}

type box struct {
	length int
	width  int
	height int
}

func (b *box) volume() int {
	v := b.height * b.length * b.width
	return v
}

type boxTwo struct {
	length int
	width  int
	height int
	volume int
}

func (b *boxTwo) calcVolume() {
	b.volume = b.height * b.length * b.width
}
