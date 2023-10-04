package main

import "fmt"

func main() {
	fmt.Println("Defer in Go lang")

	// Defer works as last in first out

	defer fmt.Println("one")
	defer fmt.Println("two")
	defer fmt.Println("three")
	defer fmt.Println("four")

	fmt.Println("Hello")

	myDefer()

}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
