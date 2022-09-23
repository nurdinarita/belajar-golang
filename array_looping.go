package main

import (
	"fmt"
)

func main() {
	var fruits = [3]string{"apple", "banana", "mango"}

	for i, v := range fruits {
		fmt.Printf("Index : %d, Value : %s\n", i, v)
	}
}
