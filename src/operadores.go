package main

import (
	"fmt"
	"math"
)

func main() {

	// Area cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area cuadrado:", areaCuadrado)

	x := 10
	y := 50

	// Suma
	result := x + y
	fmt.Println("Suma: ", result)

	// Resta
	result = x - y
	fmt.Println("Resta: ", result)

	// Multiplicación
	result = x * y
	fmt.Println("Multiplicación: ", result)

	// División
	result = x / y
	fmt.Println("División: ", result)

	// Módulo
	result = x % y
	fmt.Println("Módulo: ", result)

	// Incremental
	x++
	fmt.Println("Incremental: ", x)

	// Decremental
	x--
	fmt.Println("Decremental: ", x)

	// Rectánculo
	result = x * y
	fmt.Println("Rectángulo de ", x, " ", y, " = ", result)

	// Trapecio
	var baseUno float64 = 6
	var baseDos float64 = 15
	var alturaTrapecio float64 = 25

	areaTrapecio := ((baseUno + baseDos) * alturaTrapecio) / 2

	fmt.Println("El Area del Trapecio es :", areaTrapecio)

	// Circulo : AreaCirculo = pi por radio al cudrado
	const PI float64 = 3.14 // Constant
	var radioCirculo float64 = 10

	areaCirculo := PI * math.Pow(radioCirculo, 2)
	fmt.Println("El Area del Circulo es :", areaCirculo)
}
