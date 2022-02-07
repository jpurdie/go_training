package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println("Begin")

	const s1 = "Привет"
	//const s1 = "Hello"
	fmt.Println("Length:", len(s1), " (raw bytes)")

	for i := 0; i < len(s1); i++ {
		fmt.Printf("%d %x %T\n", i, s1[i], s1[i])
	}
	fmt.Println()

	fmt.Println("Count: ", utf8.RuneCountInString(s1), " (runes)")

	// A range loop handles strings specially and decodes each rune along with its offset in the string.
	// Meaning it iterates over the RUNE and not the (utf8) character
	for idx, runeValue := range s1 {
		fmt.Printf("%#U starts at %d %T\n", runeValue, idx, runeValue)
		examineRune(runeValue)
	}

}

func examineRune(r rune) {
	if r == 'l' {
		fmt.Println("found l")
	} else if r == 'П' {
		fmt.Println("found П")
	} else if r == 'z' {
		fmt.Println("found z")
	}
}

//https://go.dev/blog/strings
