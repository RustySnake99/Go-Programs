package main

import "fmt"

func factorial(number uint) uint {
	if number == 0 {
		return 1
	} else if number < 2 {
		return number
	} else {
		return number * factorial(number-1)
	}
}

func main() {
	var x uint

	fmt.Print("Enter a positive number (>= 0): ")
	fmt.Scan(&x)
	fmt.Println("Factorial of", x, "=", factorial(x))
}
