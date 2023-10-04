package main

import "fmt"

func main() {
	fmt.Println("Loops in Go Lang")

	days := make([]string, 7)

	days[0] = "Monday"
	days[1] = "Tuesday"
	days[2] = "Wednesday"
	days[3] = "Thursday"
	days[4] = "Friday"
	days[5] = "Saturday"
	days[6] = "Sunday"

	fmt.Println(days)

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	for index, day := range days { // First value is Index and second value is Value when using "range" keyword
		fmt.Printf("The index of %v is %v\n", day, index)
	}

	number := 0

	for number < 10 {

		if number == 8 {
			goto comment // At number 8 it will goto "comment"
		} else if number == 5 {
			number++
			continue // continues the loops but skips the number 5
		} else if number == 9 {
			break // Breaks the loop if number is 7
		}

		fmt.Println("The number is :- ", number)

		number++
	}

comment:
	fmt.Println("You have hit number 8")
}
