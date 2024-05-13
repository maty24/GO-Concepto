package main

import "fmt"

func main() {
	music := make(map[string]string) //creo un mapa con clave string y valor string
	music["rock"] = "AC/DC"
	music["pop"] = "Madonna"

	fmt.Println(music)

	tech := map[string]string{
		"company":   "Google",
		"language":  "Go",
		"framework": "React",
	}
	fmt.Println(tech)

	//delete(tech, "framework") //borro un elemento del mapa tech

	fmt.Println(music["rock"]) // para acceder a un elemento del mapa
	fmt.Println(tech["sa"])    // si la clave no existe me devuelve un string vacio

	content, ok := tech["sa"] // si la clave no existe me devuelve un string vacio y un booleano false
	fmt.Println(content, ok)

}
