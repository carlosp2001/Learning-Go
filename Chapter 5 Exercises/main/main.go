package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	firstExercise()
	secondExercise()
	thirdExercise()
}

func thirdExercise() {
	prefixer := func(s string) func(string) string {
		return func(phrase string) string {
			return s + " " + phrase
		}
	}

	prependWithFoo := prefixer("foo")
	fmt.Println(prependWithFoo("bar"))
}

func secondExercise() {
	fileLen := func(filename string) (int, error) {
		file, err := os.Open(filename)
		if err != nil {
			return 0, err
		}

		defer func(file *os.File) {
			err := file.Close()
			if err != nil {

			}
		}(file)

		fileInfo, err := file.Stat()

		if err != nil {
			return 0, err
		}

		return int(fileInfo.Size()), nil
	}

	length, err := fileLen(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Length:", length)
}

func firstExercise() {
	add := func(i int, j int) (int, error) { return i + j, nil }
	sub := func(i int, j int) (int, error) { return i - j, nil }
	multi := func(i int, j int) (int, error) { return i * j, nil }
	div := func(i int, j int) (int, error) {
		if j == 0 {
			return 0, errors.New("division by zero")
		}
		return i / j, nil
	}

	var opMap = map[string]func(int, int) (int, error){
		"+": add,
		"-": sub,
		"*": multi,
		"/": div,
	}

	expressions := [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "/", "3"},
		{"2", "%", "3"},
		{"two", "+", "three"},
		{"5"},
		{"10", "/", "0"},
	}

	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Println("invalid expression:", expression)
			continue
		}
		p1, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Println(err)
			continue
		}
		op := expression[1]
		opFunc, ok := opMap[op]
		if !ok {
			fmt.Println("unsupported operator:", op)
			continue
		}
		p2, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Println(err)
			continue
		}
		result, err := opFunc(p1, p2)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(result)
	}
}
