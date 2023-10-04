package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("If else and switch case in Go Lang")

	// IF ELSE

	number := 9

	var result string

	if number > 10 {
		result = "Number is greater than 10"
	} else if number < 10 {
		result = "Number is less than 10"
	} else {
		result = "Number is exactly 10"
	}

	fmt.Println(result)

	// We can also initialize variable and assign value in if else directly or use variable recieved from web request

	if num := 3; num > 10 {
		fmt.Println("Number is greater than 10")
	} else {
		fmt.Println("Number is not greater than 10")
	}

	// SWITCH CASE

	// Random dice number logic

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is :- ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("You got 1")
	case 2:
		fmt.Println("You got 2")
	case 3:
		fmt.Println("You got 3")
	case 4:
		fmt.Println("You got 4")
	case 5:
		fmt.Println("You got 5")
	case 6:
		fmt.Println("You got 6")
	}

}
