package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Working with files in Go lang")

	// 1. First we will dicide what we want to put in file

	content := "This will be the content of file."

	// 2. We will create text file and use os.create

	file, err := os.Create("./NewFile.txt") // It can throw an error so we will show that error and stop execution with panic()

	checkNilErr(err)

	// 3. We will write our content inside file using io package and writeString function

	length, err := io.WriteString(file, content) // It can throw an error so we will show that error and stop execution with panic()

	checkNilErr(err)

	fmt.Println("Length of content i.e. string is : ", length)

	defer file.Close() // defer will execute at the end and close the file

	readFile("./NewFile.txt")
}

// Reading file we need os package and readFile function

func readFile(fileName string) { // We give fileName as 1st para and type of data to be read

	fileContent, err := os.ReadFile(fileName)

	checkNilErr(err)

	fmt.Println(string(fileContent))
}

// Use this as common for err

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
