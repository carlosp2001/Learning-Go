package main

import (
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
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
	// El valor default para variables de funciones es nil
	myFuncVariable = f2
	fmt.Println(myFuncVariable("Hello, World!"))

	calculatorExample()
	anonymousFunctionExample()

	// Closures ejemplo
	closureExample()
	usingClosureToSortData()

	// Retornando funciones de funciones con closures
	twoBase := makeMult(3)
	threeBase := makeMult(3)
	twoBase(2)
	threeBase(3)
}

func makeMult(base int) func(int) int {
	return func(factor int) int {
		return base * factor
	}
}

func usingClosureToSortData() {
	type Person struct {
		FirstName string
		LastName  string
		Age       int
	}

	people := []Person{
		{"Pat", "Patterson", 37},
		{"Tracy", "Bobdaughter", 23},
		{"Fred", "Fredson", 18},
	}

	fmt.Println(people)

	sort.Slice(people, func(i, j int) bool {
		return people[i].LastName < people[j].LastName
	})

	fmt.Println(people)
}

func closureExample() {
	a := 20
	f := func() {
		fmt.Println("a is", a)
		a = 30 // Esta función puede leer y modificar el valor de a, incluso sin haber sido declarada y sin haber sido pasada como argumento
		// También puede suceder shadowing de la variable
		a := 40
		fmt.Println("inner a is", a)
	}
	f()
	fmt.Println("a is", a)
}

func anonymousFunctionExample() {
	f := func(j int) {
		fmt.Println("printing", j, "from inside of an anonymous function")
	}

	for i := 0; i < 5; i++ {
		f(i)
		// No es necesario asignar una función anónima a una variable
		func(j int) {
			fmt.Println("printing", j, "from inside of an anonymous function")
		}(i)
	}
}

func calculatorExample() {
	expressions := [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "/", "3"},
		{"2", "%", "3"},
		{"two", "+", "three"},
		{"5"},
	}

	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Println("invalid expression:", expression)
			continue
		}
		pl, err := strconv.Atoi(expression[0])
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
		result := opFunc(pl, p2)
		fmt.Println(result)
	}
}

var opMap = map[string]func(int, int) int{
	"+": add,
	"-": sub,
	"*": mul,
	"/": divCalc,
}

func add(i int, j int) int {
	return i + j
}

func sub(i int, j int) int {
	return i - j
}

func mul(i int, j int) int {
	return i * j
}

func divCalc(i int, j int) int {
	return i / j
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
