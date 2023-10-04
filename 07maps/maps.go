package main

import (
	"fmt"
)

func main() {

	fmt.Println("Maps in Go lang")

	// Creating map

	languages := make(map[string]string) // make creates map, it takes 2 parameters - make(map[type of key]type of value)

	// Adding value to map

	languages["js"] = "JavaScript"
	languages["rct"] = "React"
	languages["rb"] = "Ruby"

	fmt.Println("Languages map :- ", languages)
	fmt.Println("Value of js is :- ", languages["js"])

	// Deleting value from map

	delete(languages, "js")

	fmt.Println("New map :- ", languages)

	// Looping over languages map

	for key, value := range languages {
		fmt.Printf("For the key %v, value is %v\n ", key, value)
	}
}
