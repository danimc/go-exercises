package main

import "fmt"

type Person struct {
	name     string
	age      int
	lastName string
	employed bool
}

func (p Person) sayHello(hiFive string) {
	fmt.Println("Hello", p.name, hiFive)
}

func main() {
	// declare new person with struct literal
	person := Person{name: "Jhon", age: 25, lastName: "Doe", employed: true}
	// declare new person without struct literal
	person2 := Person{"Paul", 30, "Smith", false}

	person.sayHello("👋")
	person2.sayHello("How are you?")

}
