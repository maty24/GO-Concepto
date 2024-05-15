package main

import (
	"fmt"
	"strings"
)

func main() {
	lower, upper := convert("matias") //esto es para que me retorne dos valores
	fmt.Println(lower, upper)
	greet("matias", "godoy")

	nums := []int{1, 62, 3, 44, 55}

	result := filter(nums, mayor10) //esto es para que me retorne los numeros mayores a 10
	fmt.Println(result)

	plus10 := sum(10)      //esto es para que me retorne la suma de 10
	fmt.Println(plus10(5)) //esto es para que me retorne la suma de 10 + 5 ya que en la linea anterior le pase el 10
	fmt.Println(plus10(10))

	fmt.Println(sumvariatica(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)) //esto es para que me retorne la suma de los parametros que le envie
	fmt.Println(sumvariatica(1, 3, 6))

	func(name string) { //esto es una funcion anonima que se ejecuta una sola vez
		fmt.Println("Hola", name)
	}("matias")
}

func sumvariatica(nums ...int) int { //esto es para que le enviemos una cantidad de parametros que queramos
	var total int

	for _, v := range nums { //esto es para recorrer los parametros que le enviamos
		total += v //esto es para sumar los parametros que le enviamos
	}

	return total
}

func sum(a int) func(int) int { //esto es para que me retorne una funcion que recibe un entero y retorna un entero
	return func(b int) int { //esto es para que me retorne un entero
		return a + b
	}
}

func mayor10(nums int) bool { //esto es para que me retorne los numeros mayores a 10
	return nums > 10
}

func greet(firstname, lastname string) { //esto es para que reciba dos parametros de tipo string
	fmt.Println("Hola", firstname, lastname)
}

func convert(text string) (string, string) { //le digo que retorna dos strings
	lower := strings.ToLower(text)
	upper := strings.ToUpper(text)

	return lower, upper
}

func filter(nums []int, callback func(int) bool) []int { //va a recibir 2 pqrametros, un slice de enteros y una funcion que recibe un entero y retorna un booleano, y como retorno va a devolver un slice de enteros
	result := make([]int, 0, len(nums)) //esto es para crear un slice de enteros vacio

	for _, v := range nums {
		//esto es para recorrer el slice de enteros
		if callback(v) { //esto es para que si la funcion callback retorna true, entonces se va a agregar el valor al slice de enteros
			result = append(result, v)
		}
	}
	return result
}
