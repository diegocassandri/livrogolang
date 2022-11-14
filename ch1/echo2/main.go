package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", " "
	for i, arg := range os.Args[0:] {
		s += sep + arg
		fmt.Println("Indice: ", i, " Valor: ", arg)
	}
	fmt.Println(s)
}
