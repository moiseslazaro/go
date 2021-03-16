package main

import "fmt"

func main() {
	valor1 := 1
	valor2 := 2

	if valor1 == 1 {
		fmt.Println("Es 1")

	} else {
		fmt.Println("Es 2")

	}
	// With AND

	if valor1 == 1 && valor2 == 3 {
		fmt.Println("Es verdad AND")
	}

	// with OR
	if valor1 == 1 || valor2 == 3 {
		fmt.Println("Es verdadero OR ")
	}

	// Switch

	switch modulo := 4 % 2; modulo {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}
	// Sin condicion

}
