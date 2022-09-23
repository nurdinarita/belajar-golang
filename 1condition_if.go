package main

import "fmt"

func main() {
	var currentYear = 2022

	if age := currentYear - 1998; age < 17 {
		fmt.Println("Kamu belum boleh membuat SIM")
	} else {
		fmt.Println("Kamu sudah boleh membuat SIM")
	}
}
