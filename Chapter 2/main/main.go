package main

import (
	"fmt"
	"math/cmplx"
)

// En Go no se puede definir constantes con la primera letra en mayúscula porque se exportan y se pueden usar en otros paquetes, se deben definir con la primera letra en minúscula
const idKey = "id" // Error: exported const IdKey should have comment or be unexported

func main() {
	//fmt.Println("Hello, World!")
	// Predeclared types
	// bool
	var b bool // Zero value is false
	fmt.Println(b)

	// Numeric types
	var i int8 = 100
	fmt.Println(i)

	// Integer operations
	// Preferiblemente dejar que se infiera el tipo porque si no se infiere, se usará el tipo int el cual puede ser de 32 o 64 bits dependiendo de la arquitectura
	var x = 127
	x *= 2
	fmt.Println(x)

	// Floating-point types
	var f float32 = 3.14
	fmt.Println(f)

	// Complex types
	var c complex64 = complex(20.3, 10.2)
	fmt.Println(c)

	complexNumbers()
	// Strings
	var myFirstInitial rune = 'J'
	fmt.Println(myFirstInitial)

	// Type conversion
	var x1 int = 10
	var y2 float64 = 30.2
	var sum1 float64 = float64(x1) + y2

	var sum2 int = x1 + int(y2)

	fmt.Println(sum1, sum2)

	// Literals are untyped
	// Literales no tienen tipo, se les asigna un tipo dependiendo del contexto
	var x3 float64 = 10
	var y3 float64 = 200.3 * 5

	fmt.Println(x3, y3)

	//var z byte = 1000000 // Error: constant 1000000 overflows byte

	// Ways to declare variables
	// Forma más usada definiendo el tipo
	var x4 int = 10
	// Forma mas usada sin definir el tipo
	var x5 = 10
	// Forma de declarar la variable con zero value
	var x6 int
	// Forma de declarar multiples variables
	var x7, x8 int = 10, 20
	// Forma de declarar multiples variables con zero values
	var x9, x10 int
	// o con difererentes tipos
	var x11, x12 = 10, "hello"

	fmt.Println(x4, x5, x6, x7, x8, x9, x10, x11, x12)
	// Usando una lista de declaración de variables
	var (
		x13    int
		y13        = 20
		z14    int = 30
		d14, e14     = 40, "hello"
		f14, g14 string
	)

	// Forma corta de declaración
	x15 := 10 // Esta forma solo se puede declarar dentro de funciones, a nivel de paquete

	fmt.Println(x13, y13, z14, d14, e14, f14, g14, x15)

	// Constantes
	const (
		idKey   = "id"
		nameKey = "name"
	)
	fmt.Println(idKey, nameKey)
}

func complexNumbers() {
	x := complex(2.5, 3.1)
	y := complex(3.2, 4.7)
	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x * y)
	fmt.Println(x / y)
	fmt.Println(real(x))
	fmt.Println(imag(x))
	fmt.Println(cmplx.Abs(x))
}
