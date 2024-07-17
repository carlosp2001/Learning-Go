package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	// Functions function example
	myFunc, err := MyFunc(MyFuncOpts{
		FirstName: "Carlos",
		LastName:  "Pineda",
		Age:       23,
	})
	if err != nil {
		return
	} else {
		fmt.Println(myFunc)
	}

	// Uso del operador ... para pasar un número variable de parámetros
	usingVariadicParameters()

	// Multiples valores de retorno
	result, remainder, err := divAndRemainder(5, 2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(result, remainder)

	// Múltiples valores de retorno con nombres
	// No es necesario especificar los nombres de los valores de retorno en la llamada
	x, y, z := divAndRemainderNamed(5, 2)
	fmt.Println(x, y, z)

	// Funciones como valores y tipos
	var myFuncVariable func(string) int

	myFuncVariable = f1
	fmt.Println(myFuncVariable("Hello, World!"))
	myFuncVariable = f2
	fmt.Println(myFuncVariable("Hello, World!"))
}

func f1(a string) int {
	return len(a)
}

func f2(a string) int {
	total := 0
	for _, v := range a {
		total += int(v)
	}
	return total
}

// Función con multiples valores de retorno
func divAndRemainder(num, denom int) (int, int, error) {
	if denom == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}

	return num / denom, num % denom, nil
}

// Especificar nombres a los valores devueltos de una función
func divAndRemainderNamed(num, denom int) (result int, remainder int, err error) {
	if denom == 0 {
		err = errors.New("cannot divide by zero")
		return result, remainder, err
	}
	result, remainder = num/denom, num%denom
	return result, remainder, err
}

func usingVariadicParameters() {
	fmt.Println(addTo(3))
	fmt.Println(addTo(3, 2))
	fmt.Println(addTo(3, 2, 4, 6, 8))
	a := []int{4, 3}
	fmt.Println(addTo(3, a...))
	fmt.Println(addTo(3, []int{1, 2, 3, 4, 5}...))
}

func addTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))
	for _, v := range vals {
		out = append(out, base+v)
	}
	return out
}

func div(num int, denom int) int {
	if denom == 0 {
		return 0
	}

	return num / denom
}

// MyFuncOpts Funciones en go no cuentan con parameters con nombre, por lo tanto, si queremos especificarle un nombre a un parámetro
// lo mejor es el uso de structs
type MyFuncOpts struct {
	FirstName string
	LastName  string
	Age       int
}

func MyFunc(opts MyFuncOpts) (string, error) {
	if opts.FirstName == "" {
		return "", fmt.Errorf("FirstName is required")
	} else {
		return fmt.Sprintf("Hello %s %s", opts.FirstName, opts.LastName), nil
	}
}
