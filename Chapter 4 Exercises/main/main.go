package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	numbersArray := firstExercise()
	secondExercise(numbersArray)
	thirdExercise()
}

func thirdExercise() {
	var total int
	for i := 0; i < 10; i++ {
		total := total + i
		fmt.Println(total)
	}
}

func secondExercise(numbersArrays []int) {
	for _, number := range numbersArrays {
		switch {
		case number%2 == 0:
			fmt.Println("Two!")
		case number%3 == 0:
			fmt.Println("Three!")
		case number%3 == 0 && number%2 == 0:
			fmt.Println("Six!")
		default:
			fmt.Println("Never mind")
		}
	}
}

func firstExercise() []int {
	// Definir un slice con una capacidad para 100 elementos que son los que se va a agregar
	s := make([]int, 0, 100)

	for i := 0; i < 100; i++ {
		randomNumber := rand.IntN(100)
		s = append(s, randomNumber)
	}

	fmt.Printf("Length: %d, Capacity: %d\n", len(s), cap(s))

	return s
}
