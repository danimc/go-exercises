package main

import "fmt"

type Person struct {
	name     string
	age      int
	lastName string
	employed bool
}

func main() {
	// declare new person with struct literal
	person := Person{name: "Jhon", age: 25, lastName: "Doe", employed: true}

	fmt.Println("Welcome to the profile data app")
	fmt.Println("This is your information", person)
	fmt.Println("your name is", person.name)

	// declare new person without struct literal
	person2 := Person{"Paul", 30, "Smith", false}
	fmt.Println("This is your information", person2)

	fmt.Println("Update your age")
	fmt.Scanln(&person2.age)
	fmt.Println("This is your new information", person2.age)
}
