package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {

	person := Person{
		Name: "John",
		Age:  17,
	}

	if person.Age >= 18 {
		fmt.Println(person.Name, "es mayor de edad.")
	} else {
		fmt.Println(person.Name, "es menor de edad.")
	}

}
