package main

import (
	"fmt"
)

func normalFunction(message string) {
	fmt.Println(message)
}

// a and b are int
func tripleArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func doubleReturn(a int) (b, c int) {
	return a, a * 2
}

func main() {
	normalFunction("Hello World")
	tripleArgument(1, 2, "hello")
	value := returnValue(2)
	fmt.Println("Value:", value)
	value1, value2 := doubleReturn(2)
	fmt.Println("value1 and value2", value1, value2)
	value3, _ := doubleReturn(2)
	fmt.Println("value3", value3)
}
