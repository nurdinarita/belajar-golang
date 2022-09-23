package main

import (
	"fmt"
)

func main() {
	var fruits = make([]string, 3)
	_ = fruits

	fruits[0] = "apple"
	fruits[1] = "banana"
	fruits[2] = "mango"

	fmt.Printf("%#v", fruits)
}
