package main

import "fmt"

func main() {

	number := []uint{1, 2, 3, 4, 5}
	food := []string{"apple", "banana", "cherry", "date", "grape"}

	for i := range number { //el range devuelve el indice y el valor, el i es el indice y el v es el valor
		number[i] *= 2
	}

	for key, value := range food {
		fmt.Printf("%v: %v\n", key, value)
	}

	fmt.Printf("%v\n", number)
}
