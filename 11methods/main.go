package main

import "fmt"

func main() {

	fmt.Println("Methods in Go Lang")

	// creating user using structure

	shubham := User{"Shubham Sapkal", "shubham@gmail.com", 9857937692, true, 26}

	// fmt.Println("Details of user Shubham are :- ", shubham) // Println will only print values and not type

	fmt.Printf("Details of shubham :- %+v\n", shubham) // Printf along with %+v will print key value pair of user

	// fmt.Printf("Name of user is %v and email is %v.", shubham.Name, shubham.Email) // using "name.key" syntax will print that specific value

	shubham.GetStatus()
	shubham.ChangeAge()

	fmt.Printf("Details of shubham :- %+v\n", shubham)

}

// CREATING STRUCTURE IN GO LANG

// SYNTAX

// type structureName struct {
// 	key    type of value
// }

// NOTE :- ALWAYS USE UPPER CASE FOR KEY NAME

type User struct {
	Name   string
	Email  string
	Phone  int
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active : ", u.Status)
}

func (u *User) ChangeAge() {
	u.Age = 27
	fmt.Println("Changed age is :- ", u.Age)
}
