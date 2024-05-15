package main

import (
	"errors"
	"fmt"
	"strconv"
)

var errNotFound = errors.New("not found") //creamos una variable de tipo error, y le asignamos un mensaje

var food = map[int]string{
	1: "fruit",
	2: "vegetable",
}

func main() {
	num, err := strconv.Atoi("1") // me devuelve un entero y un error, el strconv.Atoi convierte un string a un entero, deve ser un numero, si no es un numero, me va a devolver un error
	if err != nil {               //el nil es como el null en otros lenguajes, le decimos que si el error es diferente de nulo, entonces que imprima el error
		fmt.Println(err)
		return
	}
	fmt.Println(num)

	found, err := search("23")
	if err == errNotFound {
		fmt.Println("error controlado")
		return
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(found)
}

func search(key string) (string, error) {
	num, err := strconv.Atoi(key)
	if err != nil {
		return "", err
	}

	frut, err := findFood(num)
	if err != nil {
		return "", err
	}
	return frut, nil
}

func findFood(id int) (string, error) {
	value, ok := food[id] // el food[id] me va a devolver un valor y un booleano, si el valor existe, el booleano va a ser true, si no existe, el booleano va a ser false
	if !ok {
		return "", errNotFound
	}
	return value, nil
}
