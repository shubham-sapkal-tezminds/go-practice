package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://www.youtube.com/watch?v=cl7_ouTMFh0&list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa&index=26"

func main() {
	fmt.Println("Handling URL in Go lang")

	result, _ := url.Parse(myurl) // It will parse URL and store it in "result"

	fmt.Println(result.Host)
	fmt.Println(result.RawQuery)

	querryparams := result.Query() // Querry stores parameters in better format

	fmt.Println(querryparams["index"]) // To access any value of parameter just use the mentioned format

	// To print all key and values of parameters us for loops as follows

	for name, value := range querryparams {
		fmt.Printf("Param for %v is %v\n", name, value)
		fmt.Println("Param is : ", value)
	}

	// CREATING URL

	createUrl := &url.URL{ // Always use "&" to create reference
		Scheme: "https",
		Host:   "www.google.com",
		// add other values onwards
	}

	newUrl := createUrl.String() // Converting url to string

	fmt.Println("New URL is : ", newUrl)

}
