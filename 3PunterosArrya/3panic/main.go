package main

import "fmt"

func main() {
	division(10, 0)
	fmt.Println("End of main function")
}

func division(dividend, divisor int) {
	validateZero(divisor)
	fmt.Println(dividend / divisor)

}

func validateZero(divisor int) {
	if divisor == 0 {
		panic("Divisor cannot be zero") // el panic es una forma de terminar el programa de forma abrupta
	}
}
