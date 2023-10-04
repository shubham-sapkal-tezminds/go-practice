package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("How much Go do you know ?")

	input , _ := reader.ReadString('\n')
	fmt.Println("Thanks for replying", input)
	fmt.Printf("The type of input is %T", input)
}