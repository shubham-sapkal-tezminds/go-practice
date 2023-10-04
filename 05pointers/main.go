package main

import "fmt"

func main() {
	//pointers hold reference to memory location and also they can store value 

	// To create pointer use & symbol

	var numberr int = 22

	var ptr = &numberr

	fmt.Println("The memory location of pointer is :- ", ptr) // Only the word ptr stores memory location
	fmt.Println("The memory location of pointer is :- ", *ptr) // When used * it stores the actual value not the copy 

	
}