package main

import "fmt"

func main() {
	var age int
	fmt.Println("Enter your age:")
	fmt.Scanln(&age)

	if age >= 18 {
		println("You can vote!")
		numberLoop()
		dayOfWeek()
	} else {
		println("You can't vote yet")
	}
}

func numberLoop() {
	for i := range 10 {
		fmt.Println("Loop:", i)
	}
}

func dayOfWeek() {

	var day int

	fmt.Println(" Enter a number between 1 and 7")
	fmt.Scanln(&day)

	switch day {
	case 1:
		fmt.Println("Sunday")
	case 2:
		fmt.Println("Monday")
	case 3:
		fmt.Println("Tuesday")
	case 4:
		fmt.Println("Wednesday")
	case 5:
		fmt.Println("Thursday")
	case 6:
		fmt.Println("Friday")
	case 7:
		fmt.Println("Saturday")
	default:
		fmt.Println("Invalid day")
	}
}
