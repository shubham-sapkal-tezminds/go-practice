package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin) //This takes in input and saves it in reader variable
    fmt.Println("Input a number")

	input, _:= reader.ReadString('\n')

    inputNumber, err := strconv.ParseFloat(strings.TrimSpace(input), 64) 
	
	// strconv is package which has ParseFloat function which converts string to number
	// strings is a package which has trim function , trimspace function removes any whitespace from start or end of string

    if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added one to input", inputNumber + 1)
	}

	fmt.Println("you gave the following input", inputNumber)
	fmt.Printf("The type of input is %T", inputNumber)
}