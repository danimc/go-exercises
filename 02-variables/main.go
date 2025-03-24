package main

import "fmt"

func main() {
	const thisYear = 2025
	var name string
	var birthdayYear int

	println("Welcome to the profile data app")
	println("Please enter the following information")
	println("What is your name?")
	fmt.Scanln(&name)

	println("What year were you born?")
	fmt.Scanln(&birthdayYear)

	// USE MATH DECLARATION FOR THIS EXAMPLE
	age := thisYear - birthdayYear

	fmt.Println("Hi", name, "you are", age, "years old")

}
