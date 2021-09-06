package main

import (
	"fmt"
	"reflect"
)

func main() {

	helloMessage := "Hello"
	worldMessage := "world"

	fmt.Println(helloMessage, worldMessage)

	// Printf
	nombre := "Platzi"
	cursos := 500
	// usando tipos de dato exactos -> buena práctica
	fmt.Printf("%s tiene mas de %d cursos\n", nombre, cursos)
	// usando variables dinámicas
	fmt.Printf("%v tiene mas de %v cursos\n", nombre, cursos)

	// Sprintf devuelve un string, no imprime en stdout
	message := fmt.Sprintf("%s tiene más de %d cursos", nombre, cursos)
	fmt.Println(message)

	// Conocer el tipo
	fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("cursos: %T\n", cursos)

	fmt.Println("helloMessage:", reflect.TypeOf(helloMessage))
}
