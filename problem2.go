package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Main goroutine starts")

	// Run the outer function as a goroutine
	go outerFunction()

	time.Sleep(time.Second * 2)

	fmt.Println("Main goroutine ends")
}

func outerFunction() {
	fmt.Println("Outer function starts")

	// Run inner functions as goroutines
	go innerFunction1()
	go innerFunction2()

	time.Sleep(time.Second)

	fmt.Println("Outer function ends")
}

// Inner function 1
func innerFunction1() {
	fmt.Println("Inner function 1 starts")
	time.Sleep(time.Millisecond * 500)
	fmt.Println("Inner function 1 ends")
}

// Inner function 2
func innerFunction2() {
	fmt.Println("Inner function 2 starts")
	time.Sleep(time.Millisecond * 300)
	fmt.Println("Inner function 2 ends")
}
