package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://lco.dev" // Put url at the top so it can be accessed anywhere

func main() {
	fmt.Println("Handling web request")

	// To make web request for this example, we are going to use "Get" from "net/http" package

	res, err := http.Get(url)

	checkNilErr(err) // This will panic and stop code execution if there is an error

	data, err := io.ReadAll(res.Body) // read content of responce and store it in data

	checkNilErr(err)

	defer res.Body.Close() // Once the get request is done it will close the connection

	fmt.Println(string(data)) // It will first convert the data to string and then print it
}

// Function to handle error

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
