package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("1. Buscar por nombre")
	fmt.Println("2. Listar")

	fmt.Print("Ingrese una opcion: ")
	opt, _ := reader.ReadString('\n')

	switch opt {
	case 1:
		fmt.Println("Ingreeso 1")
	case 2:
		fmt.Println("Ingreso 2")
	}
}
