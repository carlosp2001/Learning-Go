package main

import "fmt"

func main() {
	firstExercise()
	secondExercise()
	thirdExercise()
}

func firstExercise() {
	// Write a program that declares an integer variable called i with the value 20. Assign i to a floating-point variable named f. Print out i and f.
	var i = 20
	var f = float64(i)
	fmt.Println(i, f)
}

func secondExercise() {
	// Write a program that declares a constant called value that can be assigned to both an integer and a floating-point variable. Assign it to an integer called i and a floating-point variable called f. Print out i and f.
	const value = 20
	i := value
	f := float64(value)
	fmt.Println(i, f)
}

func thirdExercise() {
	//Write a program with three variables, one named b of type byte, one named smallI of type int32, and one named bigI of type uint64. Assign each variable the maximum legal value for its type; then add 1 to each variable. Print out their values.
	var b byte
	var bigI uint64
	b = 255
	bigI = 18446744073709551615
	b++
	bigI++
	// En este caso ambos valores imprimen 0 ya que se produce un overflow o desbordamiento de los valores máximos permitidos para los tipos de datos byte y uint64
	fmt.Println(b, bigI)
}
