package main

import (
	"fmt"
	"os"
)

func main() {

	defer fmt.Println(1) //el defer se ejecta cuando se retorna la funcion, en este caso, se ejecuta cuando se termina la funcion main

	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close() //cuando se retorna la funcion, se cierra el archivo

	_, err = file.Write([]byte("hello"))
	if err != nil {
		fmt.Println(err)
		return
	}

}
