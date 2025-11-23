package main

import "fmt"

func main() {
	var number int

	fmt.Print("Enter a number: ")
	fmt.Scan(&number)

	if number%2 == 0 {
		fmt.Println(number, "is an even number!")
	} else {
		fmt.Println(number, "is an odd number!")
	}
}
