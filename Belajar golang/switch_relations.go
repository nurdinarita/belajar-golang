package main

import "fmt"

func main() {
	var score = 8

	switch {
	case score == 8:
		fmt.Println("Perfect")
	case (score < 8) && (score > 3):
		fmt.Println("Not Bad")
	default:
		fmt.Println("Study Harder")
		fmt.Println("You Need To Learn More")
	}
}
