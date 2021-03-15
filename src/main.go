package main

import "fmt"

// Functions

func retunrValue(a int) int {
	return a * 2
}

func main() {
	// Declarion de constantes
	const pi float64 = 3.14
	const pi2 = 3.14
	fmt.Println(pi, pi2)

	// Declaracion de variables
	base := 12
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	// Zero values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	// Calculo del Area de un cuadrado

	const BaseCuadrado = 10
	areaCuadrado := BaseCuadrado * BaseCuadrado
	fmt.Println("Area cuadrado:", areaCuadrado)

	x := 10
	y := 50

	// Suma
	result := x + y
	fmt.Println("Suma:", result)

	// Resta
	result = x - y
	fmt.Println("Resta:", result)

	// Multiplicacion
	result = x * y
	fmt.Println("Multiplicacion:", result)

	// Division
	result = x / y
	fmt.Println("Division:", result)

	// Residuo de la division
	result = x % y
	fmt.Println("Residuo:", result)

	// Incremental y decremental
	x++
	y--
	fmt.Println(x, y)

	// Calcula el area de un circulo
	const radio = 10
	const pi3 = 3.14
	result2 := radio * pi3
	fmt.Println("Area:", result2)

	// Printf

	nombre1 := "Leo"
	nombre2 := "Mariafernanda"
	fmt.Printf("%s y %s son hermanos\n", nombre1, nombre2)
	fmt.Printf("%v y %v son hermanos\n", nombre1, nombre2)
	fmt.Printf("nombre1 : %T\n", nombre1)
	fmt.Printf("pi3 : %T\n", pi3)

	// Sprintf
	message := fmt.Sprintf("%v y %v son hermanos", nombre1, nombre2)
	fmt.Println(message)

	value := retunrValue(2)
	fmt.Println("Value", value)

}
