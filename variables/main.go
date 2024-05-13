package main

import "fmt"

func main() {
	//es una agrupacion de variables
	var (
		apple   string             //le asigo el tipo de que es
		platano string = "platano" //asigo un valor
		naranja string = "naranja"
	)
	//el := forma corta de declarar e iniciar una variable, esto infiera automaticamente el tipo de varible que es basandose en el valor que se le asigna
	// var apple , platano, naranja  := "manzana","platanao","naranaja"

	apple = "manzana"

	fmt.Println(apple, platano, naranja)
}
