package main

import "fmt"

func main() {

	//var flags [3]string //array de 3 elementos de tipo string

	//le asigno valores a los elementos del array
	//flags[0] = "rojo"
	//flags[1] = "verde"
	//flags[2] = "azul"

	//le pongo ... para que go infiera la cantidad de elementos
	flags := [...]string{"rojo", "verde", "azul", "amarillo"}

	//el primer valor es la posición inicial y el segundo valor es la posición final, el valor final es el indice, es el rango de entre la posición inicial y la posición final
	cars := flags[0:3]    // tambien puedo hacer [:3] ya que go infiere que es desde el inicio
	ultimo2 := flags[2:4] // tambien puedo hacer [2:] ya que go infiere que es hasta el final

	fmt.Println(ultimo2)
	fmt.Println(cars)
	fmt.Println(cars[0]) //para acceder a un elemento del array

	animales := []string{"perro", "gato", "pez", "pajaro"}
	pets := animales[0:2]
	pets = append(pets, "hamster") //agrego un elemento al array pets, el append me devuelve un nuevo array

	fmt.Println(len(pets)) //longitud del array
	fmt.Println(cap(pets)) //capacidad del array

	aves := make([]string, 0, 10) //creo un array con longitud 0 y capacidad 10
	aves = append(aves, "loro")
	fmt.Println(len(aves))
}
