package main

import "fmt"

// la estructura es un tipo de dato compuesto es como una interfaz de typescript
type Person struct {
	Name       string
	Age        int
	HasChilren bool
}

func main() {
	mati := Person{
		Name:       "Matias",
		Age:        26,
		HasChilren: false,
	}
	fmt.Printf("%+v\n", mati)
}
