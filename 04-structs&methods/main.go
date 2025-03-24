package main

import "fmt"

type Person struct {
	name     string
	age      int
	lastName string
	employed bool
}

func main() {

	person := Person{name: "Jhon", age: 25, lastName: "Doe", employed: true}

	fmt.Println("Welcome to the profile data app")
	fmt.Println("This is your information", person)
	fmt.Println("your name is", person.name)
}
