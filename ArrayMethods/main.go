package main

import "fmt"

func add_values(a []int) int {
	result := 0
	for i := 0; i < len(a); i++ {
		result += a[i]
	}
	return result
}

func max_value(a []int) int {
	max_value := a[0]

	for i := 1; i < len(a); i++ {
		if a[i] > max_value {
			max_value = a[i]
		}
	}
	return max_value
}

func main() {
	var limit int

	fmt.Print("How many values required in the array: ")
	fmt.Scan(&limit)
	numbers := make([]int, limit)

	for i := 0; i < limit; i++ {
		fmt.Printf("Enter element %d: ", i+1)
		fmt.Scan(&numbers[i])
	}
	fmt.Println("The entered array is:", numbers)
	fmt.Println("Sum of the values in the array:", add_values(numbers))
	fmt.Println("Maximum value in the array:", max_value(numbers))
}
