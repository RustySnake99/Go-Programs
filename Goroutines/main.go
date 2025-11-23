package main

import (
	"fmt"
	"time"
)

func hello() {
	for i := 0; i < 5; i++ {
		fmt.Println("\nGoroutine waves Hi!")
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	go hello() // runs concurrently
	for i := 0; i < 5; i++ {
		fmt.Println("\nThe 'main()' function waves back!")
		time.Sleep(500 * time.Millisecond)
	}
}
