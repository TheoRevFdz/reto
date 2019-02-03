package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	const pubKey = "bed2e3a8f2d65285bff7810dc68b7b08"
	const privKey = "299782ed5f6ee70e13951e761cfd15c1ee7f759e"
	const ts = "9"

	hasher := md5.New()
	hasher.Write([]byte(ts + pubKey + privKey))

	fmt.Println("1. Buscar por nombre")
	fmt.Println("2. Listar")

	fmt.Print("Ingrese una opcion: ")
	opt, _ := reader.ReadString('\n')
	opt = strings.TrimRight(opt, "\r\n")

	switch opt {
	case "1":
		fmt.Println("Ingreeso 1")
	case "2":
		fmt.Println("Ingreso 2")
	default:
		fmt.Println(fmt.Sprintf("Error: |%s|", opt))
	}
	fmt.Println()
	fmt.Println("Before: " + ts + pubKey + privKey)
	fmt.Println("Hash: " + hex.EncodeToString(hasher.Sum(nil)))
}
