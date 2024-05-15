package main

import (
	"fmt"
	"time"
)

func main() {
	day := time.Now().Weekday()

	switch day {
	case time.Monday:
		fmt.Println("Hoy es lunes.")
	case time.Tuesday:
		fmt.Println("Hoy es martes.")
	case time.Wednesday:
		fmt.Println("Hoy es miércoles.")
	case time.Thursday:
		fmt.Println("Hoy es jueves.")
	case time.Friday:
		fmt.Println("Hoy es viernes.")
	case time.Saturday:
		fmt.Println("Hoy es sábado.")
	case time.Sunday:
		fmt.Println("Hoy es domingo.")
	default:
		fmt.Println("No se pudo determinar el día.")
	}
}
