package main

import "fmt"

// las constantes no es neceario que se utilizen para compilar
const (
	os     = "Linux" // go infiere el tipo de dato
	domain = "hcsba.cl"
)

const (
	// iota es una constante que se incrementa en 1 cada vez que se usa
	// se puede usar para crear una secuencia de números
	// parte desde 0
	enero   = iota + 1 // 1
	febrero            // 2
	marzo              // 3
	abril              // 4
	mayo               // 5
	junio              // 6
)

func main() {

	//const os, domain string = "Linux", "hcsba.cl" // constantes no pueden ser reasignadas y no se puede usar el operador de asignación corta :=
	// os = "Windows" // error: no se puede reasignar una constante

	fmt.Println(os, domain)
	fmt.Println(junio)

}
