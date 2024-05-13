package main

import "fmt"

func main() {

	var a bool = true
	var b string = "Hola"
	var c uint = 10      // el uint es para que no acepte valores negativos
	var d int = -10      // el int acepta valores negativos
	var e rune = 'a'     // el rune es para los caracteres
	var f float32 = 3.14 // el float32 es para los decimales

	var z uint8 = 255
	var x uint16 = 12742
	//esto es para sumar dos variables de diferentes tipos, le pongo el tipo de dato que quiero que sea y lo pongo en parentesis y lo sumo
	v := uint16(z) + x

	fmt.Println(v)

	fmt.Println(b, c, d, e, f)

	//el %T es para mostrar el tipo de dato y el %v es para mostrar el valor
	fmt.Printf("Tipo: %T,Valor: %v", a, a)

}
