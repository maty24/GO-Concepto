package main

import (
	"flag"
	"fmt"
)

func main() {

	flagPattern := flag.String("p", "", "filtra por patron")
	flag.Parse()
	fmt.Println(*flagPattern)
}
