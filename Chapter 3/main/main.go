package main

import (
	"fmt"
	"maps"
	"slices"
)

func main() {
	// Tipos compuestos

	// Arrays no es recomendable usarlos directamente
	// En este caso se declara un array de 5 elementos de tipo int con valor 0
	var x [5]int

	// Array con valores iniciales
	var x1 = [3]int{1, 2, 3}
	fmt.Println(x, x1)

	// Aquí se declara un array de 12 elementos de tipo int, pero solo se inicializan los elementos con índice
	// 0, 5, 10 y 11
	var x2 = [12]int{1, 5: 4, 6, 10: 100, 15}
	fmt.Println(x2)

	// Otra forma de declarar un array, esto infiere el tamaño del array
	var x3 = [...]int{1, 2, 3}
	fmt.Println(x3)

	// Se pueden comparar arrays con el operador ==, pero solo si los dos arrays son del mismo tamaño y tipo.
	// Deben tener el mismo tamaño y el mismo valor en cada posición
	var x4 = [3]int{1, 2, 3}
	var x5 = [3]int{1, 2, 3}
	fmt.Println(x4 == x5) // false

	// Slices
	// Un slice es una referencia a un segmento de un array
	var x6 = []int{1, 2, 3}
	fmt.Println(x6)
	x6 = append(x6, 4)
	fmt.Println(x6)

	// El zero value de un slice es nil
	var x7 []int
	fmt.Println(x7)

	// Verificar si dos slices son iguales en tamaño y contenido
	fmt.Println(slices.Equal([]int{1, 2, 3}, []int{1, 2, 3})) // true

	// Len función que devuelve la longitud de un slice
	fmt.Println(len(x6))

	// Usando append para agregar elementos a un slice
	var x8 []int
	x8 = append(x8, 1)

	// Agregar un slice a otro slice usando el operador ...
	x8 = append(x8, []int{2, 3, 4}...)
	capacityUnderstanding()

	// Make función que crea un slice con una longitud y capacidad inicial
	createArrayWithMake()

	// Limpieza de un slice
	emptyingArrayOrSlice()

	// Formas de declarar e inicializar un slice
	declareAndInitializeSlice()

	// Dividir slices
	slicingSlices()

	// Copiar slices
	copyingSlices()

	// Convertir un array en un slice
	convertArrayToSlice()

	// Strings, runes y bytes
	var s = "Hello, 世界"
	var b = s[0] // Acceder a un string por índice devuelve un byte
	fmt.Println(s, b)

	// También funcionan las slice expressions
	var s1 = "Hello, 世界"
	var b1 = s1[4:5] // Acceder a un string por índice devuelve un byte
	fmt.Println(s1, b1)

	// Maps
	// Un map es una colección de pares clave-valor
	// El tipo de mapa se especifica con map[keyType]valueType
	var nilMap map[string]int
	fmt.Println(nilMap)

	// Declarando e inicializando un map
	var m = map[string]int{}
	fmt.Println(m)
	teams := map[string][]string{
		"Orcas":   {"Fred", "Ralph", "Bijou"},
		"Lions":   {"Sarah", "Peter", "Billie"},
		"Kittens": {"Waldo", "Raul", "Ze"},
	}
	fmt.Println(teams)
	fmt.Println(teams["Orcas"])

	readingAndWritingMap()

	// Chequear si una clave existe en un map
	m1 := map[string]int{"hello": 1}
	value, ok := m1["hello1"]
	fmt.Println(value, ok) // No se encuentra la clave en el map

	// Eliminar un elemento de un map
	delete(m1, "hello")

	// Comparar maps
	m2 := map[string]int{"hello": 1}
	m3 := map[string]int{"hello": 1}
	fmt.Println(maps.Equal(m2, m3)) // true

	// Usar un map como un conjunto(set)
	usingMapAsSet()

	// Structs
	definingStructs()
	definingAnonymousStructs()
	compareStructs()
}

func compareStructs() {
	type firstPerson struct {
		name string
		age  int
	}

	type secondPerson struct {
		name string
		age  int
	}

	// Comparar dos structs
	julia := firstPerson{"Julia", 32}
	john := secondPerson{"John", 32}
	fmt.Println(julia == firstPerson(john))
}

func definingAnonymousStructs() {
	// Declarar un struct anónimo
	person := struct {
		name string
		age  int
	}{}
	fmt.Println(person)

	// Crear un struct anónimo con valores iniciales
	julia := struct {
		name string
		age  int
	}{"Julia", 32}
	fmt.Println(julia)
}

func definingStructs() {
	type Person struct {
		name string
		age  int
		pet  string
	}

	// Declarar una variable de tipo Person
	var person Person

	// Un literal struct
	bob := Person{}
	fmt.Println(person, bob)

	// Crear un struct con valores iniciales, si se realiza de esta forma se deben especificar todos los valores
	julia := Person{"Julia", 32, "Dog"}
	fmt.Println(julia)

	// En caso de que no se quiera especificar todos los valores se puede hacer de la siguiente forma, los valores no
	// especificados se inicializan con el zero value
	john := Person{name: "John", pet: "Cat"}
	fmt.Println(john)

	// Acceder a los campos de un struct
	fmt.Println(john.name)

	// Podemos modificar los valores de un struct
	john.pet = "Fish"
	fmt.Println(john)
}

func usingMapAsSet() {
	intSet := map[int]bool{}
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
	for _, v := range vals {
		intSet[v] = true
	}
	fmt.Println(len(vals), len(intSet))
	fmt.Println(intSet[5])
	fmt.Println(intSet[500])
	intSet[100] = true
	if intSet[100] {
		fmt.Println("100 is in the set")
	}

	// Se puede utilizar un struct vacío para simular un set
	intSet1 := map[int]struct{}{}
	vals1 := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
	for _, v := range vals1 {
		intSet1[v] = struct{}{}
	}
	if _, ok := intSet1[5]; ok {
		fmt.Println("5 is in the set")
	}
}

func readingAndWritingMap() {
	totalWins := map[string][]int{}
	totalWins["Orcas"] = []int{3, 4, 2}
	totalWins["Lions"] = []int{5, 8, 9}
	fmt.Println(totalWins["Orcas"])
	fmt.Println(totalWins["Kittens"])
	totalWins["Kittens"] = append(totalWins["Kittens"], 4)
	fmt.Println(totalWins["Kittens"])
}

func convertArrayToSlice() {
	xArray := [4]int{5, 6, 7, 8}
	xSlice := xArray[:] // Se convierte el array en un slice
	fmt.Println(xArray, xSlice)

	// Convertir un subset de un array en un slice
	xArray1 := [4]int{5, 6, 7, 8}
	xSlice1 := xArray1[1:3] // Se convierte el subset del array en un slice
	fmt.Println(xArray1, xSlice1)

	// También podemos convertir un slice a un array con punteros
	xSlice2 := []int{1, 2, 3, 4}
	xArrayPointer := (*[4]int)(xSlice2) // Al convertir un slice a un array con punteros se comparte la misma memoria
	// Si se modifica el slice también se modifica el array
	fmt.Println(xSlice2, xArrayPointer)
}

func copyingSlices() {
	x := []int{1, 2, 3}
	y := make([]int, 3)
	copiedElements := copy(y, x)
	fmt.Println(x, y, copiedElements)

	// Copiar un slice en otro slice con un índice de inicio y fin
	x1 := []int{1, 2, 3, 4}
	num := copy(x1[:3], x1[1:])
	fmt.Println(x1, num)

}

func slicingSlices() {
	x := []string{"a", "b", "c", "d"} // [a b c d]
	fmt.Println(x)
	y := x[:2] // [a b]
	fmt.Println(y)
	z := x[1:] // [b c d]
	fmt.Println(z)
	d := x[1:3] // [b c]
	fmt.Println(d)
	e := x[:] // [a b c d]
	fmt.Println(e)
}

func declareAndInitializeSlice() {
	// Declarar e inicializar un slice
	// Este slice tiene el valor nil
	var data []int
	fmt.Println(data, len(data), cap(data))
	// Para inicializar un slice se puede hacer de la siguiente forma
	// Append crea un nuevo slice con un elemento y lo asigna a data, el slice original se pierde por el garbage collector
	data = append(data, 1)

	// Crear un slice con length 0 y capacity 0
	var x = []int{} // A diferencia de solo declarar un slice, este slice tiene un valor diferente a nil
	fmt.Println(x, len(x), cap(x))

	// Declara e inicializa un slice con 3 elementos
	var y = []int{1, 2, 3} // Este slice tiene un valor diferente 3 elementos y una capacidad de 3
	fmt.Println(y, len(y), cap(y))
}

func emptyingArrayOrSlice() {
	s := []string{"first", "second", "third"}
	fmt.Println(s)
	clear(s)
	fmt.Println(s, len(s), cap(s))
}

func createArrayWithMake() {
	// Se puede crear un slice con una longitud y capacidad inicial
	// En este caso se crea un slice con longitud 0 y capacidad 5
	var x = make([]int, 0, 5)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 1)
	fmt.Println(x, len(x), cap(x))
	// Se puede crear un slice con longitud 5 y capacidad 10
	// En este caso los índices de 0 a 4 son 0
	// Para cambiar el valor de un índice se puede hacer y[0] = 1
	var y = make([]int, 5, 10)
	fmt.Println(y, len(y), cap(y))
	y = append(y, 1)
	fmt.Println(y, len(y), cap(y))
}

func capacityUnderstanding() {
	// Se puede observar que la capacidad de un slice se duplica cada vez que se excede su capacidad
	var x []int
	fmt.Println(x, len(x), cap(x))
	x = append(x, 10)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 20)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 30)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 40)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 50)
	fmt.Println(x, len(x), cap(x))
}
