package main

import "fmt"

func main() {

	//puntero: variable que almacena la dirección de memoria
	var color string = "rojo"
	var punteroColor *string //estoy haciendo un puntero a un string
	punteroColor = &color    //con el operador & se obtiene la dirección de memoria de la variable color

	fmt.Printf("Tipo de dato: %T, valor: %s, dirección de memoria: %p\n", color, color, &color)
	fmt.Printf("Tipo de dato: %T, valor: %p, desferenciacion: %s\n", punteroColor, punteroColor, *punteroColor)
}
