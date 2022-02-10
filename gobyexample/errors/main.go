package main

import (
	"errors"
	"fmt"
)

func checkWord(foo string) error {
	if foo == "bar" {
		return errors.New("Do not say that.")
	}
	return nil
}

func main() {
	fmt.Println("Begin")
	foo := "bar"
	if e := checkWord(foo); e != nil {
		fmt.Println("Error: ", e)
	} else {
		fmt.Println("No Error:", foo)
	}

}
