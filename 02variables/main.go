package main

import "fmt"

//Global variable can be defined outside of main function but you need to use "var" keyword, it can be accessed inside of function


var globalVar = "global"

func main() {
	var name string = "shubham"
	fmt.Println(name)
	fmt.Printf("The type of var is : %T \n", name)
	
	var number int = 40
	fmt.Println(number)
	fmt.Printf("The type of var is : %T \n", number)


	var smallVal uint8 = 56
	fmt.Println(smallVal)
	fmt.Printf("The type of var is : %T \n", smallVal)

	var smallFloat float32 = 56.234265443
	//float32 gives number upto 5 decimal points
	fmt.Println(smallFloat)
	fmt.Printf("The type of var is : %T \n", smallFloat)

	var bigFloat float64 = 56.2342654432423453462634
	fmt.Println(bigFloat)
	fmt.Printf("The type of var is : %T \n", bigFloat)

	var isPremiumUser bool = true
	fmt.Println(isPremiumUser)
	fmt.Printf("The type of var is : %T \n", isPremiumUser)

	//DEFAULT VALUES

	var defaulValue int
	fmt.Println(defaulValue)
	
    
	
	
	// Declaring variables in Go lang

	// Implicit type , where if you dont define type of variable "lexer" 
	// gives type for it, and you can only assign value of that type next time

	var valuee = "some number"
    fmt.Printf("type of variable is : %T \n", valuee)
	//valuee = 64 //this will throw error

    // Declaring variable without giving type, 

	numberOf := 344

	fmt.Println(numberOf)
	fmt.Printf("type of variable is : %T \n", numberOf)

    // Global Variable

	fmt.Println(globalVar)
	fmt.Printf("type of variable is : %T \n", globalVar)


}