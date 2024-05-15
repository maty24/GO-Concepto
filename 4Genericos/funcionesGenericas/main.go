package main

import "fmt"

func main() {
	PrintList(1, 2, 3, 4, 5)
	PrintList("a", "b", "c", "d", "e")
}

func PrintList(list ...any) {
	for _, item := range list {
		fmt.Println(item)
	}
}
