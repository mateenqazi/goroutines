package main

import (
	"fmt"
	"time"
)

func printMessage(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Println(msg, ": ", i)
		time.Sleep(time.Millisecond * 500)
	}
}

// func main() {
// 	fmt.Println("Main goroutine starts")

// 	// Start two goroutines concurrently
// 	go printMessage("Hello from Goroutine 1")
// 	go printMessage("Hello from Goroutine 2")

// 	// Give some time for goroutines to finish before exiting the main goroutine
// 	time.Sleep(time.Second * 3)

// 	fmt.Println("Main goroutine ends")
// }
