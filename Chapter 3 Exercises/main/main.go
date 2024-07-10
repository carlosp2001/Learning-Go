package main

import (
	"fmt"
)

func main() {
	firstExercise()
	secondExercise()
	thirdExercise()
}

func thirdExercise() {
	type Employee struct {
		firstName string
		lastName  string
		id        int
	}

	employee := Employee{"Sam", "Anderson", 1001}
	employee1 := Employee{id: 1002, firstName: "John", lastName: "Doe"}
	var employee2 Employee
	employee2.lastName = "Smith"
	employee2.id = 1003
	employee2.firstName = "Jane"
	fmt.Println(employee, employee1, employee2)
}

func secondExercise() {
	var message = "Hi 👦 and 👧"
	fmt.Println(string(message[3]))
}

func firstExercise() {
	var greetings = []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}
	var firstSlice = greetings[:2]
	var secondSlice = greetings[1:4]
	var thirdSlice = greetings[3:5]
	fmt.Println(firstSlice, secondSlice, thirdSlice)
}
