package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Blocks
	shadowingBlocks()
	ifAndElseBlocks()
	forLoops()
	usingLabelsWithLoops()
	switchBlocks()
}

func switchBlocks() {
	words := []string{"hello", "apple", "banana", "world"}
	for _, word := range words {
		switch size := len(word); size {
		case 1, 2, 3, 4:
			fmt.Println(word, "is a short word")
		case 5:
			wordLen := len(word)
			fmt.Println(word, "is exactly the right length:", wordLen)
		case 6, 7, 8, 9:
		default:
			fmt.Println(word, "is a long word")
		}
	}

	// Existen casos en los que necesitemos utilizar break para salir no del case sino de un for loop
loop:
	for i := 0; i < 3; i++ {
		switch {
		case i == 0:
			fmt.Println("i == 0")
		case i == 1:
			fmt.Println("i == 1")
			break loop
		case i == 2:
			fmt.Println("i == 2")
		}

	}

	a := 1

	// Es mejor utilizarlo como una expresión
	switch a {
	case 1:
		fmt.Println("a is 1")
	case 2:
		fmt.Println("a is 2")
	case 3:
		fmt.Println("a is 3")
	case 4:
		fmt.Println("a is 4")
	default:
		fmt.Println("a is not 1, 2, 3, or 4")
	}
}

func usingLabelsWithLoops() {
outerLoop:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j == 1 {
				fmt.Println("Continuing outer loop at j == 1")
				continue outerLoop // Salta a la siguiente iteración del bucle exterior
			}
			fmt.Printf("i = %d, j = %d\n", i, j)
		}
	}
}

func forLoops() {
	// Go uses only the for loop
	// The for loop can be used in four ways
	// A complete, C-style for loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// Infinite for loop
	// for {
	// 	fmt.Println("Infinite loop")
	// }

	// One condition for loop
	// This is equivalent to a while loop in other languages
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	// Using break and continue
	// break exits the loop
	for x := 0; x < 10; x++ {
		if x == 5 {
			break
		}
		fmt.Println(x)
	}

	// continue skips the current iteration
	for z := 0; z < 10; z++ {
		if z == 5 {
			continue
		}
		fmt.Println(z)
	}

	// Using for range loop to iterate over slices, arrays, maps and strings
	// Using with slice
	evenVals := []int{2, 4, 6, 8, 10}
	for i, v := range evenVals {
		fmt.Println(i, v)
	}

	// Ignoring the index value
	for _, v := range evenVals {
		fmt.Println(v)
	}

	// Iterating over maps
	m := map[string]int{
		"a": 1,
		"c": 3,
		"b": 2,
	}

	for i := 0; i < 3; i++ {
		fmt.Println("Loop", i)
		for k, v := range m { // We can appreciate that the order of the keys is not guaranteed
			fmt.Println(k, v)
		}
	}

	// Iterating over strings
	samples := []string{"hello", "apple"}
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
		}
		fmt.Println()
	}
}

func ifAndElseBlocks() {
	n := rand.Intn(10)
	if n == 0 {
		fmt.Println("That's too low")
	} else if n > 5 {
		fmt.Println("That's too big", n)
	} else {
		fmt.Println("That's a good number", n)
	}

	// We can define a variable in the if block and use it in both the if and else blocks
	if y := rand.Intn(10); y == 0 {
		fmt.Println("That's too big", n)
	} else {
		fmt.Println("That's a good number", n)

	}
}

func shadowingBlocks() {
	// Shadowing blocks
	// Variables declared in a block are only available within that block
	// If a variable is declared in a block with the same name as a variable in an outer block, the outer variable is shadowed
	// The inner variable is the one that is used
	// The outer variable is not accessible within the inner block
	// The outer variable is accessible after the inner block
	outerVar := 10
	if outerVar > 5 {
		fmt.Println(outerVar)
		outerVar := 5
		fmt.Println(outerVar) // 5
	}
	fmt.Println(outerVar)

	// Shadowing fmt package
	// The fmt package is shadowed by the fmt variable
	fmt := "oops"
	println(fmt)
}
